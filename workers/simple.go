package workers

import (
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
        c <- num;
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