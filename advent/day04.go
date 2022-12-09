package advent

import (
	"AdventOfCode2021/util"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BingoBoard struct {
	Numbers [][]int
}

func (b BingoBoard) Width() int {
	if len(b.Numbers) == 0 {
		return 0
	}
	return len(b.Numbers[0])
}

func (b BingoBoard) Height() int {
	return len(b.Numbers)
}

func (b BingoBoard) Solved(numbers map[int]int) bool {
	for x := 0; x < b.Width(); x++ {
		all := true
		for y := 0; y < b.Height(); y++ {
			if _, ok := numbers[b.Numbers[y][x]]; !ok {
				all = false
				break
			}
		}
		if all {
			return true
		}
	}

	for y := 0; y < b.Height(); y++ {
		all := true
		for x := 0; x < b.Width(); x++ {
			if _, ok := numbers[b.Numbers[y][x]]; !ok {
				all = false
				break
			}
		}
		if all {
			return true
		}
	}

	return false
}

func (b BingoBoard) WinScore(numbers map[int]int, n int) int {
	usum := 0

	for y := 0; y < b.Height(); y++ {
		for x := 0; x < b.Width(); x++ {
			if _, ok := numbers[b.Numbers[y][x]]; !ok {
				usum += b.Numbers[y][x]
			}
		}
	}

	return usum * n
}

func (b BingoBoard) String(numbers map[int]int) string {
	sb := ""

	for y := 0; y < b.Height(); y++ {
		for x := 0; x < b.Width(); x++ {
			if x > 0 {
				sb += " "
			}
			if _, ok := numbers[b.Numbers[y][x]]; !ok {
				a := fmt.Sprintf("%d", b.Numbers[y][x])
				if len(a) == 1 {
					a = " " + a
				}
				sb += a
			} else {
				sb += "##"
			}
		}
		sb += "\n"
	}

	return sb
}

func Day04Part1(ctx *util.AOCContext) (string, error) {
	numbers, boards, err := loadBingoBoards(ctx)
	if err != nil {
		return "", err
	}

	ctx.Printf("\n")
	ctx.Printf("====== INITIAL ======\n")
	ctx.Printf("\n")
	for _, b := range boards {
		ctx.Printf("%s\n", b.String(make(map[int]int, 0)))
	}

	numMap := make(map[int]int, len(numbers))

	for i, n := range numbers {
		numMap[n] = n

		ctx.Printf("\n")
		ctx.Printf("====== %v ======\n", numbers[0:(i+1)])
		ctx.Printf("\n")
		for _, b := range boards {
			ctx.Printf("%s\n", b.String(numMap))
		}

		for _, board := range boards {
			if board.Solved(numMap) {
				return strconv.Itoa(board.WinScore(numMap, n)), nil
			}
		}
	}

	return "", errors.New("no board won")
}

func Day04Part2(ctx *util.AOCContext) (string, error) {
	numbers, boards, err := loadBingoBoards(ctx)
	if err != nil {
		return "", err
	}

	ctx.Printf("\n")
	ctx.Printf("====== INITIAL ======\n")
	ctx.Printf("\n")
	for _, b := range boards {
		ctx.Printf("%s\n", b.String(make(map[int]int, 0)))
	}

	numMap := make(map[int]int, len(numbers))

	for i, n := range numbers {
		numMap[n] = n

		ctx.Printf("\n")
		ctx.Printf("====== %v ======\n", numbers[0:(i+1)])
		ctx.Printf("\n")
		for _, b := range boards {
			ctx.Printf("%s\n", b.String(numMap))
		}

		remainingBoards := make([]BingoBoard, 0, len(boards))

		for _, board := range boards {
			if board.Solved(numMap) {
				if len(boards) == 1 {
					return strconv.Itoa(board.WinScore(numMap, n)), nil
				}
			} else {
				remainingBoards = append(remainingBoards, board)
			}
		}
		boards = remainingBoards
	}

	return "", errors.New("no board won")
}

func loadBingoBoards(ctx *util.AOCContext) ([]int, []BingoBoard, error) {
	var err error

	input := ctx.InputLines()

	numberStr := strings.Split(input[0], ",")

	numbers := make([]int, 0, len(numberStr))
	for _, v := range numberStr {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, nil, err
		}
		numbers = append(numbers, n)
	}

	ctx.Printf("Numbers: %v\n", numbers)

	boards := make([]BingoBoard, 0)

	rex := regexp.MustCompile("\\s+")

	for i := 2; i+4 < len(input); i += 6 {
		rowarr := make([][]int, 0, 5)
		for j := 0; j < 5; j++ {
			colarr := make([]int, 5)
			for k, v := range rex.Split(strings.TrimSpace(input[i+j]), -1) {
				colarr[k], err = strconv.Atoi(v)
				if err != nil {
					return nil, nil, err
				}
			}
			rowarr = append(rowarr, colarr)
		}
		boards = append(boards, BingoBoard{Numbers: rowarr})
	}

	return numbers, boards, nil
}
