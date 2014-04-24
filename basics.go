package main

import (
	"fmt"
)

func arrayPlay() {
	fmt.Println("ARRS!")
	arr := [3]int{4, 5, 6}

	for i, v := range arr {
		fmt.Println(i, ": ", v)
	}

}

func niceSlice() {
	fmt.Println("SLICES!")
	slice := []int{5, 6, 7, 8, 9}

	fmt.Println("Slice Length: ", len(slice))
	fmt.Println("Slice Capacity: ", cap(slice))

	fmt.Println(slice[3:5])
}

func meanSlice() {
	fmt.Println("MORE SLICES!")

	slice := make([]int, 2, 4)
	fmt.Println("Slice Length: ", len(slice))
	fmt.Println("Slice Capacity: ", cap(slice))

	fmt.Println(slice)

	slice[1] = 2
	fmt.Println(slice)

	slice = slice[:4]
	fmt.Println(slice)

}

func sliceMadeInteresting() {
	fmt.Println("MORE SLICES GO!")

	s := []int{1, 2, 3, 4, 5, 6, 7}
	s2 := s[2:4]
	s3 := dontCare(s)

	fmt.Println(&s[2])
	fmt.Println(&s2[0])
	fmt.Println(&s3[0])
}

func dontCare(s []int) (interesting []int) {
	bit := s[2:4]

	interesting = make([]int, len(bit))

	copy(interesting, bit)
	return
}

func mapsMadeFun() {
	fmt.Println("MAPS!")
	m := make(map[string]int)
	m["Foo"] = 3
	m["Bar"] = 4
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, ": ", v)
	}
}

func main() {
	//arrayPlay()
	//niceSlice()
	//meanSlice()

	//sliceMadeInteresting()

	mapsMadeFun()

}
