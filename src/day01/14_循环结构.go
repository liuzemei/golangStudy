package main

import "fmt"

func main() {
	// for 初始化条件;判断条件;条件变化{
	// }
	// 1 ~ 100 累加
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// range
	str := "abc"

	// 通过 for 打印每个字符
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
	// 迭代打印每个元素，默认返回 2 个值：一个是元素的位置，一个是元素本身
	for i, data := range str {
		fmt.Printf("str[%d]=%c\n", i, data)
	}

	for i := range str { // 第 2 个返回值，默认丢弃，返回元素的下标
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
}
