package main

import (
	"fmt"
	"github.com/sillyhatxu/code-complete-go/design-patterns/Builder/build"
	"time"
)

func main() {
	user := build.New().Id("1").Name("haha").Age(15).Status(true).CreateTime(time.Now()).Build()
	fmt.Println(user.String())
	user = build.New().Id("1").Name("haha").Status(true).CreateTime(time.Now()).Build()
	fmt.Println(user.String())
}
