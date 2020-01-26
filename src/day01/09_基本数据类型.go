package main

import "fmt"

func main() {
	// 1. bool 零值为 false
	var a bool
	fmt.Println("a = ", a)

	// 2. 浮点型
	var f1 float32
	f1 = 3.14
	fmt.Println("f1 = ", f1)

	// 自动推导类型为 float64
	f2 := 3.14
	fmt.Printf("f2 type is %T\n", f2)

	// 3. 字符类型
	var ch byte
	ch = 97
	// 格式化输出， %c以字符方式打印，%d 以整型方式打印
	fmt.Printf("ch = %c\n", ch)

	ch = 'a'
	fmt.Printf("ch = %d\n", ch)

	// 大小写转换，相差 32
	fmt.Printf("大写: %d, 小写: %d\n", 'A', 'a')
	fmt.Printf("大写转小写： %c\n", 'A'+32)
	fmt.Printf("小写转大写： %c\n", 'a'-32)

	// 4. 字符串类型
	// 1. 双引号
	// 2. 字符串有 1 个或多个字符组成
	// 3. 字符串都是隐藏了一个结束符， '\0' 的
	var str1 string
	str1 = "abc"
	fmt.Println("str1 = ", str1)

	// 内建函数 len()可以测字符串的长度
	fmt.Println("len(str1) = ", len(str1))

	// 5. 复数类型
	var t complex128
	t = 2.1 + 3.14i
	fmt.Println("t = ", t)
	t2 := 3.3 + 3.14i
	fmt.Printf("t2 type is %T\n", t2)
	fmt.Printf("real(t2) = %f, imag(t2) = %f", real(t2), imag(t2))
}
