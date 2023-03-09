package main

import (
	"fmt"
	"strconv"
)

func main() {
	res1 := test(23121, []int{2, 4, 9})
	fmt.Println(res1 == 22999)
	res2 := test(33121, []int{1, 2, 4, 9})
	fmt.Println(res2 == 29999)
}

//23121.  [2, 4, 9] 22999
//33121.  [1, 2, 4, 9]. 29999
//33121.  [4, 9]. 9999. 44444
func test(input int, arr []int) int {
	inputSrc := strconv.Itoa(input)
	getMax := false
	res := make([]int, len(inputSrc), len(inputSrc))
	for i := 0; i < len(inputSrc); i++ {
		current := int(inputSrc[i] - '0')
		if getMax {
			res[i] = arr[len(arr)-1]
			continue
		}
		for j := 0; j < len(arr); j++ {
			if arr[j] <= current {
				res[i] = arr[j]
			} else {
				break
			}
		}
		if res[i] < current {
			getMax = true
		}
	}
	temp := ""
	for i := 0; i < len(res); i++ {
		temp += strconv.Itoa(res[i])
	}
	result, _ := strconv.Atoi(temp)
	return result
}
