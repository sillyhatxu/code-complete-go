package impl

import (
	"fmt"
	"github.com/sillyhatxu/code-complete-go/design-patterns/structural-patterns/Adapter/service"
)

type BirdAdapter struct {
	bird service.Bird
}

func (ba *BirdAdapter) squeak() {
	fmt.Println("BirdAdapter squeak")
}
