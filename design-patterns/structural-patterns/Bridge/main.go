package main

import (
	"fmt"
	"github.com/sillyhatxu/code-complete-go/design-patterns/structural-patterns/Bridge/model"
)

func main() {
	openGL := &model.OpenGL{}
	direct2D := &model.Direct2D{}

	circle := &model.Circle{
		Center: model.Point{X: 100, Y: 100},
		Radius: 50,
	}
	fmt.Println()
	circle.DrawingContext = openGL
	circle.Draw()
	fmt.Println()
	circle.DrawingContext = direct2D
	circle.Draw()
}
