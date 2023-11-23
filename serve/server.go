package serve

import (
	"fmt"
	"golang.org/x/net/websocket"
	"gopkg.in/yaml.v3"
	"lacos/channel"
	"lacos/config"
	lu "lacos/util"
	"log"
	"os"
	"os/signal"
	"strconv"
)

func InitDiscovery(GConfig interface{}) {
	v, err := lu.ReadConfig("bootstrap")
	if err != nil {
		fmt.Println(err)
	}
	err = v.Unmarshal(GConfig)
	if err != nil {
		fmt.Println(err)
	}
}

func GetRemoteConfig(GConfig interface{}) {
	InitDiscovery(config.RCofnig)
	serverAddr := "ws://" +
		config.RCofnig.Discovery.Host +
		":" +
		strconv.Itoa(config.RCofnig.Discovery.Port) +
		"/" +
		config.RCofnig.Discovery.ConfigName
	ws, err := websocket.Dial(serverAddr, "", "http://"+config.RCofnig.Discovery.Host)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	fmt.Println("成功建立 WebSocket 连接")
	go func() {
		for {
			var message = make([]byte, 512)
			n, err := ws.Read(message)
			if err != nil {
				log.Println("读取消息失败:", err)
				return
			}
			err = yaml.Unmarshal(message[:n], GConfig)
			fmt.Printf("全局配置%+v\n", GConfig)
			channel.Channel <- 1
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("接收到中断信号，关闭程序...")
}
