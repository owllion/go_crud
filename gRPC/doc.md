# gRPC 創建流程

-------gRPC---------------------------------

1. .proto 檔案：
- 創建 .proto檔案，定義service和他的req/res(不僅定義了服務 (service)，還定義了該服務所使用的資料結構 (message))

2. 生成 .pb & gRPC code：
- 跑指令(要安裝兩個go套件) 把 .proto 檔案轉換為 .pb.go 檔案，此檔包含了 gRPC 服務的客戶端和伺服器代碼，白話文: 裡面就是你呼叫gRPC的一些設定，初始化gRPC server需要用到，下面是安裝指令(要用go install)
``` shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

3. 服務實現(Impl)：
- 新增該服務的實現 (Impl) 檔案。此檔案包含了服務的具體邏輯，例如：UserService 可能包含 getStudent 和 updateStudent 等功能。
  - !!注意!!
     必須在其中定義一個結構體，例如 <服務名稱>ServiceImpl，這將用於具體實現 gRPC 服務。這個結構體需要嵌入 Unimplemented<服務名稱>ServiceServer。這是為了保證兼容性，使得當 .proto 文件中添加了新方法時，不必強制實現所有方法。

4. gRPC 伺服器：
- 都好了後就寫server檔，(可能不用寫但)可以加上close client的defer函數，會在gRPC server 被終止/關閉後執行(可確保當伺服器關閉時，進行釋放資源或進行清理工作)

5. gRPC 客戶端：
- 要初始化一個gRPC client，在server中引入並統一初始化(client.InitGRPCClient())，就不用每次呼叫api都要重新初始化一次

6. ??
- 去根目錄的server.go引入 gRPC server，要注意其中一個server要運行在另個goroutine裡面，因為server是blocking的



----------GIN----------------------------------
1. Controller 邏輯：
- 以上都完成後，可去controller開始撰寫相關邏輯CRUD，就像寫一般的Restful API 一樣。可以想象為一個在 GIN 中的橋接器，使得 HTTP 請求可以通過 GIN 伺服器被轉發到 gRPC 服務。
- 補充: 當你的 GIN 路由接收到 HTTP 請求後，它會用 gRPC 客戶端呼叫你的 gRPC 服務。
- 補充: 要修改service的邏輯，去Impl修改即可，不用重新產生.pb.go檔案那些!

2. 定義路由：
- 在 GIN 的路由設定中添加對應的路由，並連接到相應的 controller。

3. 啟動伺服器：
- 完成後就一樣把route給寫到router裡面
- 執行go run . 啟動伺服器(兩個都要，目前測試只啟動一個會導致兩個都沒開)
- 去Postman測試，port是GIN server的port，呼叫後會打到 ??
