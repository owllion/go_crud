package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func DownloadImages(urls []string, folderPath string, c chan<- string, stop chan<- string) {
	// 獲取用戶的家目錄路徑
	//QA: 哪裡可以查到?
	homeDir := os.Getenv("USERPROFILE")
	if homeDir == "" {
		// 如果環境變數不存在，可以使用預設值或者報錯
		fmt.Println("USERPROFILE 環境變數不存在")
		return
	}

	// 構建下載文件夾路徑
	folderPath = filepath.Join(homeDir, "Downloads")

	// 遍歷所有圖片 URL 並下載
	counter := 0
	for i, url := range urls {
		// 下載圖片
		counter++
		err := Download(url, folderPath, counter)
		if err != nil {
			fmt.Printf("Error downloading image from %s: %v\n", url, err)
			stop <- err.Error()
		}
		c <- fmt.Sprintf("Download Image %d successfully", i+1)
	}
	stop <- "done"
}

func Download(url string, folderPath string, counter int) error {

	// fmt.Println("folderPath----------------------", folderPath)
	// 下載圖片數據
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Receive non 200 status code")

	}

	// 創建保存文件
	// 清理文件名，刪除無效字符
	// filename := filepath.Base(url)

	now := time.Now().UTC().Format("2006-01-02")
	filepath := filepath.Join(folderPath, now+"-"+strconv.Itoa(counter)+".jpg")
	// fmt.Println("filepath-------------", filepath)
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 將圖片數據寫入文件
	_, err = io.Copy(file, resp.Body)
	return err
}

func main() {
	urls := []string{
		"https://images.unsplash.com/photo-1682687220305-ce8a9ab237b1?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.unsplash.com/photo-1707499929621-8e55b938aeb7?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw0fHx8ZW58MHx8fHx8",
		"https://images.unsplash.com/photo-1558585918-601f4cf404b0?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw1fHx8ZW58MHx8fHx8",
		"https://images.unsplash.com/photo-1707216171962-9f1514c0bda6?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw5fHx8ZW58MHx8fHx8",
		"https://images.unsplash.com/photo-1683009427041-d810728a7cb6?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxMXx8fGVufDB8fHx8fA%3D%3D",
		"https://images.unsplash.com/photo-1544009520-e2ea9189f15e?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OHx8bm9yd2F5fGVufDB8fDB8fHww",
		"https://images.unsplash.com/photo-1566230724840-0fe03c62884d?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MjJ8fG5vcndheXxlbnwwfHwwfHx8MA%3D%3D",
		"https://images.unsplash.com/photo-1582920771007-0a1080b99a31?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MjZ8fG5vcndheXxlbnwwfHwwfHx8MA%3D%3D",
		"https://images.unsplash.com/photo-1502021680532-838cfc650323?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MzB8fG5vcndheXxlbnwwfHwwfHx8MA%3D%3D",
	}

	numOfPics := len(urls)
	//Buffered Channel
	c := make(chan string, numOfPics)
	stop := make(chan string)
	//預設下載到Downloads
	go DownloadImages(urls, "", c, stop)

	for {
		select {
		case msg := <-c: //有收到訊息就印出
			fmt.Println("收到的msg----------", msg)
		case <-stop:
			fmt.Println("完成下載")
			return
		}

	}

}
