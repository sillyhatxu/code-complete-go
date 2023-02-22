package less_than_n_and_largest_number

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(strings.Replace(fmt.Sprint(a), "", "", -1))
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(a), " ", "", -1), "[]"))
}

func Test_lessThanNAndLargestNumber1(t *testing.T) {
	assert.EqualValues(t, 2499, lessThanNAndLargestNumber1(2533, []int{1, 2, 4, 9}))
	assert.EqualValues(t, 9999, lessThanNAndLargestNumber1(12345, []int{6, 7, 9}))
	assert.EqualValues(t, 999, lessThanNAndLargestNumber1(2533, []int{0, 3, 9}))
	assert.EqualValues(t, 2499, lessThanNAndLargestNumber1(2533, []int{1, 2, 4, 9}))
	assert.EqualValues(t, 0, lessThanNAndLargestNumber1(1, []int{0, 0}))
	assert.EqualValues(t, 222, lessThanNAndLargestNumber1(1111, []int{2}))
	//assert.EqualValues(t, 12121, lessThanNAndLargestNumber1(12121, []int{1, 2}))
	assert.EqualValues(t, 2449, lessThanNAndLargestNumber1(2451, []int{4, 2, 5, 9}))
	assert.EqualValues(t, 222, lessThanNAndLargestNumber1(1111, []int{2}))
	assert.EqualValues(t, 22222, lessThanNAndLargestNumber1(33333, []int{2}))
	assert.EqualValues(t, 313, lessThanNAndLargestNumber1(321, []int{1, 2, 3}))
	assert.EqualValues(t, 2, lessThanNAndLargestNumber1(3, []int{1, 2, 3}))
	assert.EqualValues(t, 43999, lessThanNAndLargestNumber1(44321, []int{4, 3, 9}))
	assert.EqualValues(t, 22555, lessThanNAndLargestNumber1(24131, []int{2, 4, 5}))
	assert.EqualValues(t, 2299, lessThanNAndLargestNumber1(2409, []int{1, 2, 4, 9}))
	assert.EqualValues(t, 4919, lessThanNAndLargestNumber1(4921, []int{1, 2, 4, 9}))
	assert.EqualValues(t, 1999, lessThanNAndLargestNumber1(2100, []int{1, 2, 4, 9}))
	assert.EqualValues(t, 4955, lessThanNAndLargestNumber1(4956, []int{4, 5, 6, 9}))
	assert.EqualValues(t, 4949, lessThanNAndLargestNumber1(4956, []int{0, 4, 9}))
	assert.EqualValues(t, 99, lessThanNAndLargestNumber1(589, []int{6, 9}))
	assert.EqualValues(t, 1111, lessThanNAndLargestNumber1(4956, []int{1}))
	assert.EqualValues(t, 888, lessThanNAndLargestNumber1(1200, []int{6, 7, 8}))
	assert.EqualValues(t, 2555, lessThanNAndLargestNumber1(3211, []int{2, 3, 5}))
	assert.EqualValues(t, 1199, lessThanNAndLargestNumber1(1201, []int{1, 2, 9}))
	assert.EqualValues(t, 994, lessThanNAndLargestNumber1(999, []int{1, 2, 4, 9}))
	assert.EqualValues(t, 999, lessThanNAndLargestNumber1(1111, []int{1, 2, 4, 9}))
	assert.EqualValues(t, 2499, lessThanNAndLargestNumber1(2533, []int{1, 2, 4, 9}))
	assert.EqualValues(t, 4919, lessThanNAndLargestNumber1(4921, []int{1, 2, 4, 9}))
	assert.EqualValues(t, 1244, lessThanNAndLargestNumber1(1249, []int{1, 2, 4, 9}))
	assert.EqualValues(t, 22555, lessThanNAndLargestNumber1(24131, []int{2, 4, 5}))
}

func Test_lessThanNAndLargestNumberInterview(t *testing.T) {
	//assert.EqualValues(t, 2499, lessThanNAndLargestNumberInterview(2533, []int{1, 2, 4, 9}))
	//assert.EqualValues(t, 9999, lessThanNAndLargestNumberInterview(12345, []int{6, 7, 9}))
	//assert.EqualValues(t, 999, lessThanNAndLargestNumberInterview(2533, []int{0, 3, 9}))
	//assert.EqualValues(t, 2499, lessThanNAndLargestNumberInterview(2533, []int{1, 2, 4, 9}))
	//assert.EqualValues(t, 0, lessThanNAndLargestNumberInterview(1, []int{0, 0}))
	//assert.EqualValues(t, 222, lessThanNAndLargestNumberInterview(1111, []int{2}))
	//assert.EqualValues(t, 12121, lessThanNAndLargestNumberInterview(12121, []int{1, 2}))
	assert.EqualValues(t, 2449, lessThanNAndLargestNumberInterview(2451, []int{4, 2, 5, 9}))
	assert.EqualValues(t, 222, lessThanNAndLargestNumberInterview(1111, []int{2}))
	assert.EqualValues(t, 22222, lessThanNAndLargestNumberInterview(33333, []int{2}))
	assert.EqualValues(t, 313, lessThanNAndLargestNumberInterview(321, []int{1, 2, 3}))
	//assert.EqualValues(t, 2, lessThanNAndLargestNumberInterview(3, []int{1, 2, 3}))
	//assert.EqualValues(t, 43999, lessThanNAndLargestNumberInterview(44321, []int{4, 3, 9}))
	//assert.EqualValues(t, 22555, lessThanNAndLargestNumberInterview(24131, []int{2, 4, 5}))
	//assert.EqualValues(t, 2299, lessThanNAndLargestNumberInterview(2409, []int{1, 2, 4, 9}))
	//assert.EqualValues(t, 4919, lessThanNAndLargestNumberInterview(4921, []int{1, 2, 4, 9}))
	//assert.EqualValues(t, 1999, lessThanNAndLargestNumberInterview(2100, []int{1, 2, 4, 9}))
	//assert.EqualValues(t, 4955, lessThanNAndLargestNumberInterview(4956, []int{4, 5, 6, 9}))
	//assert.EqualValues(t, 4949, lessThanNAndLargestNumberInterview(4956, []int{0, 4, 9}))
	//assert.EqualValues(t, 99, lessThanNAndLargestNumberInterview(589, []int{6, 9}))
	//assert.EqualValues(t, 1111, lessThanNAndLargestNumberInterview(4956, []int{1}))
	//assert.EqualValues(t, 888, lessThanNAndLargestNumberInterview(1200, []int{6, 7, 8}))
	//assert.EqualValues(t, 2555, lessThanNAndLargestNumberInterview(3211, []int{2, 3, 5}))
	//assert.EqualValues(t, 1199, lessThanNAndLargestNumberInterview(1201, []int{1, 2, 9}))
	//assert.EqualValues(t, 994, lessThanNAndLargestNumberInterview(999, []int{1, 2, 4, 9}))
	//assert.EqualValues(t, 999, lessThanNAndLargestNumberInterview(1111, []int{1, 2, 4, 9}))
	//assert.EqualValues(t, 2499, lessThanNAndLargestNumberInterview(2533, []int{1, 2, 4, 9}))
	//assert.EqualValues(t, 4919, lessThanNAndLargestNumberInterview(4921, []int{1, 2, 4, 9}))
	//assert.EqualValues(t, 1244, lessThanNAndLargestNumberInterview(1249, []int{1, 2, 4, 9}))
	//assert.EqualValues(t, 22555, lessThanNAndLargestNumberInterview(24131, []int{2, 4, 5}))
}
