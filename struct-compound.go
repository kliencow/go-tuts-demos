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

func (stooge *Stooge) AddScheme(name string, scheme Scheme) {
    if (stooge.Schemes == nil) {
        stooge.Schemes = make(map[string]Scheme)
    }

    stooge.Schemes[name] = scheme
}

func (stooge Stooge) ListSchemes() {
    fmt.Println(stooge.Name)
    for name, scheme := range stooge.Schemes {
        fmt.Print("   ", name, ":")
        fmt.Println(scheme.Hilarity)
    }
}

func (stooge *Stooge) RealListSchemes() {
    fmt.Println(stooge.Name )
    for name, scheme := range stooge.Schemes {
        fmt.Print("   ", name, ":")
        fmt.Println(scheme.Hilarity)
    }
}


func MakeStooges() (moe, curly, larry Stooge) {
    moe = Stooge {
        "Moe",
        map[string] Scheme {
            "Bankers": Scheme {
                Hilarity: 55,
                ChanceOfSuccess: 0,                
            },
        },
    }

    curly = Stooge {
        Name: "Curly",  
        Schemes: make(map[string]Scheme),  
    }

    larry = Stooge{}

    PaintingScheme := Scheme{
        Hilarity: 30,
        ChanceOfSuccess: 0.0,
        GagType: "Schtick",
    }
    DanceInstructorScheme := Scheme {
        Hilarity: 75,
        ChanceOfSuccess: 0.0,
        GagType: "EyeWoolPullingOf",
    }

    larry.Name = "Larry"
    larry.AddScheme("Painters", PaintingScheme)
    larry.AddScheme("Dance Instructors", DanceInstructorScheme)

    return
}


func main() {
    moe, curly, larry := MakeStooges()
    fmt.Println(moe.Name)
    fmt.Println(curly.Name)
    fmt.Println(larry.Name)

    fmt.Println()
    fmt.Println("List Schemes Simple")    

    moe.ListSchemes()
    curly.ListSchemes()
    larry.ListSchemes()

    fmt.Println()
    fmt.Println("List Schemes by Refenence")

    moe.RealListSchemes()
    curly.RealListSchemes()
    larry.RealListSchemes()
}

