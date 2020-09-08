package main

import (
	"fmt"
	"github.com/sillyhatxu/code-complete-go/design-patterns/creational-patterns/Singleton/singleton"
)

func main() {
	t := singleton.New()
	t.Test = "Singleton"
	fmt.Println(t.Test)
}
