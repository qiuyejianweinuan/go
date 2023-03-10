package main

import (
	"fmt"
	"goTest/goroutines/utils"
	"time"
)

// 通道(channels) 是连接多个协程的管道。
// 你可以从一个协程将值发送到通道，然后在另一个协程中接收。
func main() {

	num1 := utils.Num1
	fmt.Println(num1)
	// 使用 make(chan val-type) 创建一个新的通道。
	//通道类型就是他们需要传递值的类型。
	messages := make(chan string, 2)

	done := make(chan bool)
	// 使用 channel <- 语法 发送 一个新的值到通道中。
	//这里我们在一个新的协程中发送 "ping" 到上面创建的 messages 通道中。
	go func() { messages <- "ping" }()

	// 使用 <-channel 语法从通道中 接收 一个值。
	//这里我们会收到在上面发送的 "ping" 消息并将其打印出来
	msg := <-messages
	//我们运行程序时，通过通道，
	//成功的将消息 "ping" 从一个协程传送到了另一个协程中
	fmt.Println(msg)
	go func() {
		for {
			i, m := <-messages
			if m {
				fmt.Println(i)
			} else {
				done <- true
				return
			}
		}
	}()
	messages <- "www"
	messages <- "ww333w"

	time.Sleep(2 * time.Second)

	// 默认发送和接收操作是阻塞的，直到发送方和接收方都就绪。
	//这个特性允许我们，不使用任何其它的同步操作， 就可以在程序结尾处等待消息 "ping"
}
