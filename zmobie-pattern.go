package main

import (
	"fmt"
	"math/rand"
	"time"
)

var rando *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type ZmobieList []Zmobie

type Zmobie struct {
	MyInt   int
	NumArms int
	NumLegs int
	Name    string

	DerFuncs map[string]func(z *Zmobie)
	Shout    func(z *Zmobie)
}

func (z *Zmobie) Passer(funcName string) {
	z.DerFuncs[funcName](z)
}

func ShoutName(z *Zmobie) {
	fmt.Println(z.Name)
}

func ShoutArms(z *Zmobie) {
	fmt.Println("I've got", z.NumArms, "arms!")
}

func ShoutLegs(z *Zmobie) {
	fmt.Println("I've got", z.NumLegs, "legs!")
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
		Shout: ZmobieYell(),
	}
}

func ZmobieYell() func(z *Zmobie) {
	chance := rando.Intn(30)
	switch {
	case chance < 10:
		return ShoutName
	case chance < 20:
		return ShoutLegs
	}
	return ShoutArms
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
		z.Shout(&z)
	}
}

func main() {
	zl := make(ZmobieList, 0, 10)
	zl.SpamZmobies(20)
	zl.RollCall()
}
