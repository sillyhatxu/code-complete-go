package _79_Largest_Number

import (
	"bytes"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	max := len(nums)
	temp := make([]string, max, max)
	for i := 0; i < max; i++ {
		temp[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i]+temp[j] > temp[j]+temp[i]
	})
	if temp[0] == "0" {
		return "0"
	}
	res := ""
	for i := 0; i < max; i++ {
		res += temp[i]
	}
	return res
}

func largestNumber1(nums []int) string {
	var numsSrc []string
	for _, num := range nums {
		numsSrc = append(numsSrc, strconv.Itoa(num))
	}
	sort.Slice(numsSrc, func(i, j int) bool {
		a, b := numsSrc[i]+numsSrc[j], numsSrc[j]+numsSrc[i]
		return a > b
	})
	if numsSrc[0] == "0" {
		return "0"
	}
	var res bytes.Buffer
	for i := range numsSrc {
		res.WriteString(numsSrc[i])
	}
	return res.String()
}
