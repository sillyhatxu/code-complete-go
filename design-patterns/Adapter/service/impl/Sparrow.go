package impl

import "fmt"

type Sparrow struct {
}

func (s *Sparrow) fly() {
	fmt.Println("Sparrow : bird fly")
}
func (s *Sparrow) makeSound() {
	fmt.Println("Sparrow : bird sound")
}
