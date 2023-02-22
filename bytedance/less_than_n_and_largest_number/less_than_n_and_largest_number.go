package less_than_n_and_largest_number

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func lessThanNAndLargestNumber1(n int, array []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sort.Ints(array)
	nSrc := strconv.Itoa(n)
	res, num := 0, 0
	var helper func(depth int)
	helper = func(depth int) {
		if depth == len(nSrc) {
			return
		}
		for i := 0; i < len(array); i++ {
			num = num*10 + array[i]
			if num >= n {
				num = (num - array[i]) / 10
				return
			}
			res = max(res, num)
			helper(depth + 1)
			num = (num - array[i]) / 10 //退回
		}
	}
	helper(0)
	return res
}

func lessThanNAndLargestNumber2(n int, array []int) int {
	sort.Ints(array)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	res, num := 0, 0
	var helper func()
	helper = func() {
		for i := 0; i < len(array); i++ {
			num = num*10 + array[i]
			if num > n { // 剪枝，跳到同层下一个元素
				num = (num - array[i]) / 10
				continue // 注意是continue不是return
			}
			res = max(res, num)
			helper()
			num = (num - array[i]) / 10 // 撤销
		}
	}
	helper()
	return res
}

func lessThanNAndLargestNumberInterview(n int, array []int) int {
	sort.Ints(array)
	nSrc := strconv.Itoa(n)
	minimum := getMinimum(array, nSrc)
	if minimum > n {
		//array的最小值仍然大于n，所以从array取最大值但是数字位数小一位 如:12345 [6,7,8,9] return 9999
		return getSameNumber(array[len(array)-1], len(nSrc)-1)
	}
	//2533, []int{1, 2, 4, 9}
	getMax := false
	res := make([]int, len(nSrc), len(nSrc))
	for i := 0; i < len(nSrc); i++ {
		curr := int(nSrc[i] - '0')
		if getMax {
			res[i] = array[len(array)-1]
			continue
		}
		for j := 0; j < len(array); j++ {
			if array[j] <= curr {
				res[i] = array[j]
			} else {
				break
			}
		}
		if res[i] < curr {
			getMax = true
		}
	}
	temp := strings.Trim(strings.Replace(fmt.Sprint(res), " ", "", -1), "[]")
	result, _ := strconv.Atoi(temp)
	return result
}

func getMinimum(array []int, nSrc string) int {
	if array[0] == 0 {
		temp := ""
		for i := 0; i < len(array); i++ {
			if array[i] == 0 {
				continue
			}
			temp = strconv.Itoa(array[i])
			break
		}
		if temp == "" {
			return 0
		}
		for i := 0; i < len(nSrc)-1; i++ {
			temp += "0"
		}
		res, _ := strconv.Atoi(temp)
		return res
	}
	temp := ""
	for i := 0; i < len(nSrc); i++ {
		temp += strconv.Itoa(array[0])
	}
	res, _ := strconv.Atoi(temp)
	return res
}

func getSameNumber(num, digit int) int {
	temp := ""
	for i := 0; i < digit; i++ {
		temp += strconv.Itoa(num)
	}
	res, _ := strconv.Atoi(temp)
	return res
}
