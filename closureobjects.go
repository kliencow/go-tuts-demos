package main

import (
    "fmt"
)


func Fibo() func() int {
    oneAgo := 1
    twoAgo := 0

    return func() int {
       next := oneAgo + twoAgo
       twoAgo = oneAgo
       oneAgo = next
       
       return next
    }
}


func main() {
    fmt.Println("Starting")
    fibo := Fibo();

    for i := 0; i < 10; i++ {
        fmt.Println(fibo())
    }

}
