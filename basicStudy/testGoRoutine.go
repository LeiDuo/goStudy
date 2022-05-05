package main

import (
	"fmt"
	"time"
)

func mygo(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("In goroutine %s\t i=%d\n", name, i)
		// 为了避免第一个协程执行过快，观察不到并发的效果，加个休眠
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go mygo("协程1号") // 第一个协程
	go mygo("协程2号") // 第二个协程
	time.Sleep(time.Second)
}
