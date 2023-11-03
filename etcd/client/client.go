package etcdClient

import (
	"context"
	"fmt"
	"time"

	etcd "go.etcd.io/etcd/client/v3"
)

func StartEtcd() {
	// 創建 etcd 客戶端
	client, err := etcd.New(etcd.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd 服務的地址
		DialTimeout: 10 * time.Second,
	})
	//time.Second 單位是那納秒，常用來做延遲、超時、間隔等功能
	fmt.Println("*---------------------------------client", client)
	if err != nil {
		fmt.Println("Failed to create etcd client:", err)
		return
	}
	defer client.Close()

	fmt.Println("-------------------client.Close--------------------")

	//創建一個租賃
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	leaseGrantResp, err := client.Grant(ctx, 60)

	// leaseGrantResp, err := client.Grant(context.TODO(), 60) // 60 秒的 TTL
	fmt.Println("----------------------------leaseGrantResp---------------", leaseGrantResp)
	if err != nil {
		fmt.Println("Failed to grant lease:", err)
		return
	}

	//定義一個註冊信息
	key := "serviceA_addr"
	value := "localhost:50051"
	fmt.Println("-----------------------k,v:------------------------", key, value)

	//將註冊信息設置到 etcd，並關聯到租賃
	//lease一段時間後就會失效
	_, err = client.Put(context.TODO(), key, value, etcd.WithLease(leaseGrantResp.ID))
	if err != nil {
		fmt.Println("Failed to put key:", err)
		return
	}
	fmt.Println("------------------------after Put --------------------------------")

	fmt.Println("Registered serviceA with address:", value)


	// 模擬程序運行一段時間，例如10秒
	// time.Sleep(10 * time.Second)
}
