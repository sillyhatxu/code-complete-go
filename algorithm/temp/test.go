package main

import (
	"fmt"
	"reflect"
)

type Test2 struct {
	y int
}

func (Test2) Test2Function() {}

type Struct1 struct {
	value int
	Test2
	*Test1
}

func (*Struct1) M2() {}

type Q *Struct1

var t Struct1 = Struct1{Test2: Test2{}, Test1: &Test1{}}
var p = &Struct1{Test2: Test2{}, Test1: &Test1{}}
var q Q = p

type Test1 struct {
	x int
}

func (*Test1) Test1Function() {
	fmt.Println("1")
}

type AAAA *Test1

func main() {
	var l = AAAA(&Test1{x: 1})
	fmt.Println("type:", reflect.TypeOf((*l)))
	fmt.Println("type:", reflect.TypeOf(l))
	//var q = Q(&Struct1{Test2: Test2{}, Test1: &Test1{}})
	//
	//q.Test1.Test1Function()
	//(*q).Test1Function()
	//p.Test1Function()
}
