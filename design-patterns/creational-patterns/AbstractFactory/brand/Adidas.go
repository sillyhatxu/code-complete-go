package brand

import (
	"github.com/sillyhatxu/code-complete-go/design-patterns/creational-patterns/AbstractFactory/brand/basic"
	"github.com/sillyhatxu/code-complete-go/design-patterns/creational-patterns/AbstractFactory/factory"
)

type Adidas struct {
	Test string
}

func (a *Adidas) MakeShoe() factory.ShoeInterface {
	return &AdidasShoe{Shoe: basic.Shoe{Logo: "Adidas", Size: 14}}
}

func (a *Adidas) MakeShort() factory.ShortInterface {
	return &AdidasShort{Short: basic.Short{Logo: "Adidas", Size: 34}}
}

type AdidasShoe struct {
	basic.Shoe
}

type AdidasShort struct {
	basic.Short
}
