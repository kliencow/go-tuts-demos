package main

import (
    "tuts/workers"
    "sync"
)

func main() {
    // Create a waitgroup to wait for stuff
    wg := sync.WaitGroup{}

    // Since we could have several CPUS, why not have a channel with a buffer equal to the number of cpus
    c := make(chan int)    


    // launch of bunch of go-routines to fill the channels
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(num int) {
            defer wg.Done()
            workers.FindStuff(num, c)
        }(i)
    }

    defer wg.Done()
    workers.DoStuff(c)

    // wait until it's all finished
    wg.Wait()
}
