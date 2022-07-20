### 类型

--- 

#### 1.变量

Go 是静态类型语⾔，不能在运⾏期改变变量类型。

使⽤关键字 var 定义变量，⾃动初始化为零值。如果提供初始化值，可省略变量类型，由编译器⾃动推断。
```go
var x int
var f float32 = 1.6
var s = "abc"
```

在函数内部，可⽤更简略的 ":=" ⽅式定义变量。
```go
func main() {
    x := 123 // 注意检查，是定义新局部变量，还是修改全局变量。该⽅式容易造成错误。
}
```

可一次定义多个变量。
```go
var x, y, z int
var s, n = "abc", 123

var (
	a int
	b float32
)

func main() {
	n, s := 0x1234, "Hello, World!"
	println(x, s, n)
}
```

多变量赋值时，先计算所有相关值，然后再从左到右依次赋值。
```go
data, i = [3]int{0, 1, 2}, 0
i, data[i] = 2, 100   // (i = 0) -> (i = 2), (data[0] = 100)
```

特殊只写变量 "_"，⽤于忽略值占位。
```go
func test() (int, string) {
	return 1, "abc"
}

func main() {
	_, s := test()
	println(s)
}
```

编译器会将未使⽤的局部变量当做错误。
```go
var s string  // 全局变量没问题。

func main() {
	i := 0    // Error: i declared but not used。(可使用 "_ = i" 规避)
}
```

注意重新赋值与定义新同名变量的区别。
```go
s := "abc"
println(&s)
s, y := "hello", 20 // 重新赋值: 与前 s 在同⼀层次的代码块中，且有新的变量被定义。
println(&s, y) // 通常函数多返回值 err 会被重复使⽤。
{
    s, z := 1000, 30 // 定义新同名变量: 不在同⼀层次代码块。
    println(&s, z)
}
println(&s)
// 输出
// 0xc000046768
// 0xc000046768 20
// 0xc000046738 30
// 0xc000046768
```


#### 2.常量

常量值必须是编译期可确定的数字、字符串、布尔值。
```go
const x, y int = 1, 2 // 多常量初始化
const s = "Hello, World!" // 类型推断
const ( // 常量组
    a, b = 10, 100
    c bool = false
)
func main() {
    const x = "xxx" // 未使⽤局部常量不会引发编译错误。
}
```

不⽀持 1UL、2LL 这样的类型后缀。

在常量组中，如不提供类型和初始化值，那么视作与上⼀常量相同。
```go
const (
    s = "abc"
    x           // x = "abc"
)
```

常量值还可以是 len、cap、unsafe.Sizeof 等编译期可确定结果的函数返回值。
```go
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(b)
)
```

如果常量类型⾜以存储初始化值，那么不会引发溢出错误。
```go
const (
    a byte = 100 // int to byte
    b int = 1e20 // float64 to int, overflows
)
```

##### 枚举

关键字 iota 定义常量组中从 0 开始按⾏计数的⾃增枚举值。
```go
const (
    Sunday = iota   // 0
    Monday          // 1，通常省略后续⾏表达式。
    Tuesday         // 2
    Wednesday       // 3
    Thursday        // 4
    Friday          // 5
    Saturday        // 6
)

const (
    _ = iota                        // iota = 0
    KB int64 = 1 << (10 * iota)     // iota = 1
    MB                              // 与 KB 表达式相同，但 iota = 2
    GB
    TB
)
```

在同⼀常量组中，可以提供多个 iota，它们各⾃增⻓。
```go
const (
	A, B = iota, iota << 10 	// 0, 0 << 10
	C, D 						// 1, 1 << 10
)
```

如果 iota ⾃增被打断，须显式恢复。
```go
const (
    A = iota        // 0
    B               // 1
    C = "c"         // c
    D               // c，与上⼀⾏相同。
    E = iota        // 4，显式恢复。注意计数包含了 C、D 两⾏。
    F               // 5
)
```

