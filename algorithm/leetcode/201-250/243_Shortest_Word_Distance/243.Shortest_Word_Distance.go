package main

import (
	"fmt"
)

func abs(i, j int) int {
	if i > j {
		return i - j
	} else {
		return j - i
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// O(n)2
func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	result := len(wordsDict)
	for i, iWord := range wordsDict {
		if iWord == word1 {
			for j, jWord := range wordsDict {
				if jWord == word2 {
					result = min(result, abs(i, j))
				}
			}
		}
	}
	return result
}

func main() {
	input1 := []string{"practice", "makes", "perfect", "coding", "makes"}
	fmt.Println(shortestDistance(input1, "coding", "practice") == 3)
	fmt.Println(shortestDistance(input1, "makes", "coding") == 1)
}
