package _90_Reverse_Bits

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverse(t *testing.T) {
	input1 := "00000010100101000001111010011100"
	var temp1 []string
	for i := 0; i < len(input1); i++ {
		temp1 = append(temp1, string(input1[i]))
	}
	reverse(&temp1, 0, len(temp1)-1)
	for i := 0; i < len(temp1); i++ {
		fmt.Print(temp1[i])
	}
	fmt.Println()
	fmt.Println("00111001011110000010100101000000")
}

func reverse(temp *[]string, start int, end int) {
	for start < end {
		(*temp)[start], (*temp)[end] = (*temp)[end], (*temp)[start]
		start++
		end--
	}
}

func Test_toInt(t *testing.T) {
	fmt.Println(uint32(43261596))
	//test := uint32(00111001011110000010100101000000)
	//fmt.Println(int(test))
}
func Test_reverseBits(t *testing.T) {
	assert.EqualValues(t, 964176192, reverseBits(uint32(43261596)))
	assert.EqualValues(t, 3221225471, reverseBits(uint32(4294967293)))
	//assert.EqualValues(t, 00111001011110000010100101000000, reverseBits(uint32(00000010100101000001111010011100)))
	//assert.EqualValues(t, 10111111111111111111111111111111, reverseBits(uint32(11111111111111111111111111111101)))
}

//00000010100101000001111010011100
//00111001011110000010100101000000
