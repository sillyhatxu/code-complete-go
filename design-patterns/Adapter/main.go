package main

import (
	"github.com/sillyhatxu/code-complete-go/design-patterns/Adapter/model"
)

func main() {
	pb := &model.PrintBanner{}
	pb.PrintBanner("Hello")
	pb.PrintWeek()
	pb.PrintStrong()
}
