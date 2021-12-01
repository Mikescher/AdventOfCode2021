package input

import (
	"embed"
	"fmt"
)

//go:embed *
var assets embed.FS

func GetInput(day int) (string, error) {
	bin, err := assets.ReadFile(fmt.Sprintf("day%02d.txt", day))
	if err != nil {
		return "", err
	}

	return string(bin), nil
}