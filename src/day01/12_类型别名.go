package main

import "fmt"

func main() {
	// 给 int64 起一个别名
	type bigint int64
	var a bigint
	fmt.Printf("a type is %T\n", a)

	// 批量起别名
	type (
		long int64
		char byte
	)

	var b long = 11
	var ch char = 'a'
	fmt.Printf("b type is %T, b is %d\n", b, b)
	fmt.Printf("ch type is %T, ch is %c\n", ch, ch)

}
