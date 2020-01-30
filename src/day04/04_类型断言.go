package main

import "fmt"

// 可以接受任意多个任意类型的参数

type Student struct {
	name string
	id   int
}

func main() {
	//typeAsset1()
	typeAsset2()
}

// 通过 if 实现
func typeAsset1() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello go"
	i[2] = Student{"Neo", 666}

	// 类型查询，类型断言
	for index, data := range i {
		// 第一个返回值，第二个返回判断结果的真假
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d]类型为 int，内容为%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d]类型为 string，内容为%s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d]类型为 Student，内容为%v\n", index, value)
		}
	}
}

// 通过 switch 实现
func typeAsset2() {

	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello go"
	i[2] = Student{"Neo", 666}
	// 类型查询，类型断言
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d]类型为 int，内容为%d\n", index, value)
		case string:
			fmt.Printf("x[%d]类型为 string，内容为%s\n", index, value)
		case Student:
			fmt.Printf("x[%d]类型为 Student，内容为%v\n", index, value)
		}
	}
}
