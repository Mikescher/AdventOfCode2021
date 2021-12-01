package util

import (
	"strings"
)

type AOCContext struct {
	Day   int
	Part  int
	Input string
}

func NewContext() *AOCContext {
	return &AOCContext{
		Day:   -1,
		Part:  -1,
		Input: "",
	}
}

func (c *AOCContext) InputLines() []string {
	r := make([]string, 0)
	for _, v := range strings.Split(c.Input, "\n") {
		if v != "" {
			r = append(r, v)
		}
	}
	return r
}