可通过⾃定义类型来实现枚举类型限制。
```go
type Color int

const (
    Black Color = iota
    Red
    Blue
)

func test(c Color) {}

func main() {
    c := Black
    test(c)
    x := 1
    test(x)     // Error: cannot use x (type int) as type Color in function argument
    test(1)     // 常量会被编译器⾃动转换。
}
```

#### 3.基本类型

更明确的数字类型命名，⽀持 Unicode，⽀持常⽤数据结构。

类型 | 长度 | 默认值 | 说明
--- | --- | --- | ---
bool | 1 | false | 
byte | 1 | 0 | uint8
rune | 4 | 0 | Unicode Code Point, int32
int,uint | 4或8 | 0 | 32位或64位
int8,uint8 | 1 | 0 | -128 ~ 127, 0 ~ 255
int16,uint16 | 2 | 0 | -32768 ~ 32767, 0 ~ 65535
int32,uint32 | 4 | 0 | -21亿 ~ 21 亿, 0 ~ 42 亿
int64,uint64 | 8 | 0 | 
float32 | 4 | 0.0 | 
float64 | 8 | 0.0 |
complex64 | 8 |  | 
complex128 | 16 |  |
uintptr | 4或8 |  | ⾜以存储指针的 uint32 或 uint64 整数
array |  |  | 值类型
struct |  |  | 值类型
string |  | "" | UTF-8 字符串
slice |  | nil | 引⽤类型
map |  | nil | 引⽤类型
channel |  | nil | 引⽤类型
interface |  | nil | 接⼝
function |  | nil | 函数

⽀持⼋进制、⼗六进制，以及科学记数法。标准库 math 定义了各数字类型取值范围。
```go
a, b, c, d := 071, 0x1F, 1e9, math.MinInt16
```

空指针值 nil，⽽⾮ C/C++ NULL。


#### 4.引用类型

引⽤类型包括 slice、map 和 channel。它们有复杂的内部结构，除了申请内存外，还需要初始化相关属性。

内置函数`new`计算类型⼤⼩，为其分配零值内存，***返回指针***。⽽`make`会被编译器翻译成具体的创建函数，由其分配内存和初始化成员结构，***返回对象⽽⾮指针***。
```go
a := []int{0, 0, 0}     // 提供初始化表达式。
a[1] = 10

b := make([]int, 3)     // makeslice
b[1] = 10

c := new([]int)
c[1] = 10               // Error: invalid operation: c[1] (index of type *[]int)
```

#### 5.类型转换

不⽀支持隐式类型转换，即便是从窄向宽转换也不⾏行。
```go
var b byte = 100
// var n int = b    // Error Cannot use 'b' (type byte) as the type int
var n int = int(b)  // 显示转换
```

使⽤用括号避免优先级错误。
```go
*Point(p)           // 相当于 *(Point(p)) 
(*Point)(p)
<-chan int(c)       // 相当于 <-(chan int(c)) 
(<-chan int)(c)
```

同样不能将其他类型当 bool 值使⽤用。
```go
a := 100
if a {              // Error: non-bool a (type int) used as if condition
    println("true")
}
```

#### 6.字符串

字符串是不可变值类型，内部⽤用指针指向 UTF-8 字节数组。
- 默认值是空字符串 ""。
- ⽤用索引号访问某字节，如 s[i]。
- 不能⽤用序号获取字节元素指针，&s[i] ⾮非法。 • 不可变类型，⽆无法修改字节数组。
- 字节数组尾部不包含 NULL。

runtime.h
```c
struct String
{
    byte* str;
    intgo len;
}
```

使⽤用索引号访问字符 (byte)。
```go
s := "abc"
println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)    // 输出 true true true
```

使⽤用 "`" 定义不做转义处理的原始字符串，⽀支持跨⾏行。
```go
s := `a
b\r\n\x00
c`
println(s)
```

