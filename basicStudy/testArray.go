package main

import "fmt"

func main() {
	var intArr [10]int8
	//intArr = [10]int8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//intArr := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	intArr = [10]int8{1: 2, 4: 10, 8: 38}
	fmt.Println(intArr)
	fmt.Println(len(intArr))
	for index, value := range intArr {
		fmt.Printf("Index:%d\tValue:%d\tElement[%d]:%d\n", index, value, index, intArr[index])
	}
}
