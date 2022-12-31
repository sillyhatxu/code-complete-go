package _38_Find_All_Anagrams_in_a_String

import "sort"

func findAnagrams(s string, p string) []int {
	dict, validate, size := make([]int, 26, 26), make([]int, 26, 26), len(p)
	compare := func(a, b []int) bool {
		for i := 0; i < 26; i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	var res []int
	for i := 0; i < len(s); i++ {
		dict[s[i]-'a']++
		if i < size {
			validate[p[i]-'a']++
			if i == size-1 && compare(dict, validate) {
				// finish validate
				res = append(res, 0)
			}
		} else {
			dict[s[i-size]-'a']--
			if compare(dict, validate) {
				res = append(res, i-len(p)+1)
			}
		}
	}
	return res
}

func findAnagramsLongTime(s string, p string) []int {
	size := len(p)
	pTemp := []byte(p)
	sort.Slice(pTemp, func(i, j int) bool {
		return pTemp[i] < pTemp[j]
	})
	p = string(pTemp)
	sortS := func(s string, start int) string {
		temp := []byte(s[start : start+size])
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
		return string(temp)
	}
	var res []int
	for i := 0; i < len(s)-size+1; i++ {
		validate := sortS(s, i)
		if validate == p {
			res = append(res, i)
		}
	}
	return res
}
