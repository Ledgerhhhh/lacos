package main

import (
	"fmt"
	"lacos/channel"
	"lacos/config"
	"lacos/serve"
	"net/http"
)

func init() {
	go func() {
		serve.GetRemoteConfig(config.GConfig)
	}()
}

func main() {
	<-channel.Channel
	http.HandleFunc("/editConfig", handler)
	port := 11111
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("===============%+v", config.GConfig)
	// 启动 HTTP 服务器
	fmt.Printf("Server is listening on port %d...\n", port)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	// 处理请求的逻辑
	fmt.Println("hhhh")
}
