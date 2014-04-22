package main

import (
	"fmt"
	"time"
	"tuts/hub"
)

func main() {
	fmt.Println("Starting")
	hub.Test()

	go hub.Add("passing through")
	go hub.Move()
	go hub.MoveAgain()
	go hub.Get()

	time.Sleep(1000 * time.Millisecond)
}
