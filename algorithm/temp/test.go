package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1
	f := func() int {
		//defer func() {a--}()
		a++
		fmt.Printf("aaaaaaaa : %d \n", a)
		return a
	}

	t := func() int {
		//defer func() {a--}()
		a++
		time.Sleep(5 * time.Second)
		fmt.Printf("ttttttttt : %d \n", a)
		return a
	}

	m := map[int]int{t(): f()}
	fmt.Println(m)
	a++
	fmt.Println(m)
}
