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

	aoc.Add(1, advent.Day1Part1, advent.Day1Part2)
	aoc.Add(2, nil, nil)
	aoc.Add(3, nil, nil)
	aoc.Add(4, nil, nil)
	aoc.Add(5, nil, nil)
	aoc.Add(6, nil, nil)
	aoc.Add(7, nil, nil)
	aoc.Add(8, nil, nil)
	aoc.Add(9, nil, nil)
	aoc.Add(10, nil, nil)
	aoc.Add(11, nil, nil)
	aoc.Add(12, nil, nil)
	aoc.Add(13, nil, nil)
	aoc.Add(14, nil, nil)
	aoc.Add(15, nil, nil)
	aoc.Add(16, nil, nil)
	aoc.Add(17, nil, nil)
	aoc.Add(18, nil, nil)
	aoc.Add(19, nil, nil)
	aoc.Add(20, nil, nil)
	aoc.Add(21, nil, nil)
	aoc.Add(12, nil, nil)
	aoc.Add(23, nil, nil)
	aoc.Add(24, nil, nil)

	return aoc
}

func (aoc *AdventOfCode) Add(day int, p1, p2 AOCFunc) {
	aoc.Code[day*1000+1] = p1
	aoc.Code[day*1000+2] = p2
}

func (aoc *AdventOfCode) Get(day, part int) AOCFunc {
	return aoc.Code[day*1000+part]
}

func (aoc *AdventOfCode) Run(ctx *util.AOCContext, day, part int) (string, error) {

	in, err := input.GetInput(day)
	if err != nil {
		return "", err
	}

	ctx.Day = day
	ctx.Part = part
	ctx.Input = in

	return aoc.Get(day, part)(ctx)
}