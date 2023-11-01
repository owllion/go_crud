package main

import (
	"context"
	"fmt"
	"log"
	"time"

	etcd "go.etcd.io/etcd/client/v3"
)

func main() {
	// 創建 etcd 客戶端
	client, err := etcd.New(etcd.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd 服務的地址
		DialTimeout: 5 * time.Second,
	})
	//time.Second 單位是那納秒，常用來做延遲、超時、間隔等功能
	
	if err != nil {
		fmt.Println("Failed to create etcd client:", err)
		return
	}
	defer client.Close()

	// 創建一個租賃
	leaseGrantResp, err := client.Grant(context.TODO(), 60) // 60 秒的 TTL
	if err != nil {
		fmt.Println("Failed to grant lease:", err)
		return
	}

	// 租賃的 ID
	leaseID := leaseGrantResp.ID

	// 定義一個註冊信息
	key := "serviceA_addr"
	value := "localhost:50051"

	// 將註冊信息設置到 etcd，並關聯到租賃
	_, err = client.Put(context.TODO(), key, value, etcd.WithLease(leaseGrantResp.ID))
	if err != nil {
		fmt.Println("Failed to put key:", err)
		return
	}

	fmt.Println("Registered serviceA with address:", value)

	// 進行續租操作
	keepAliveChan, err := etcd.Lease.KeepAlive(context.TODO(), leaseID)
	if err != nil {
		log.Fatal(err)
	}

	// 監聽 keepAliveChan 以獲得續租的更新
	go func() {
		for ka := range keepAliveChan {
			fmt.Println("Received keepalive response", ka)
		}
	}()


	// 模擬程序運行一段時間，例如10秒
	time.Sleep(10 * time.Second)
}
