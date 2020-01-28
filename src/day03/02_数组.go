package main

import "fmt"

func main() {
	//arrayStudy()
	//arrayStudy1()
	//arrayStudy2()
	//arrayStudy3()
	arrayStudy4()

}

func arrayStudy() {
	var id [50]int
	//for i := 0; i < len(id); i++ {
	//	id[i] = i + 1
	//	fmt.Printf("id[%d] = %d\n", i, id[i])
	//}
	for i, data := range id {
		// 这个 data 是内存空间的引用，不是真正的 id[i]
		id[i] = i + 1
		fmt.Printf("id[%d] = %d\n", i, data)
	}
	fmt.Printf("id[50] = %v\n", id)
}
func arrayStudy1() {
	// 定义一个数组，[10]int 和 [5]int 是不同类型
	// [数字]，这个数字作为数组元素个数
	// 1. 全部初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	// 注意：指定的数组元素个数必须是常量
	fmt.Printf("%v\n", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", b)

	// 2. 部分初始化，没有初始化的元素自动赋值为 0 值

	// 3. 指定某个元素初始化
	d := [5]int{2: 10, 4: 20}
	fmt.Printf("d = %v\n", d)
}

func arrayStudy2() {
	// 二维数组
	var a [3][4]int
	for i := range a {
		for j := range a[i] {
			a[i][j] = (i)*10 + j
		}
	}
	fmt.Printf("a =  %v \n", a)
	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("b = ", b)
}

func arrayStudy3() {
	// 数组支持比较，只支持 == 或 != ，比较是不是每一个元素都一样，2 个数组比较，类型要一样
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3, 5, 4}
	fmt.Printf(" a == b is %v\n", a == b) // true
	fmt.Printf(" a == c is %v\n", a == c) // false
}

func arrayStudy4() {
	a := [5]int{1, 2, 3, 4, 5}
	p := &a
	p0 := &a[0]
	p1 := &a[1]
	//fmt.Println("p", p)
	fmt.Printf(" &a = %p\n", p)
	fmt.Printf(" &a[0] = %p\n", p0)
	fmt.Printf(" &a1 = %p\n", p1)
}
