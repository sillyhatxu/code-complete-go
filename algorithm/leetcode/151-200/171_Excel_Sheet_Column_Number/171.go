package _71_Excel_Sheet_Column_Number

func titleToNumber(columnTitle string) int {
	res := 0
	for _, c := range columnTitle {
		res = res*26 + int(c-'A'+1)
	}
	return res
}

func titleToNumber1(columnTitle string) int {
	result := 0
	x := len(columnTitle) - 1
	for i := 0; i < len(columnTitle); i++ {
		if x == 0 {
			result += int(columnTitle[i] - 'A' + 1)
			continue
		} else {
			temp := 26
			for j := 1; j < x; j++ {
				temp *= 26
			}
			result += int(columnTitle[i]-'A'+1) * temp
			x--
		}
	}
	return result
}

func titleToNumberSample(columnTitle string) int {
	result := 0
	for i := 0; i < len(columnTitle); i++ {
		result = result*26 + int(columnTitle[i]-'A'+1)
	}
	return result
}
