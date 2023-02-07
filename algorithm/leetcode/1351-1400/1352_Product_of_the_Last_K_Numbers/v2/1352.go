package v2

type ProductOfNumbers struct {
	number []int
	last   int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{number: []int{}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.number = []int{}
		this.last = 0

		return
	}

	this.last = num
	if len(this.number) > 0 {
		this.last *= this.number[len(this.number)-1]
	}
	this.number = append(this.number, this.last)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if l := len(this.number); k > l {
		return 0
	} else if k == l {
		return this.last
	}
	return this.last / this.number[len(this.number)-k-1]
}
