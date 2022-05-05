package main

import "fmt"

const constA string = "这是一个常量"

func main() {
	a := 1
	fmt.Printf("a's value %d\na's pointer %d\n定义的常量 %s\n", a, &a, constA)
	var x int
	var y float64
	//读取键盘的输入，通过操作地址，赋值给x和y   阻塞式
	fmt.Println("请输入一个整数，一个浮点类型：")
	scan, err := fmt.Scanln(&x, &y)
	if err != nil {
		goto handleError
	}
	fmt.Println("你输入了", scan, "个字符")
	fmt.Printf("x的数值：%d，y的数值：%f\n", x, y)
	scan, err = fmt.Scanf("%d next %f", &x, &y)
	if err != nil {
		goto handleError
	}
	fmt.Println("你输入了", scan, "个字符")
	fmt.Printf("x:%d,y:%f\n", x, y)
	return
handleError:
	fmt.Println("error message:", err)
	return

}
