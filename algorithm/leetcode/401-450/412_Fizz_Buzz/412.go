package _12_Fizz_Buzz

import "fmt"

var (
	fizz = "Fizz"
	buzz = "Buzz"
	fbzz = "FizzBuzz"
)

func fizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		isFizz := (i % 3) == 0
		isBuzz := (i % 5) == 0
		if isFizz && isBuzz {
			res = append(res, fbzz)
		} else if isFizz {
			res = append(res, fizz)
		} else if isBuzz {
			res = append(res, buzz)
		} else {
			res = append(res, fmt.Sprintf("%d", i))
		}
	}
	return res
}
