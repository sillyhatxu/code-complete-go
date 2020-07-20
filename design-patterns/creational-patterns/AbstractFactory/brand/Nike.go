package brand

import (
	"github.com/sillyhatxu/code-complete-go/design-patterns/creational-patterns/AbstractFactory/brand/basic"
	"github.com/sillyhatxu/code-complete-go/design-patterns/creational-patterns/AbstractFactory/factory"
)

type Nike struct {
}

func (a *Nike) MakeShoe() factory.ShoeInterface {
	return &NikeShoe{Shoe: basic.Shoe{Logo: "Nike", Size: 20}}
}

func (a *Nike) MakeShort() factory.ShortInterface {
	return &NikeShort{Short: basic.Short{Logo: "Nike", Size: 40}}
}

type NikeShoe struct {
	basic.Shoe
}

type NikeShort struct {
	basic.Short
}
