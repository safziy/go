package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	goroutineLeak1()
}

func goroutineLeak1() {
	test := func(dataChan chan int) {
		// do something
		dataChan <- 1 // 向一个非缓冲的通道的写入操作会阻塞,直到该有消费者从该通道读取数据
	}
	intChan := make(chan int)
	go test(intChan)
	time.Sleep(time.Second * 2)
	fmt.Println(runtime.NumGoroutine()) // 2
}
