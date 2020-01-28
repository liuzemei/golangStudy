package main

import "fmt"

func main() {
	//mapStudy1()
	mapStudy2()
}

func mapStudy1() {
	// 定义一个变量，类型为 map[int]string
	var m1 map[int]string
	fmt.Println("m1 = ", m1) // map[]

	// 对于 map 只有 len，没有 cap
	fmt.Println("len(m1) = ", len(m1)) // 0

	// 可以通过 make 创建
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)           // map[]
	fmt.Println("len(m2) = ", len(m2)) // 0

	// 通过 make 创建，可以指定长度，只是指定了容量，但是里面确实一个数据也没有
	// 会自动扩容，不过提前分配一些空间会有利于性能提升。
	m3 := make(map[int]string, 10)
	fmt.Println("m3 = ", m3)           // map[]
	fmt.Println("len(m3) = ", len(m3)) // 0

	// 初始化 键值是唯一的
	m4 := map[int]string{1: "mike", 2: "go", 3: "c++"}
	fmt.Println("m4 = ", m4) // map[]
}

func mapStudy2() {
	m1 := map[int]string{1: "mike", 2: "go", 3: "c++"}
	for key, value := range m1 {
		fmt.Printf("m1[%d] = %s\n", key, value)
	}

	// 如何判断一个 key 值是否存在
	value, ok := m1[5]
	if ok == true {
		fmt.Println("m1[1] = ", value)
	} else {
		fmt.Println("key 不存在")
	}

	// 删除 map 中的内容
	delete(m1, 1) // 删除 key 为 1 的内容
}

func mapStudy3() {

}
