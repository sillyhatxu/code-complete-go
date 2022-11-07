package _91_Decode_Ways

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	//dp[i] = dp[i-1] + dp[i-2]
	current, previous := 1, 1 //dp[i], dp[i-1]
	for i := 1; i < len(s); i++ {
		num := getRealNum(s[i-1])*10 + getRealNum(s[i])
		if num >= 10 && num <= 26 {
			//[10-26]
			if s[i] == '0' {
				current, previous = previous, 0
			} else {
				previous, current = current, current+previous
			}
		} else {
			//(0-9]
			if s[i] == '0' {
				return 0
			} else {
				previous = current
			}
		}
	}
	return current
}

func getRealNum(input byte) byte {
	return input - '0'
}
