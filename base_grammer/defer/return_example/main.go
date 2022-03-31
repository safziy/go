package main

import "fmt"

// 1 修改的是局部变量
func deferReturn1() int {
	a := 1
	defer func() {
		a++
	}()
	return a
}

// 2 defer可以读取到有名的返回值
func deferReturn2() (a int) {
	defer func() {
		a++
	}()
	return 1
}

// 1
func deferReturn3() (b int) {
	a := 1
	defer func() {
		a++
	}()
	return 1
}

// 1?
func deferReturn4() (a int) {
	defer func(a int) {
		a++
	}(a)
	return 1
}

func main() {
	fmt.Println(deferReturn1())
	fmt.Println(deferReturn2())
	fmt.Println(deferReturn3())
	fmt.Println(deferReturn4())
}
