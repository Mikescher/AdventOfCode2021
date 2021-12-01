package main

import (
	"AdventOfCode2021/util"
	"fmt"
)

func main()  {
	aoc := NewAOC()

	result, err := aoc.Run(util.NewContext(), 0, 0)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Println(result)
}