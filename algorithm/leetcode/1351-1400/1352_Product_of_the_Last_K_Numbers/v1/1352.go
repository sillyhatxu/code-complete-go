package v1

type ProductOfNumbers struct {
	arr []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{arr: []int{}}
}

func (this *ProductOfNumbers) Add(num int) {
	this.arr = append(this.arr, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	res := 1
	for i := len(this.arr) - 1; i >= len(this.arr)-k; i-- {
		res = res * this.arr[i]
	}
	return res
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
