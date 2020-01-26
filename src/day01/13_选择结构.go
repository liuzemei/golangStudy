package main

import "fmt"

func main() {

	// 1. 单 if
	s := "王思聪"
	// if 和 { 之间的就是条件，条件通常都是关系运算符
	if s == "王思聪" {
		fmt.Println("左手一个妹子，右手一个大妈")
	}
	// if 支持 1 个初始化语句,初始化语句和判断条件以分好分割
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	}

	// 2. if else
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else {
		fmt.Println("a < 10")
	}

	// 3. switch
	// 1. 基本用法
	num := 1
	switch num { // switch 后面写的是变量本身
	case 1:
		fmt.Printf("按下的是1楼\n")
		//break // go 语言保留了 break 关键字，跳出 switch 语句，不写，默认就包含
		//fallthrough 不跳出 switch，后边的无条件执行
	case 2:
		fmt.Printf("按下的是2楼\n")
	case 3:
		fmt.Printf("按下的是3楼\n")
	default:
		fmt.Println("按下的是 xxx 楼")
	}

	// 2. switch 也支持一个初始化语句，初始化语句和变量本身，以分号分割
	// 也可以放多个常量
	switch num1 := 1; num1 {
	case 1:
		fmt.Printf("按下的是1楼\n")
	case 2:
		fmt.Printf("按下的是2楼\n")
	case 3, 4, 5: // 也可以放多个常量
		fmt.Printf("按下的是3楼\n")
	default:
		fmt.Println("按下的是 xxx 楼")
	}

	// 3. 可以没有条件 case 后边可以放条件
	score := 85
	switch {
	case score > 90:
		fmt.Println("优秀")
	case score > 80:
		fmt.Println("良好")
	case score > 70:
		fmt.Println("中等")
	case score > 60:
		fmt.Println("及格")
	}

}
