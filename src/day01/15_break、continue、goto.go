package main

import "fmt"

func main() {
	//i := 0
	//for { // for 后面不写任何东西，这个循环条件死循环
	//	i++
	//	time.Sleep(time.Second)
	//	if i == 5 {
	//		//break // 跳出循环，如果嵌套多个循环，跳出最内层的循环
	//		continue
	//	}
	//	fmt.Println("i = ", i)
	//}

	// break 只能在 loop switch select 中使用
	// continue 只能在 loop 中使用
	// goto 可以用在任何地方，但是不能跨函数使用
	fmt.Println("111111111111111")
	goto End
	fmt.Println("222222222222222")
End:
	fmt.Println("333333333333333")
}
