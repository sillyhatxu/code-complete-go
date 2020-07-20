package brand

import (
	"github.com/sillyhatxu/code-complete-go/design-patterns/creational-patterns/AbstractFactory/brand/basic"
	"github.com/sillyhatxu/code-complete-go/design-patterns/creational-patterns/AbstractFactory/factory"
)

type Fila struct {
	Test string
}

func (a *Fila) MakeShoe() factory.ShoeInterface {
	return &FilaShoe{Shoe: basic.Shoe{Logo: "Fila", Size: 16}}
}

func (a *Fila) MakeShort() factory.ShortInterface {
	return &FilaShort{Short: basic.Short{Logo: "Fila", Size: 36}}
}

type FilaShoe struct {
	basic.Shoe
}

type FilaShort struct {
	basic.Short
}
