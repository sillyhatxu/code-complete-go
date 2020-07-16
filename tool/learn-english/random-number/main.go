package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Single() string {
	return Output(1, "", 0)
}

func PhotoNumber(digits int) string {
	return Output(9, " ", 3)
}

func Money(digits int) string {
	result := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < digits; i++ {
		result += strconv.Itoa(rand.Intn(10))
	}
	return Output(6, ",", 3)
}

func Output(digits int, delimiter string, delimiterNumber int) string {
	result := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < digits; i++ {
		if i != 0 && delimiterNumber > 0 {
			if i%delimiterNumber == 0 {
				result = delimiter + result
			}
		}
		result = strconv.Itoa(rand.Intn(10)) + result
	}
	return result
}

func main() {
	number := 50
	for i := 0; i < number; i++ {
		fmt.Println(Single())
		time.Sleep(500 * time.Millisecond)
	}
	for i := 0; i < number; i++ {
		fmt.Println(PhotoNumber(9))
		time.Sleep(3 * time.Second)
	}
	for i := 0; i < number; i++ {
		fmt.Println(Money(6))
		time.Sleep(3 * time.Second)
	}
}
