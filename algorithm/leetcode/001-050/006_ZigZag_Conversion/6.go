package main

import (
	"bytes"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	temp := make([]string, numRows, numRows)
	row, up := 0, false
	for i := range s {
		temp[row] += string(s[i])
		if row == numRows-1 {
			up = true
		} else if row == 0 {
			up = false
		}
		if up {
			row--
		} else {
			row++
		}
	}
	return strings.Join(temp, "")
}

func convert1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	temp := make([]string, numRows)
	row, up := 0, true
	for i := range s {
		temp[row] += string(s[i])
		if row == numRows-1 {
			up = false
		} else if row == 0 {
			up = true
		}
		if up {
			row++
		} else {
			row--
		}
	}
	return strings.Join(temp, "")
}

func convertHistory2(s string, numRows int) string {
	temp := make([][]byte, numRows, numRows)
	i := 0
	for i < len(s) {
		currentRow := 0
		for currentRow < numRows && i < len(s) {
			temp[currentRow] = append(temp[currentRow], s[i])
			currentRow++
			i++
		}
		currentRow -= 2
		for currentRow > 0 && i < len(s) {
			temp[currentRow] = append(temp[currentRow], s[i])
			currentRow--
			i++
		}
	}
	var res bytes.Buffer
	for i := 0; i < numRows; i++ {
		res.WriteString(string(temp[i]))
	}
	return res.String()
}

func convertHistory(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	if numRows == 2 {
		odd := ""
		event := ""
		for i, v := range s {
			if i%2 == 0 {
				odd += string(v)
			} else {
				event += string(v)
			}
		}
		return odd + event
	}
	var result []byte
	rows, rIndex, upAndDown := make([][]byte, numRows), 0, false
	for i := range s {
		rows[rIndex] = append(rows[rIndex], s[i])
		if rIndex == numRows-1 || rIndex == 0 {
			upAndDown = !upAndDown
		}
		if upAndDown {
			rIndex++
		} else {
			rIndex--
		}
	}
	for i := range rows {
		result = append(result, rows[i]...)
	}
	return string(result)
}
