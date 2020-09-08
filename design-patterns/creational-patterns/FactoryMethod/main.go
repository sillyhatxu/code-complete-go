package main

import (
	"github.com/sillyhatxu/code-complete-go/design-patterns/creational-patterns/FactoryMethod/fm"
	"os"
)

func main() {
	doc := &fm.Document{
		ShapeFactories: []fm.ShapeFactory{
			&fm.CircleFactory{},
			&fm.RactangleFactory{},
		},
	}
	err := doc.Draw(os.Stdout)
	if err != nil {
		panic(err)
	}
}
