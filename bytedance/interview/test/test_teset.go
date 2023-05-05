package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	assert.EqualValues(t, "6", Calculator("1+2+3"))
}
