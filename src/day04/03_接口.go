package main

import "fmt"

func main() {
	//interfaceStudy1()
	//interfaceStudy2()
	//interfaceStudy3()
	interfaceStudy4()
}

// 1. 接口类型的定义
type Humaner interface {
	// 方法，只有声明，没有实现
	sayhi()
}

type Student struct {
	name string
	id   int
}

type Teacher struct {
	name  string
	group string
}
type MyStr string

func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi \n", tmp.name, tmp.id)
}

func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher[%s, %s] sayhi \n", tmp.name, tmp.group)
}

func (tmp *MyStr) sayhi() {
	fmt.Printf("MyStr[%s] sayhi \n", *tmp)
}

func interfaceStudy1() {
	// 1. 定义接口类型变量
	var i Humaner
	// 只要实现了此接口方法的类型，那么这个类型的变量就可以给 i 赋值
	s := &Student{"neo", 18}
	i = s
	i.sayhi()

	t := &Teacher{"bj", "go"}
	i = t
	i.sayhi()

	var str MyStr
	str = "hello"
	i = &str
	i.sayhi()
}

// 2. 接口类型的使用
// 定义一个普通函数，函数的参数为接口类型
func WhoSayHi(i Humaner) {
	i.sayhi()
}

func interfaceStudy2() {
	s := &Student{"neo", 18}
	t := &Teacher{"bj", "go"}
	var str MyStr
	// 只有一个函数，可以有不同表现  =====>  多态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	// 创建一个切片（接口类型）
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str
	for i, data := range x {
		WhoSayHi(x[i])
		data.sayhi()
	}
}

// 3. 接口的嵌入和实现
type Personer interface {
	Humaner //匿名字段，继承了 sayhi()
	sing(lrc string)
}

func (tmp *Student) sing(lrc string) {
	fmt.Println("Student 在唱着：", lrc)
}

func interfaceStudy3() {
	// 定义一个接口类型的变量
	var i Personer
	s := &Student{"neo", 666}
	i = s
	i.sayhi()
	// 超集可以转换为子集，反过来不行
	i.sing("十年")
}

// 可以接受任意多个任意类型的参数
func xxx(args ...interface{}){
}

func interfaceStudy4() {
	// 空接口万能类型，保存任意类型的值
	var i interface{} = 1
	fmt.Println("i = ", i)
	i = "hello"
	fmt.Println("i = ", i)
}
