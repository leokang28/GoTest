package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(header string, channel chan<- string) {
	for {

		// 生产发送
		channel <- fmt.Sprintf("%s:%v", header, rand.Int31())

		time.Sleep(time.Second)
	}
}

func customer(channel <-chan string) {
	// 接受消费
	for {
		message := <-channel
		fmt.Println(message)
	}
}

func main() {
	// 字符串类型管道
	channel := make(chan string)

	// goroutine
	go producer("cat", channel)
	go producer("dog", channel)

	customer(channel)
}
