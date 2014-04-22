package hub

import (
	"fmt"
)

type Channels struct {
	TestChan  chan string
	TestChan2 chan string
	TestChan3 chan string
}

type Aberations struct {
	Error chan string
	Warn  chan string
}

var cs = Channels{
	TestChan:  make(chan string, 0),
	TestChan2: make(chan string, 0),
	TestChan3: make(chan string, 0),
}

func Add(s string) {
	cs.TestChan <- s
}

func Move() {
	s := <-cs.TestChan
	cs.TestChan2 <- s
}

func MoveAgain() {
	s := <-cs.TestChan2
	cs.TestChan3 <- s
}

func Get() {
	fmt.Println(<-cs.TestChan3)
}

func Test() {
	fmt.Println("Test")
}
