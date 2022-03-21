## Golang中defer的规则

#### 规则一 当defer被声明时，其参数就会被实时解析

例如下面代码中，只会输出0

```go
package main

import "fmt"

func demo() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
```

#### 规则二 defer执行顺序为先进后出

```go
package main

import "fmt"

func f() int {
	var i int
	defer func() {
		i++
		fmt.Println("f defer1:", i)
	}()

	defer func() {
		i++
		fmt.Println("f defer2:", i)
	}()
	return i
}

func main() {
	fmt.Println("f return:", f())
}
```


#### 规则三 defer可以读取有名返回值