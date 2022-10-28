package main

import "fmt"

type Error struct {
	errCode uint8
}

func (e *Error) Error() string {
	switch e.errCode {
	case 1:
		return "file not found"
	case 2:
		return "time out"
	case 3:
		return "permission denied"
	default:
		return "unknown error"
	}
}

func checkError(err error) {
	fmt.Println("----------")
	fmt.Println(err)
	fmt.Println(err != nil)
	fmt.Println("----------")
	if err != nil {
		panic(err)
	}
}
func main() {
	var v interface{}
	fmt.Println("----------")
	fmt.Println(v)
	fmt.Println(v == nil)
	fmt.Println("----------")
	var e *Error
	fmt.Println(e)
	fmt.Println(e != nil)
	if e != nil {
		fmt.Println(123321)
		panic(e)
	}
	checkError(e)
}
