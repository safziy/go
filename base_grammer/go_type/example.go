package main

import (
	"fmt"
	"unsafe"
)

const (
	_        = iota             // iota = 0
	KB int64 = 1 << (10 * iota) // iota = 1
	MB                          // 与 KB 表达式相同，但 iota = 2
	GB
	TB
)

const (
	A, B = iota, iota << 10 // 0, 0 << 10
	C, D                    // 1, 1 << 10
	E    = iota
	F    = iota
)

func main() {
	// 多变量赋值时，先计算所有相关值，然后再从左到右依次赋值。
	data, i := [3]int{0, 1, 2}, 0
	i, data[i] = 2, 100 // // (i = 0) -> (i = 2), (data[0] = 100)
	println(i, data[0], data[1], data[2])

	// 编译器会将未使⽤的局部变量当做错误。
	d := 0 // Error: d declared but not used
	_ = d  // 可以使用 `_ = d` 来规避Error

	// 注意重新赋值与定义新同名变量的区别。
	s := "abc"
	println(&s)
	s, y := "hello", 20 // 重新赋值: 与前 s 在同⼀层次的代码块中，且有新的变量被定义。
	println(&s, y)      // 通常函数多返回值 err 会被重复使⽤。
	{
		s, z := 1000, 30 // 定义新同名变量: 不在同⼀层次代码块。
		println(&s, z)
	}
	println(&s)

	// 关键字 iota 定义常量组中从 0 开始按⾏计数的⾃增枚举值。
	println(KB, MB, GB, TB)

	// 在同⼀常量组中，可以提供多个 iota，它们各⾃增⻓。
	println(A, B, C, D, E, F)

	// 单引号字符常量表⽰示 Unicode Code Point，⽀支持 \uFFFF、\U7FFFFFFF、\xFF 格式。 对应 rune 类型，UCS-4。
	fmt.Printf("%T\n", 'a')
	var c1, c2 rune = '\u6211', '们'
	println(c1 == '我', string(c2) == "\xe4\xbb\xac")

	f := struct {
		s string
		x int
	}{"abc", 100}

	p := uintptr(unsafe.Pointer(&f)) // *struct -> Pointer -> uintptr
	p += unsafe.Offsetof(f.x)        // uintptr + offset
	p2 := unsafe.Pointer(p)          // uintptr -> Pointer
	px := (*int)(p2)                 // Pointer -> *int
	*px = 200                        // d.x = 200

	fmt.Printf("%#v\n", f) // 输出 struct { s string; x int }{s:"abc", x:200}

	println("aaaa")

}
