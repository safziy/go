package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println("goroutine ", i)
			time.Sleep(time.Second)
		}
	}()
	// 主协程执行完后，所有协程都会结束
	for i := 0; i < 10; i++ {
		fmt.Println("main ", i)
		time.Sleep(time.Second)
	}
}