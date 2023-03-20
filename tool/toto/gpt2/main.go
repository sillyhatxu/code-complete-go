package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var pastResults = [][]int{
	{4, 14, 30, 33, 37, 42},
	{1, 3, 15, 18, 23, 48},
	{1, 10, 30, 31, 38, 45},
	//{21, 30, 35, 36, 37, 45},
	//{15, 22, 24, 26, 29, 36},
	//{6, 26, 32, 39, 44, 48},
	//{2, 5, 19, 24, 34, 44},
}

//var pastResults = [][]int{
//	{3, 5, 10, 24, 31, 48},
//	{9, 18, 19, 23, 32, 36},
//	{1, 10, 12, 26, 27, 35},
//	{8, 12, 16, 18, 22, 40},
//	{2, 6, 16, 17, 29, 42},
//	{3, 5, 14, 16, 21, 33},
//	{4, 11, 17, 28, 29, 41},
//}

func main() {
	rand.Seed(time.Now().UnixNano())
	freq := getFrequency(pastResults)
	var numbers []int
	for len(numbers) < 6 {
		num := generateNumber(freq)
		if !contains(numbers, num) {
			numbers = append(numbers, num)
		}
	}
	sort.Ints(numbers)
	fmt.Println(numbers)
}

func getFrequency(results [][]int) map[int]map[int]int {
	freq := make(map[int]map[int]int)
	for i := 1; i <= 49; i++ {
		freq[i] = make(map[int]int)
	}
	for _, res := range results {
		for i, num := range res {
			freq[num][i]++
		}
	}
	return freq
}

func generateNumber(freq map[int]map[int]int) int {
	var num int
	maxCount := 0
	for i := 0; i < 6; i++ {
		for n, count := range freq {
			if count[i] > maxCount || (count[i] == maxCount && rand.Intn(2) == 0) {
				maxCount = count[i]
				num = n
			}
		}
	}
	return num
}

func contains(numbers []int, num int) bool {
	for _, n := range numbers {
		if n == num {
			return true
		}
	}
	return false
}
