package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("goroutine ", i)
			time.Sleep(time.Second)
		}
		c <- true
	}()
	// 阻塞主协程，等待其他协程返回的数据
	<- c
	fmt.Println("over")
}