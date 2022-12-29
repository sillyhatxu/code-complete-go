package _94_Decode_String

import "bytes"

func decodeString(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			res += s[i : i+1]
		} else {
			temp, num, c := "", 0, s[i]-'0'
			for c <= 9 {
				num = num*10 + int(c)
				i++
				c = s[i] - '0'
			}
			if s[i] == '[' {
				index, curr := i, 0
				for index < len(s) {
					if s[index] == '[' {
						curr++
					}
					if s[index] == ']' {
						curr--
					}
					if curr == 0 {
						break
					}
					index++
				}
				temp = decodeString(s[i+1 : index])
				i = index
			}
			for j := 0; j < num; j++ {
				res += temp
			}
		}
	}
	return res
}

func decodeStringStack(s string) string {
	var nums []int
	var dicks []string
	for i := 0; i < len(s); i++ {
		c := s[i] - '0'
		if c <= 9 {
			//number
			num := int(c)
			for {
				next := s[i+1] - '0'
				if next > 9 {
					break
				}
				num = num*10 + int(next)
				i++
			}
			nums = append(nums, num)
		} else if s[i] == ']' {
			index, temp := len(dicks)-1, ""
			for dicks[index] != "[" {
				temp = dicks[index] + temp
				index--
			}
			dicks = dicks[:index]
			curr, times := "", nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			for j := 0; j < times; j++ {
				curr += temp
			}
			dicks = append(dicks, curr)
		} else {
			//letter or '['
			dicks = append(dicks, string(s[i]))
		}
	}
	var res bytes.Buffer
	for i := 0; i < len(dicks); i++ {
		res.WriteString(dicks[i])
	}
	return res.String()
}
