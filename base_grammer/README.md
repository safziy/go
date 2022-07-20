## go基础语法

---

### 1.[类型](go_type/README.md)

#### 1.1 [变量]()

##### goroutine泄露
goroutine泄露是指客户端生成一个routine来做一些异步任务，并在任务完成后将数据写入一个channel，但是

- 没有监听程序消耗该channel的数据写入
```go
func main() {
	test := func(dataChan chan int) {
		// do something
		dataChan <- 1 // 向一个非缓冲的通道的写入操作会阻塞,直到该有消费者从该通道读取数据 
	}
	intChan := make(chan int)
	go test(intChan)
	time.Sleep(time.Second * 2)
	fmt.Println(runtime.NumGoroutine()) // 2
}
```

- 
