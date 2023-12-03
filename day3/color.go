package main

import (
	"github.com/fatih/color"
	"strings"
)

func PrintLines(text string, formatter func(format string, a ...interface{})) {
	color.NoColor = false // force colour to be on
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		formatter("%s", line)
	}
}

func PrintRed(text string) {
	PrintLines(text, color.Red)
}

func PrintMagenta(text string) {
	PrintLines(text, color.Magenta)
}

func PrintBlue(text string) {
	PrintLines(text, color.Blue)
}

func PrintBlack(text string) {
	PrintLines(text, color.Black)
}
