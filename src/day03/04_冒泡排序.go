package main

import "fmt"

func main() {
	a := []int{1,23,412,34,231,423,4,12,4,32,14,23,41,234,21,4,23,14,1234}
	b := sort(a)
	fmt.Println("b = ", b)
}

func sort(arr []int) (result []int) {
	switch len(arr) {
	case 0, 1:
	case 2:
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
	default:
		for i, len := 0, len(arr); i < len-2; i++ {
			for j := 0; j < len-1-i; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}
	result = arr
	return
}
