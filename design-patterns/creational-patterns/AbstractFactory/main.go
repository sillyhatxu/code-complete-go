package main

import (
	"fmt"
	"github.com/sillyhatxu/code-complete-go/design-patterns/creational-patterns/AbstractFactory/build"
	"github.com/sillyhatxu/code-complete-go/design-patterns/creational-patterns/AbstractFactory/factory"
)

func main() {
	adidasFactory, err := build.Factory(build.Adidas)
	if err != nil {
		panic(err)
	}
	filaFactory, err := build.Factory(build.Fila)
	if err != nil {
		panic(err)
	}
	hmFactory, err := build.Factory(build.HM)
	if err != nil {
		panic(err)
	}
	nikeFactory, err := build.Factory(build.Nike)
	if err != nil {
		panic(err)
	}
	uniqloFactory, err := build.Factory(build.UNIQLO)
	if err != nil {
		panic(err)
	}
	PrintFactory(adidasFactory)
	PrintFactory(filaFactory)
	PrintFactory(hmFactory)
	PrintFactory(nikeFactory)
	PrintFactory(uniqloFactory)
}

func PrintFactory(f factory.BrandFactory) {
	fmt.Printf("Shoe -> Logo: %s; Size: %d\n", f.MakeShoe().GetLogo(), f.MakeShoe().GetSize())
	fmt.Printf("Short -> Logo: %s; Size: %d\n", f.MakeShort().GetLogo(), f.MakeShort().GetSize())
}
