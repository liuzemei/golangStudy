package main

import "fmt"

func main() {
	//sliceStudy()
	//sliceStudy1()
	//sliceStudy2()
	sliceStudy3()
}

func sliceStudy() {
	array := [5]int{1, 2, 3, 4, 0}
	// 切片的定义
	slice := array[1:3:5]
	fmt.Println("slice = ", slice)
	fmt.Println("len(slice) = ", len(slice)) // 长度： 3 - 1
	fmt.Println("cap(slice) = ", cap(slice)) // 容量： 5 - 1
	array[2] = 10
	fmt.Println("slice = ", slice) // 2 10 -> 只是引用，不是值
}

func sliceStudy1() {
	// 切片和数组的区别
	// 数组[]里面的长度是固定的一个常量，不能修改长度，len() == cap()
	// 切片 []里面为空，或者为...，切片的长度或容量可以不固定
	// 切片的定义
	slice := []int{}
	fmt.Printf("slice len = %d, cap = %d\n", len(slice), cap(slice))
	slice = append(slice, 11) // 给切片末尾追加一个成员
	fmt.Println("slice = ", slice)
	fmt.Printf("append: len = %d, cap = %d\n", len(slice), cap(slice))
}

func sliceStudy2() {
	// 切片的创建
	// 1. 自动推导类型，同时初始化
	slice1 := []int{1, 2, 3, 4}
	fmt.Println("slice1 = ", slice1)

	// 2. 借助 make 函数 make(type, len, [cap])
	// 没有指定容量的话，len = cap
	slice2 := make([]int, 3, 5)
	fmt.Println("slice2 = ", slice2)
}

func sliceStudy3() {
	// 切片的操作
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 前面一个不写就是 0，后边一个不写就是 len()
	s1 := array[:] // => [0:len(array)]
	fmt.Println("s1 = ", s1)

	// 操作单个元素，跟数组操作方式一样
	data := s1[1]
	fmt.Println("data = ", data)

	// 内建函数 append
	// 如果超过原来的容量，通常以 2 倍容量扩容
	s1 = append(s1, 2, 2, 3, 4, 5)
	fmt.Println("s1 = ", s1)
	fmt.Println("s1 = ", s1, "lens(s1) = ", len(s1))
	fmt.Printf("s1 = %v, len(s1) = %v, cap(s1) = %v\n", s1, len(s1), cap(s1))
	fmt.Println("array = ", array)

	// 内建函数 copy
	// 将后边的 slice 直接追加到前边的 slice 尾部，不扩容。
	srcSlice := []int{1, 2}
	dstSlice := []int{6, 6, 6, 6, 6, 6}
	copy(dstSlice, srcSlice)
	fmt.Println("dist = ", dstSlice)
	fmt.Printf("dstSlice = %v, len(dstSlice) = %v, cap(dstSlice) = %v\n", dstSlice, len(dstSlice), cap(dstSlice))

}

func sliceStudy4() {

}
