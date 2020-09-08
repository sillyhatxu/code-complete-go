package main

import (
	"fmt"
	"github.com/sillyhatxu/code-complete-go/algorithm/word_frequency/statistics"
	"path/filepath"
)

func main() {
	inputFilePath, _ := filepath.Abs("../algorithm-go/word_frequency/files/article.txt")
	outputFilePath, _ := filepath.Abs("../algorithm-go/word_frequency/files/result.txt")
	fmt.Println(inputFilePath)
	fmt.Println(outputFilePath)
	statistics.CountTestBase(inputFilePath, outputFilePath)
}
