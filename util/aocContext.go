package util

import (
	"fmt"
	"strings"
)

type AOCContext struct {
	Day   int
	Part  int
	Input string

	Silent bool
}

func NewContext() *AOCContext {
	return &AOCContext{
		Day:   -1,
		Part:  -1,
		Input: "",

		Silent: false,
	}
}

func (c *AOCContext) InputLines() []string {
	r := make([]string, 0)
	for _, v := range strings.Split(c.Input, "\n") {
		r = append(r, v)
	}
	for r[len(r)-1] == "" {
		r = r[0 : len(r)-1]
	}
	return r
}

func (c *AOCContext) Printf(format string, a ...interface{}) {
	if !c.Silent {
		fmt.Printf(format, a...)
	}
}
