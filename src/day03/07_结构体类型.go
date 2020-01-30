package main

import "fmt"

// 结构体的定义
type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//structStudy1()
	structStudy2()
}
func structStudy1() {
	// 1. 初始化
	// 顺序初始化，每个成员必须初始化
	var s1 = Student{1, "Neo", 0, 18, "北京朝阳"}
	fmt.Println("s1 = ", s1)

	// 指定成员初始化，没有初始化的成员自动赋值为零值
	s2 := Student{name: "Neo", addr: "bg"}
	fmt.Println("s2 = ", s2)

	// 直接拷贝是深拷贝，值传递
	s3 := s1
	// (.)点操作符来访问和修改成员
	s3.name = "Neooo"
	fmt.Println("s1 = ", s1)
	fmt.Println("s3 = ", s3)

	// 指针有合法指向后，才操作成员
	p1 := &s1
	(*p1).name = "neooo"
	fmt.Println("s1 = ", s1)
	fmt.Println("*p1 = ", *p1)
	fmt.Println("p1 = ", p1)
	fmt.Printf("p1 type is %T\n", p1) // *main.Student

	// 2. 通过 new 申请一个结构体
	p2 := new(Student)
	p2.id = 1
	p2.name = "Neooo"
	p2.sex = 'm'
	p2.age = 18
	p2.addr = "bj"
	fmt.Println("p2 = ", p2)
}

func structStudy2() {
	// 结构体的比较
	s1 := Student{1, "neo", 'm', 18, "bj"}
	s2 := Student{1, "neo", 'm', 18, "bj"}
	s3 := Student{1, "neoooo", 'm', 18, "bj"}
	fmt.Println("s1 == s2 is ", s1 == s2) // true
	fmt.Println("s1 == s3 is ", s1 == s3) // false
}


