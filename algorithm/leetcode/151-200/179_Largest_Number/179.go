package _79_Largest_Number

import (
	"bytes"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
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
