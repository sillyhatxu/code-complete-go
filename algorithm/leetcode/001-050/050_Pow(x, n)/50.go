package _50_Pow_x__n_

func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if n < 0 {
		return 1 / recursion(x, n*-1)
	}
	return recursion(x, n)
}

func recursion(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	half := recursion(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}