连接跨⾏行字符串时，"+" 必须在上⼀一⾏行末尾，否则导致编译错误。
```go
s := "Hello, " +
     "World!"
s2 := "Hello, "
    + "World!"    // Error: invalid operation: + untyped string
```

⽀支持⽤用两个索引号返回⼦子串。⼦子串依然指向原字节数组，仅修改了指针和⻓长度属性。
```go
s := "Hello, World!"
s1 := s[:5]     // Hello
s2 := s[7:]     // World!
s3 := s[1:5]    // ello
```

单引号字符常量表⽰示 Unicode Code Point，⽀支持 \uFFFF、\U7FFFFFFF、\xFF 格式。 对应 rune 类型，UCS-4。
```go
func main() { 
	fmt.Printf("%T\n", 'a')                              // 输出 int32  rune是int32的别名
    var c1, c2 rune = '\u6211', '们'
    println(c1 == '我', string(c2) == "\xe4\xbb\xac")    // 输出 true true
}
```


要修改字符串，可先将其转换成 []rune 或 []byte，完成后再转换为 string。⽆无论哪种转 换，都会重新分配内存，并复制字节数组。
```go
func main() {
    s := "abcd"
    bs := []byte(s)
    bs[1] = 'B'
    println(string(bs))         // 输出 aBcd
    
    u := "电脑"
    us := []rune(u)
    us[1] = '话'
    println(string(us))         // 输出 电话
}
```

⽤用 for 循环遍历字符串时，也有 byte 和 rune 两种⽅方式。
```go
func main() {
    s := "abc汉字"
    for i := 0; i < len(s); i++ {           // byte
    	fmt.Printf("%c,", s[i])
    }
    
    for _, r := range s {                   // rune
    	fmt.Printf("%c,", r)
    } 
}
```

#### 7.指针

⽀支持指针类型 *T，指针的指针 **T，以及包含包名前缀的 *<package>.T。
- 默认值 nil，没有 NULL 常量。
- 操作符 "&" 取变量地址，"*" 透过指针访问⺫⽬目标对象。
- 不⽀支持指针运算，不⽀支持 "->" 运算符，直接⽤用 "." 访问⺫⽬目标成员。
```go
func main() {
    type data struct{ a int }
    var d = data{1234}
    var p *data
    p = &d
    fmt.Printf("%p, %v\n", p, p.a)      // 直接⽤用指针访问目标对象成员，无须转换。 
}
```

不能对指针做加减法等运算。
```go
x := 1234
p := &x
p++          // Error: invalid operation: p += 1 (mismatched types *int and int)
```

可以在 unsafe.Pointer 和任意类型指针间进⾏行转换。
```go
func main() {
    x := 0x12345678
    
    p := unsafe.Pointer(&x)             // *int -> Pointer
    n := (*[4]byte)(p)                  // Pointer -> *[4]byte
    
    for i := 0; i < len(n); i++ {
    	fmt.Printf("%X ", n[i])         // 输出 78 56 34 12
    } 
}
```

返回局部变量指针是安全的，编译器会根据需要将其分配在 GC Heap 上。
```go
func test() *int {
    x := 100
    return &x   // 在堆上分配 x 内存。但在内联时，也可能直接分配在目标栈。 
}
```

将 Pointer 转换成 uintptr，可变相实现指针运算。
```go
func main() {
    f := struct {
        s string
        x   int
    }{"abc", 100}
    
    p := uintptr(unsafe.Pointer(&f))		// *struct -> Pointer -> uintptr
    p += unsafe.Offsetof(f.x)				// uintptr + offset
    p2 := unsafe.Pointer(p)					// uintptr -> Pointer
    px := (*int)(p2)						// Pointer -> *int
    *px = 200								// d.x = 200
    
    fmt.Printf("%#v\n", f)			        // 输出 struct { s string; x int }{s:"abc", x:200}
}
```

> 注意:GC 把 uintptr 当成普通整数对象，它⽆无法阻⽌止 "关联" 对象被回收。