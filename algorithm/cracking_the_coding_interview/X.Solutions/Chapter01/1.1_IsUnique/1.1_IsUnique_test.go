package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println('a')
	fmt.Println('z')
	fmt.Println(int('b'))
	fmt.Println(int('c'))
	fmt.Println(int('d'))
	fmt.Println(int('e'))
	fmt.Println(int('f'))
}

//
//func TestIsUniqueChars(t *testing.T) {
//	result := IsUniqueChars("abcdefghijklmnopqrstuvwxyzg")
//	assert.Equal(t, false, result)
//	result = IsUniqueChars("abcdefghijklmnopqrstuvwxyz")
//	assert.Equal(t, true, result)
//	result = IsUniqueChars("qwertyuiopasdfghjklzxcvbnma")
//	assert.Equal(t, false, result)
//}
//
//func TestIsUniqueChars2(t *testing.T) {
//	result := IsUniqueChars2("abcdefghijklmnopqrstuvwxyzg")
//	assert.Equal(t, false, result)
//	result = IsUniqueChars2("abcdefghijklmnopqrstuvwxyz")
//	assert.Equal(t, true, result)
//	result = IsUniqueChars2("qwertyuiopasdfghjklzxcvbnma")
//	assert.Equal(t, false, result)
//}
//
//func TestIsUniqueChars3(t *testing.T) {
//	//result := IsUniqueChars3("1234567890abcdefghijklmnopqrstuvwxyz")
//	result := IsUniqueChars3("abcdefghijklmnopqrstuvwxyzg")
//	assert.Equal(t, false, result)
//	result = IsUniqueChars3("abcdefghijklmnopqrstuvwxyz")
//	assert.Equal(t, true, result)
//	result = IsUniqueChars3("qwertyuiopasdfghjklzxcvbnma")
//	assert.Equal(t, false, result)
//}

func BenchmarkString(b *testing.B) {
	fmt.Println(101 >> 3)
	assert.Equal(b, int32(115), 's')
	assert.Equal(b, int32(97), 'a')
	assert.Equal(b, int32(98), 'b')
	assert.Equal(b, int32(99), 'c')
	assert.Equal(b, int32(46), '.')
}

func BenchmarkIsUniqueChars1(b *testing.B) {
	for i := 0; i < 100000; i++ {
		result := IsUniqueChars1("abcdefghijklmnopqrstuvwxyzg")
		assert.Equal(b, false, result)
		result = IsUniqueChars1("abcdefghijklmnopqrstuvwxyz")
		assert.Equal(b, true, result)
		result = IsUniqueChars1("qwertyuiopasdfghjklzxcvbnma")
		assert.Equal(b, false, result)
	}
}

func BenchmarkIsUniqueChars2(b *testing.B) {
	for i := 0; i < 100000; i++ {
		result := IsUniqueChars2("abcdefghijklmnopqrstuvwxyzg")
		assert.Equal(b, false, result)
		result = IsUniqueChars2("abcdefghijklmnopqrstuvwxyz")
		assert.Equal(b, true, result)
		result = IsUniqueChars2("qwertyuiopasdfghjklzxcvbnma")
		assert.Equal(b, false, result)
	}
}

func BenchmarkIsUniqueChars3(b *testing.B) {
	for i := 0; i < 100000; i++ {
		//result := IsUniqueChars3("1234567890abcdefghijklmnopqrstuvwxyz")
		result := IsUniqueChars3("abcdefghijklmnopqrstuvwxyzg")
		assert.Equal(b, false, result)
		result = IsUniqueChars3("abcdefghijklmnopqrstuvwxyz")
		assert.Equal(b, true, result)
		result = IsUniqueChars3("qwertyuiopasdfghjklzxcvbnma")
		assert.Equal(b, false, result)
	}
}
