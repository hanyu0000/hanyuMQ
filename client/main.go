package main

import (
	"hanyuMQ/client/pks"
	"fmt"
	"os"

	"time"
)

func main() {
	
	option := os.Args[1]
	port := ""
	if len(os.Args) == 3{
		port = os.Args[2]
	}else{
		port = "null"
	}
	
	switch option{
	case "p":

		//创建一个生产者实例，连接到 Zookeeper 地址 0.0.0.0:2181
		//"producer1" 是生产者的名字（用来注册）
		producer, _ := clients.NewProducer("0.0.0.0:2181", "producer1")

		for {
			msg := clients.Message{
				Topic_name:    	"phone_number",// 要发送的主题（类似一个分类）
				Part_name:      "yclchuxue", // 分区名（为了支持并发、负载均衡）
				Msg:  			[]byte("18788888888"),// 实际发送的消息内容，手机号
			}
			err := producer.Push(msg, -1)
			if err != nil {
				fmt.Println(err)
			}

			time.Sleep(5 * time.Second)
		}

	case "c":	
		consumer,_ := clients.NewConsumer("0.0.0.0:2181", "consumer1", port)
		go consumer.Start_server()

		consumer.Subscription("phone_number", "yclchuxue", 2)

		consumer.StartGet(clients.Info{
			Offset: 0,
			Topic: "phone_number",
			Part: "yclchuxue",
			Option: 2,
		})
	}

}
