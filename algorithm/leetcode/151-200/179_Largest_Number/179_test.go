package _79_Largest_Number

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

/**
系统对test1和test2进行逐字节比较，也就是先比较
test1[0] = 3
test2[0] = 3
结果相等，继续往后判断
test1[1] = 4
test2[1] = 0
判断test1[1] > test2[1], 返回结果，不用继续比较
*/

func Test_compare(t *testing.T) {
	test1 := "34"
	test2 := "301"
	fmt.Println(test1 > test2)
	fmt.Println([]byte(test1))
	fmt.Println([]byte(test2))
	fmt.Println(strings.Compare(test1, test2))
	fmt.Println("303" > "330")
}

func Test_largestNumber(t *testing.T) {
	assert.EqualValues(t, "9534330", largestNumber([]int{3, 30, 34, 5, 9}))
}
