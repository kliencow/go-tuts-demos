package main

import (
	"fmt"
	"math/rand"
	"time"
)

var rando *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type ZmobieList []Zmobie

type Zmobie struct {
	MyInt    int
	Name     string
	DerFuncs map[string]func(z *Zmobie)
}

func (z *Zmobie) Passer(funcName string) {
	z.DerFuncs[funcName](z)
}

func (z *Zmobie) Shout() {
	fmt.Println(z.Name)
}

func MoreZmobies() Zmobie {
	return Zmobie{
		MyInt:   4,
		NumArms: rando.Intn(3),
		NumLegs: rando.Intn(3),
		Name:    ZmobieName(),
		DerFuncs: map[string]func(z *Zmobie){
			"PrintMyInt": func(z *Zmobie) {
				fmt.Println(z.MyInt)
			},
			"PrintTwiceMyInt": func(z *Zmobie) {
				fmt.Println(z.MyInt * 2)
			},
			"DoubleMyInt": func(z *Zmobie) {
				z.MyInt = z.MyInt * 2
			},
		},
	}
}

func ZmobieName() string {
	nameLength := rando.Intn(17) + 3
	name := ""

	for i := 0; i < nameLength; i++ {
		char := rando.Intn(5)
		switch {
		case char == 0:
			name += "G"
		case char == 1:
			name += "U"
		case char == 2:
			name += "R"
		case char == 3:
			name += "L"
		case char == 4:
			name += "B"
		}
	}

	return name
}

func (zl *ZmobieList) SpamZmobies(spamCount int) {
	for i := 0; i < spamCount; i++ {
		*zl = append(*zl, MoreZmobies())
	}
}

func (zl ZmobieList) RollCall() {
	for _, z := range zl {
		z.Shout()
	}
}

func main() {
	zl := make(ZmobieList, 0, 10)
	zl.SpamZmobies(20)
	zl.RollCall()
}
