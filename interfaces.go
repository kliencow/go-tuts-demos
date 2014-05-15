package main

import (
	"fmt"
)

type Fooer interface {
	SetString(s string)
	GetString() string
}

type Foo struct {
	s string
}

func (f *Foo) GetString() string {
	return f.s
}

func (f *Foo) SetString(s string) {
	f.s = s
}

func PrintTwice(f Fooer) {
	fmt.Println(f.GetString())
	fmt.Println(f.GetString())
}

func main() {
	fmt.Println("Starting")

	f := new(Foo)

	f.SetString("YO")
	PrintTwice(f)
}
