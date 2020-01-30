package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  byte
}

type Student struct {
	// 通过匿名字段 实现继承
	Person
	id   int
	addr string
	name string
}

func main() {
	//AnonymousFields()
	//AnonymousFields1()
	AnonymousFields2()
}

func AnonymousFields() {
	// 1. 顺序初始化
	s1 := Student{Person{"Neo", 18, 'm'}, 1, "bg", "haha"}
	fmt.Println("s1 = ", s1)
	s2 := new(Student)
	s2.name = "haha"
	fmt.Println("s2 = ", s2)
	// %+v ，显示更详细
	fmt.Printf("s2 = %+v\n", s2)

	// 2. 部分初始化, 其他的就是零值
	s3 := Student{id: 1}
	fmt.Println("s3 = ", s3)
}

func AnonymousFields1() {
	// 重名变量先访问外层的
	s1 := Student{Person{"Neo", 18, 'm'}, 1, "bg", "haha"}
	fmt.Println(s1.name, s1.sex, s1.age, s1.addr)
	// 如果要访问内层的，需要加上结构体名称
	fmt.Println(s1.Person.name, s1.sex, s1.age, s1.addr)
}

type Student1 struct {
	*Person // 指针类型
	id   int
	addr string
}

func AnonymousFields2() {
	// 1. 顺序初始化 初始化指针类型的成员，就直接取地址就 ok
	s1 := Student1{&Person{"yoyo", 18, 'm'}, 1, "bj"}
	fmt.Println(s1, s1.name)
	var s2 Student1
	s2.Person = new(Person) // 分配空间
	s2.name = "neo"
	fmt.Println(s2, s2.name)
}
