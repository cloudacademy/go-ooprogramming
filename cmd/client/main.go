package main

import (
	"fmt"

	"github.com/cloudacademy/go-ooprogramming/pkg/crazystrings"
)

func main() {
	var strs []crazystrings.CrazyString

	cs1, _ := crazystrings.NewCrazyString("composition")
	cs2, _ := crazystrings.NewBigCrazyString("composition")

	cs3, _ := crazystrings.NewCrazyString("polymorphism")
	cs4, _ := crazystrings.NewBigCrazyString("polymorphism")

	strs = append(strs, cs1)
	strs = append(strs, cs2)
	strs = append(strs, cs3)
	strs = append(strs, cs4)

	for _, str := range strs {
		fmt.Println(str.Scramble())
	}
}
