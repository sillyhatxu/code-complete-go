package main

import (
	"fmt"
)

func stringsRearrangement(inputArray []string) bool {
	for perm := range permutations(inputArray) {
		if oneLetterDiff(perm) {
			return true
		}
	}
	return false
}

func oneLetterDiff(perm []string) bool {
	last := len(perm) - 1
	for i := 0; i < last; i++ {
		if perm[i] == perm[i+1] {
			return false
		}
		mistakeNum := 1
		for j := range perm[i] {
			if perm[i][j] != perm[i+1][j] {
				mistakeNum--
				if mistakeNum < 0 {
					return false
				}
			}
		}
	}
	return true
}

func permutations(data []string) <-chan []string {
	c := make(chan []string)
	go func(c chan []string) {
		defer close(c)
		permutate(c, data)
	}(c)
	return c
}

func permutate(c chan []string, inputs []string) {
	output := make([]string, len(inputs))
	copy(output, inputs)
	c <- output
	size := len(inputs)
	p := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		p[i] = i
	}
	for i := 1; i < size; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		tmp := inputs[j]
		inputs[j] = inputs[i]
		inputs[i] = tmp
		output := make([]string, len(inputs))
		copy(output, inputs)
		c <- output
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

func main() {
	fmt.Println(stringsRearrangement([]string{"abc", "abx", "axx", "abc"}) == false)
	fmt.Println(stringsRearrangement([]string{"q", "q"}) == false)
	fmt.Println(stringsRearrangement([]string{"aba", "bbb", "bab"}) == false)
	fmt.Println(stringsRearrangement([]string{"ab", "bb", "aa"}))
	fmt.Println(stringsRearrangement([]string{"zzzzab", "zzzzbb", "zzzzaa"}))
	fmt.Println(stringsRearrangement([]string{"ab", "ad", "ef", "eg"}) == false)
	fmt.Println(stringsRearrangement([]string{"abc", "bef", "bcc", "bec", "bbc", "bdc"}))
	fmt.Println(stringsRearrangement([]string{"abc", "abx", "axx", "abx", "abc"}))
	fmt.Println(stringsRearrangement([]string{"f", "g", "a", "h"}))
	fmt.Println(stringsRearrangement([]string{"ff", "gf", "af", "ar", "hf"}))
	fmt.Println(stringsRearrangement([]string{"a", "b", "c"}))
}
