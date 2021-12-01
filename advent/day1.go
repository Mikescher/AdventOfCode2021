package advent

import (
	"AdventOfCode2021/util"
	"math"
	"strconv"
)

func Day1Part1(ctx *util.AOCContext) (string, error) {

	incCount := 0

	val := int64(math.MaxInt64)
	for _, v := range ctx.InputLines() {
		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return "", err
		}

		if val < num {
			incCount++
		}
		val = num
	}

	return strconv.Itoa(incCount), nil
}