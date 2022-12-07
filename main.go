package main

import (
	"AdventOfCode2021/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//
// Usage:
//   ./aoc2021              [--verbose] [--silent]
//   ./aoc2021 <day>        [--verbose] [--silent]
//   ./aoc2021 <day> <part> [--verbose] [--silent]
//

func main() {
	aoc := NewAOC()

	ctx := util.NewContext()

	verbose := false

	args := os.Args[1:]
	realArgs := make([]string, 0, len(args))
	for _, v := range args {
		if strings.HasPrefix(v, "--") {
			if v == "--verbose" {
				verbose = true
			} else if v == "--silent" {
				verbose = false
			} else {
				panic("Unknown cmd argument: '" + v + "'")
			}
		} else {
			realArgs = append(realArgs, v)
		}
	}

	ctx.Silent = !verbose

	if len(realArgs) == 0 {

		for day := 1; day <= 25; day++ {

			part := 1
			result1, found, err := aoc.Run(ctx, int(day), 1)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}
			if !found {
				fmt.Printf("AOC Problem <%d|%d> not found\n", day, part)
				return
			}
			fmt.Printf("Day %02d (Part %d): %s\n", day, 1, result1)

			if day != 25 {
				part = 2
				result2, found, err := aoc.Run(ctx, int(day), 2)
				if err != nil {
					fmt.Printf("%v\n", err)
					return
				}
				if !found {
					fmt.Printf("AOC Problem <%d|%d> not found\n", day, part)
					return
				}
				fmt.Printf("Day %02d (Part %d): %s\n", day, 2, result2)
			}
		}

	} else if len(realArgs) == 1 {

		day, err := strconv.ParseInt(realArgs[0], 10, 32)
		if err != nil {
			fmt.Printf("Failed to parse <day> param %v\n", err)
			return
		}

		part := 1
		result1, found, err := aoc.Run(ctx, int(day), part)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		if !found {
			fmt.Printf("AOC Problem <%d|%d> not found\n", day, part)
			return
		}
		fmt.Printf("Day %02d (Part %d): %s\n", day, 1, result1)

		part = 2
		result2, found, err := aoc.Run(ctx, int(day), part)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		if !found {
			fmt.Printf("AOC Problem <%d|%d> not found\n", day, part)
			return
		}
		fmt.Printf("Day %02d (Part %d): %s\n", day, 2, result2)

	} else if len(realArgs) == 2 {

		day, err := strconv.ParseInt(realArgs[0], 10, 32)
		if err != nil {
			fmt.Printf("Failed to parse <day> param %v\n", err)
			return
		}

		part, err := strconv.ParseInt(realArgs[1], 10, 32)
		if err != nil {
			fmt.Printf("Failed to parse <day> param %v\n", err)
			return
		}

		result, found, err := aoc.Run(ctx, int(day), int(part))
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		if !found {
			fmt.Printf("AOC Problem <%d|%d> not found\n", day, part)
			return
		}
		fmt.Printf("Day %02d (Part %d): %s\n", day, part, result)
	} else {
		panic(fmt.Sprintf("Too many arguments (%d)", len(realArgs)))
	}

}
