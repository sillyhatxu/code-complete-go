package impl

import (
	"fmt"
	"github.com/sillyhatxu/code-complete-go/design-patterns/Adapter/service"
)

type BirdAdapter struct {
	bird service.Bird
}

func (ba *BirdAdapter) squeak() {
	bird.
		fmt.Println("BirdAdapter squeak")
}
