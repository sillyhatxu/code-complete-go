package build

import (
	"fmt"
	"github.com/sillyhatxu/code-complete-go/design-patterns/AbstractFactory/brand"
	"github.com/sillyhatxu/code-complete-go/design-patterns/AbstractFactory/factory"
)

const (
	Adidas = "Adidas"
	Fila   = "FILA"
	HM     = "HM"
	Nike   = "NIKE"
	UNIQLO = "UNIQLO"
)

func Factory(name string) (factory.BrandFactory, error) {
	if name == Adidas {
		return &brand.Adidas{}, nil
	} else if name == Fila {
		return &brand.Fila{}, nil
	} else if name == HM {
		return &brand.HM{}, nil
	} else if name == Nike {
		return &brand.Nike{}, nil
	} else if name == UNIQLO {
		return &brand.UNIQLO{}, nil
	}
	return nil, fmt.Errorf("wrong name type passed")
}
