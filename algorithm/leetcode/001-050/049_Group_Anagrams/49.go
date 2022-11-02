package _49_Group_Anagrams

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	tempMap := map[[26]int][]string{}
	for _, str := range strs {
		key := [26]int{}
		for i := 0; i < len(str); i++ {
			key[str[i]-'a'] += 1
		}
		tempMap[key] = append(tempMap[key], str)
	}
	var result [][]string
	for _, value := range tempMap {
		result = append(result, value)
	}
	return result
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
