package advent

import (
	"AdventOfCode2021/util"
	"errors"
	"fmt"
	"strconv"
)

func Day03Part1(ctx *util.AOCContext) (string, error) {

	gamma := ""
	epsilon := ""
	for i := 0; i < 12; i++ {
		count0 := 0
		count1 := 0
		for _, line := range ctx.InputLines() {
			if line[i] == '0' {
				count0++
			} else {
				count1++
			}
		}
		if count1 > count0 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	intGamma, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return "", err
	}
	intEpsilon, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", intGamma * intEpsilon), nil
}

func Day03Part2(ctx *util.AOCContext) (string, error) {
	return "", errors.New("unimplemented")
}