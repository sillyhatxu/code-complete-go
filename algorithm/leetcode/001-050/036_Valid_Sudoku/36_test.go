package main

import "testing"

func Test_isValidSudoku(t *testing.T) {
	var input1 [][]byte
	input1 = append(input1, []byte("0"))
	isValidSudoku(input1)
}
