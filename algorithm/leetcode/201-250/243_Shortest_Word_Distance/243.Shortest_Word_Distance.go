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
func shortestDistanceBasic(wordsDict []string, word1 string, word2 string) int {
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

//O(n)
func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	iIndex := -1
	jIndex := -1
	result := len(wordsDict)
	for i, word := range wordsDict {
		if word == word1 {
			iIndex = i
		} else if word == word2 {
			jIndex = i
		}
		if iIndex != -1 && jIndex != -1 {
			result = min(result, abs(iIndex, jIndex))
		}
	}
	return result
}

func main() {
	input1 := []string{"practice", "makes", "perfect", "coding", "makes"}
	fmt.Println(shortestDistance(input1, "coding", "practice") == 3)
	fmt.Println(shortestDistance(input1, "makes", "coding") == 1)
}
