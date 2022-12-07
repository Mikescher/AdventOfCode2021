package main

import (
	"AdventOfCode2021/advent"
	"AdventOfCode2021/input"
	"AdventOfCode2021/util"
)

type AOCFunc func(ctx *util.AOCContext) (string, error)

type AdventOfCode struct {
	Code map[int]AOCFunc
}

func NewAOC() *AdventOfCode {
	aoc := &AdventOfCode{
		Code: make(map[int]AOCFunc, 24),
	}

	aoc.Add(1, advent.Day01Part1, advent.Day01Part2)
	aoc.Add(2, advent.Day02Part1, advent.Day02Part2)
	aoc.Add(3, advent.Day03Part1, advent.Day03Part2)
	aoc.Add(4, advent.Day04Part1, advent.Day04Part2)
	aoc.Add(5, advent.Day05Part1, advent.Day05Part2)
	aoc.Add(6, advent.Day06Part1, advent.Day06Part2)
	aoc.Add(7, advent.Day07Part1, advent.Day07Part2)
	aoc.Add(8, advent.Day08Part1, advent.Day08Part2)
	aoc.Add(9, advent.Day09Part1, advent.Day09Part2)
	aoc.Add(10, advent.Day10Part1, advent.Day10Part2)
	aoc.Add(11, advent.Day11Part1, advent.Day11Part2)
	aoc.Add(12, advent.Day12Part1, advent.Day12Part2)
	aoc.Add(13, advent.Day13Part1, advent.Day13Part2)
	aoc.Add(14, advent.Day14Part1, advent.Day14Part2)
	aoc.Add(15, advent.Day15Part1, advent.Day15Part2)
	aoc.Add(16, advent.Day16Part1, advent.Day16Part2)
	aoc.Add(17, advent.Day17Part1, advent.Day17Part2)
	aoc.Add(18, advent.Day18Part1, advent.Day18Part2)
	aoc.Add(19, advent.Day19Part1, advent.Day19Part2)
	aoc.Add(20, advent.Day20Part1, advent.Day20Part2)
	aoc.Add(21, advent.Day21Part1, advent.Day21Part2)
	aoc.Add(12, advent.Day22Part1, advent.Day22Part2)
	aoc.Add(23, advent.Day23Part1, advent.Day23Part2)
	aoc.Add(24, advent.Day24Part1, advent.Day24Part2)
	aoc.Add(25, advent.Day25Part1, advent.Day25Part2)

	return aoc
}

func (aoc *AdventOfCode) Add(day int, p1, p2 AOCFunc) {
	aoc.Code[day*1000+1] = p1
	aoc.Code[day*1000+2] = p2
}

func (aoc *AdventOfCode) Get(day, part int) (AOCFunc, bool) {
	if v, ok := aoc.Code[day*1000+part]; ok {
		return v, true
	} else {
		return v, false
	}
}

func (aoc *AdventOfCode) Run(ctx *util.AOCContext, day, part int) (string, bool, error) {

	in, err := input.GetInput(day)
	if err != nil {
		return "", false, err
	}

	ctx.Day = day
	ctx.Part = part
	ctx.Input = in

	fn, found := aoc.Get(day, part)
	if !found {
		return "", false, nil
	}

	res, err := fn(ctx)
	if err != nil {
		return "", true, err
	}

	return res, true, nil
}
