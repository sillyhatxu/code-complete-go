package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var numbers []int
	for len(numbers) < 6 {
		num := rand.Intn(49) + 1
		if !contains(numbers, num) {
			numbers = append(numbers, num)
		}
	}
	fmt.Println(numbers)
}

func contains(numbers []int, num int) bool {
	for _, n := range numbers {
		if n == num {
			return true
		}
	}
	return false
}
