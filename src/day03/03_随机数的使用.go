package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置种子，只需一次
	rand.Seed(time.Now().UnixNano())
	// 产生随机数,注意，如果种子一样，那么每次生成的数都一样。
	for i := 0; i < 50; i++ {
		fmt.Println("rand = ", rand.Intn(100)) // 0~99
	}
}
