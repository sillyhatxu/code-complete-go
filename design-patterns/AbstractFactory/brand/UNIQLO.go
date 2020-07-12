package brand

import (
	"github.com/sillyhatxu/code-complete-go/design-patterns/AbstractFactory/brand/basic"
	"github.com/sillyhatxu/code-complete-go/design-patterns/AbstractFactory/factory"
)

type UNIQLO struct {
}

func (a *UNIQLO) MakeShoe() factory.ShoeInterface {
	return &UNIQLOShoe{Shoe: basic.Shoe{Logo: "UNIQLO", Size: 22}}
}

func (a *UNIQLO) MakeShort() factory.ShortInterface {
	return &UNIQLOShort{Short: basic.Short{Logo: "UNIQLO", Size: 42}}
}

type UNIQLOShoe struct {
	basic.Shoe
}

type UNIQLOShort struct {
	basic.Short
}
