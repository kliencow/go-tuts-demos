package main

import (
    "fmt"
)

type Stooge struct {
    Name string
    Schemes map[string]Scheme
}

type Scheme struct {
    Hilarity int
    ChanceOfSuccess float64
    GagType string
}


func MakeStooges() (moe, curly, larry Stooge) {
    moe = Stooge {
        "Moe",
        nil,
    }

    curly = Stooge {
        Name: "Curly",
    }

    larry = Stooge{}
    larry.Name = "Larry"

    return
}



func main() {
    fmt.Println("Starting")

    moe, larry, curly := MakeStooges()

    fmt.Println(moe.Name)
    fmt.Println(curly.Name)
    fmt.Println(larry.Name)
}

