package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(test, len(test), cap(test))
	temp := test[2:4]
	fmt.Println(temp, len(temp), cap(temp))
	fmt.Println(temp[0])
	fmt.Println(temp[1])
	fmt.Println(temp[0:6])
	fmt.Println(temp[2])
}
