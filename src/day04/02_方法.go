package main

import "fmt"

func main() {
	//MethodStudy1()
	//MethodStudy2()
	//MethodStudy3()
	//MethodStudy4()
	MethodStudy5()
}

func MethodStudy1() {
	var i long
	i = 1
	var j long
	j = 10
	// 调用方式： 变量名.函数名(所需参数)
	// 注意参数和返回值类型需要完全匹配。
	s := i.Add01(j)
	fmt.Println("s = ", s)
}

type long int

// 给某一个类型绑定一个函数
func (tmp long) Add01(other long) long {
	return tmp + other
}

// 方法的定义和调用
func MethodStudy2() {
	p := Person{"Neo", 18, 'm'}
	p.PrintInfo()

	var p2 Person
	// 这两种调用方式都可以
	//(&p2).SetInfo("neooo", 20, 'm')
	p2.SetInfo("neooo", 20, 'm')
	p2.PrintInfo()
}

type Person struct {
	name string
	age  int
	sex  byte
}

func (tmp Person) PrintInfo() {
	fmt.Println("tmp = ", tmp)
}

// 注意，结构体是值传递，如果要更改本身的内容，需要使用指针来完成
func (tmp *Person) SetInfo(name string, age int, sex byte) {
	tmp.name = name
	tmp.age = age
	tmp.sex = sex
}

// 指针变量的方法集
func MethodStudy3() {
	// 结构体变量是一个指针变量，它能够调用的方法就是它的方法集。
	p := &Person{"mike", 18, 'm'}

	// 指针变量和值变量都可以调用，内部会自动转换为值或者指针。
	p.SetInfo("neo", 10, 'f')
	p.PrintInfo()
}

// 方法的继承
// Person 类型，实现了 2 个方法
type Student struct {
	Person // 匿名字段
	id   int
	addr string
}

//  方法的重写（同名方法）
func (tmp Student) PrintInfo() {
	fmt.Println("tmp = ", tmp)
}
func MethodStudy4() {
	s := Student{Person{"test", 18, 'm'}, 1, "bj"}
	s.Person.PrintInfo()
	s.PrintInfo()
}

func MethodStudy5() {
	// 换一种方式调用：方法值和方法表达式
	p := Person{"neo", 18, 'm'}
	// 方法值
	f := p.PrintInfo
	f()
	// 方法表达式
	f1 := (*Person).PrintInfo
	f1(&p) // 显式的把接受者传递过去
}
