package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println(len(Easy))
	fmt.Println(len(Medium))
	fmt.Println(len(Hard))
	fmt.Println(len(Easy) + len(Medium) + len(Hard))
}
func main1() {

	tempEasy := make([]int, len(Easy), len(Easy))
	tempMedium := make([]int, len(Medium), len(Medium))
	copy(tempEasy, Easy)
	copy(tempMedium, Medium)
	Shuffle(tempEasy)
	Shuffle(tempMedium)

	test := func(tag string, input []int, number int) []string {

		var res []string
		for len(input) > 0 {
			temp := fmt.Sprintf("%s:", tag)
			currNumber := min(number, len(input))
			for i := 0; i < currNumber; i++ {
				temp += fmt.Sprintf(" %d,", input[i])
			}
			input = input[currNumber:]
			temp = strings.TrimRight(temp, ",")
			res = append(res, temp)
		}
		return res
	}
	easy := test("Easy", tempEasy, 1)
	medium := test("Medium", tempMedium, 3)
	index := 0
	fmt.Println("```")
	for index < len(easy) && index < len(medium) {
		if index < len(easy) {
			fmt.Print(easy[index], "   ")
		}
		if index < len(medium) {
			fmt.Print(medium[index], "   ")
		}
		fmt.Println("")
		index++
	}
	fmt.Println("```")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Shuffle(slice []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

var Easy = []int{
	1, 9, 13, 14, 20, 21, 26, 27, 35, 66, 69, 70, 88, 94,
	101, 104, 108, 112, 118, 121, 125, 136, 141, 144, 145, 157, 160, 169, 171, 190, 191,
	202, 206, 217, 226, 234, 242, 268, 283,
	326, 344, 350, 387, 392,
	412, 543, 771,
}
var Medium = []int{
	2, 3, 5, 6, 7, 8, 11, 12, 15, 16, 17, 18, 19, 22, 24, 28, 29, 31, 33, 34, 36, 38, 39, 40, 43, 45, 46, 48, 49, 50,
	53, 54, 55, 56, 62, 64, 73, 74, 75, 78, 79, 91, 98,
	102, 103, 105, 106, 114, 116, 122, 128, 130, 131, 134, 138, 139, 142, 146, 148,
	150, 152, 153, 155, 162, 166, 172, 179, 189, 198, 199,
	200, 204, 207, 208, 210, 215, 227, 230, 236, 237, 238, 240,
	251, 253, 277, 279, 285, 287, 289,
	300, 322, 324, 328, 334, 340, 341, 347, 348,
	371, 380, 384, 394, 395,
	416, 438, 454, 560, 739, 763, 994,
}
var Hard = []int{4, 37, 1312}
