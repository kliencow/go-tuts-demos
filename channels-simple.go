package main

import (
    "runtime"
    "fmt"
    "time"
    "math/rand"    
)

/*
    NextPause is a closure around getting new durations from a consistent seed. 
    It was found that we could run out of random numbers, granter we were using numbers on the order
    of 10M * 200M. Still, it should no be running out of numbers. To combant this issue, we create 
    a PNRG for each go-routine.
*/
func NextPause() func() (pause time.Duration) {
    rando := rand.New(rand.NewSource(time.Now().UnixNano()))
    return func() (pause time.Duration) {
        pause = time.Duration((rando.Intn(500)) * int(time.Millisecond))
        return
    }
}


/*
    FindStuff is a worker meant to be used in a go-routine. His job is to fill the channel with cruct. In this
    case his cruft is his channel number.
*/
func FindStuff(num int, c chan int) {    
    pause := NextPause();

    for i := 0; i < 10; i++ {                
        time.Sleep(pause())
        c <- num
    }        
}


/*
    DoStuff is a worker meant to be used in a go-routine. His job is to take messages off of the channel and print
    the number he finds. 
*/
func DoStuff(c chan int) {
    for v := range c {
         fmt.Println(v)
    }
}


func main() {
    // For fun, we choose to use the maximum number of CPUs availble to do work.
    cpus := runtime.NumCPU()
    runtime.GOMAXPROCS(cpus)    

    // Since we could have several CPUS, why not have a channel with a buffer equal to the number of cpus
    c := make(chan int, cpus)    

    // launch of bunch of go-routines to fill the channels
    for i := 0; i < 3; i++ {
        go FindStuff(i, c)
    }

    // lauch a go-routine to take stuff off of the channel
    go DoStuff(c)

    // since we're being simple, let's just wait some arbitrary amount of time for the work to be done.
    // Better solutions would be:
    //      a non-go-routine DoStuff() with a switch. 
    //      using wait-groups
    //      create some complex system of non-go-routine channel counters to make sure we're finished before releasing
    //      the loops.
    time.Sleep(10000 * time.Millisecond)
}