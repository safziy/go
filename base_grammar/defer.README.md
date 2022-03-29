## Golang中defer的规则

#### 规则一 当defer被声明时，其参数就会被实时解析
```go
package main

import "fmt"

func f() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func main() {
	f()
}
```
> 输出结果为<br>
> 0

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
> 输出结果为<br>
> f defer2: 1<br>
> f defer1: 2<br>
> f return: 0<br>


#### 规则三 defer可以读取有名返回值
```go
package main

import "fmt"

func f()  (i int) {
	defer func() {
		i++
	}()
	return 1
}

func main() {
	fmt.Println("f return:", f())
}
```
> 输出结果为<br>
> f return: 2