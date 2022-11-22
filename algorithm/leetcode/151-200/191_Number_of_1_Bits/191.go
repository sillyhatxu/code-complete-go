package _91_Number_of_1_Bits

func hammingWeight(num uint32) int {
	res := 0
	for i := 0; i < 32; i++ {
		res += int(num & 1)
		num = num >> 1
	}
	return res
}
