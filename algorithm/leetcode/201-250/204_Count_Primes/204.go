package _04_Count_Primes

import "math"

func countPrimes(n int) int {
	if n < 3 {
		//The first Prime number is 2
		return 0
	}
	nums := make([]byte, n, n)
	nums[0] = '1'
	nums[1] = '1'
	sqrt := math.Sqrt(float64(n))
	for i := 2; i <= int(sqrt); i++ {
		if nums[i] == '1' {
			continue
		}
		for j := i * i; j < n; j += i {
			nums[j] = '1'
		}
	}
	//get all prime numbers
	res := 1
	for i := 3; i < n; i++ {
		if nums[i] != '1' {
			res++
		}
	}
	return res
}
