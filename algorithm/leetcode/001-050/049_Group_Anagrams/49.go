package _49_Group_Anagrams

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	helper := func(s string) string {
		temp := []byte(s)
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
		return string(temp)
	}
	valid := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		s := helper(strs[i])
		if _, ok := valid[s]; ok {
			valid[s] = append(valid[s], strs[i])
		} else {
			valid[s] = []string{strs[i]}
		}
	}
	var res [][]string
	for _, v := range valid {
		res = append(res, v)
	}
	return res
}

func groupAnagramsOriginal(strs []string) [][]string {
	valid := make(map[[26]int][]string)
	for _, str := range strs {
		key := [26]int{}
		for i := 0; i < len(str); i++ {
			key[str[i]-'a']++
		}
		valid[key] = append(valid[key], str)
	}
	var res [][]string
	for _, v := range valid {
		res = append(res, v)
	}
	return res
}

func groupAnagrams1(strs []string) [][]string {
	if len(strs) == 0 || len(strs) == 1 {
		return [][]string{strs}
	}
	temp := make(map[string][]string)
	for _, str := range strs {
		sortSrc := srcSort(str)
		if _, ok := temp[sortSrc]; ok {
			temp[sortSrc] = append(temp[sortSrc], str)
		} else {
			temp[sortSrc] = []string{str}
		}
	}
	var result [][]string
	for _, value := range temp {
		result = append(result, value)
	}
	return result
}

func srcSort(input string) string {
	inputByte := []byte(input)
	sort.Slice(inputByte, func(i, j int) bool {
		return inputByte[i] < inputByte[j]
	})
	return string(inputByte)
}
