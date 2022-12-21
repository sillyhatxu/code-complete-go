package main

func romanToInt(s string) int {
	h := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res, max := 0, len(s)-1
	for i := 0; i <= max; i++ {
		temp := h[s[i]]
		if i < max && s[i] == 'I' && s[i+1] == 'V' {
			temp = 4
			i++
		} else if i < max && s[i] == 'I' && s[i+1] == 'X' {
			temp = 9
			i++
		} else if i < max && s[i] == 'X' && s[i+1] == 'L' {
			temp = 40
			i++
		} else if i < max && s[i] == 'X' && s[i+1] == 'C' {
			temp = 90
			i++
		} else if i < max && s[i] == 'C' && s[i+1] == 'D' {
			temp = 400
			i++
		} else if i < max && s[i] == 'C' && s[i+1] == 'M' {
			temp = 900
			i++
		}
		res += temp
	}
	return res
}

/**
Runtime: 16 ms, faster than 100.00% of Go online submissions for Roman to Integer.
Memory Usage: 3 MB, less than 22.00% of Go online submissions for Roman to Integer.
*/
func romanToIntOriginal(s string) int {
	h := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		current := h[s[i]]
		next := 0
		if i+1 < len(s) {
			next = h[s[i+1]]
		}
		if current >= next {
			result += current
		} else {
			result += next - current
			i++
		}
	}
	return result
}
