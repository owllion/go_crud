# defer執行時機

- 作用: 確保某a函數在主函數結束執行之前，他一定會執行到該a函數
- 範例
    - gRPC建立server後，會使用grpcServer.Serve(lis)去持續運作
    - Serve 是一個阻塞調用，它會一直運行，直到服務器停止或發生錯誤，因此 defer 語句只有在 Serve 方法返回後才會執行。
