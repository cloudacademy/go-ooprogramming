package main

import (
	"fmt"

	"github.com/cloudacademy/go-ooprogramming/pkg/crazystrings"
)

func main() {
	var strs []crazystrings.CrazyString

	cs1, _ := crazystrings.NewCrazyString("encapsulation")
	cs2, _ := crazystrings.NewBigCrazyString("encapsulation")

	cs3, _ := crazystrings.NewCrazyString("abstraction")
	cs4, _ := crazystrings.NewBigCrazyString("abstraction")

	cs5, _ := crazystrings.NewCrazyString("polymorphism")
	cs6, _ := crazystrings.NewBigCrazyString("polymorphism")

	cs7, _ := crazystrings.NewCrazyString("composition")
	cs8, _ := crazystrings.NewBigCrazyString("composition")

	cs9, _ := crazystrings.NewFatCrazyString("ooprogramming")
	cs10, _ := crazystrings.NewVeryFatCrazyString("ooprogramming", 3)

	strs = append(strs, cs1)
	strs = append(strs, cs2)
	strs = append(strs, cs3)
	strs = append(strs, cs4)
	strs = append(strs, cs5)
	strs = append(strs, cs6)
	strs = append(strs, cs7)
	strs = append(strs, cs8)
	strs = append(strs, cs9)
	strs = append(strs, cs10)

	for _, str := range strs {
		fmt.Printf("%T : %s\n", str, str.Scramble())
	}
}
