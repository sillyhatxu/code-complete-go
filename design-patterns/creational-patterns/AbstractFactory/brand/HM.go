package brand

import (
	"github.com/sillyhatxu/code-complete-go/design-patterns/creational-patterns/AbstractFactory/brand/basic"
	"github.com/sillyhatxu/code-complete-go/design-patterns/creational-patterns/AbstractFactory/factory"
)

type HM struct {
}

func (a *HM) MakeShoe() factory.ShoeInterface {
	return &HMShoe{Shoe: basic.Shoe{Logo: "HM", Size: 18}}
}

func (a *HM) MakeShort() factory.ShortInterface {
	return &HMShort{Short: basic.Short{Logo: "HM", Size: 38}}
}

type HMShoe struct {
	basic.Shoe
}

type HMShort struct {
	basic.Short
}
