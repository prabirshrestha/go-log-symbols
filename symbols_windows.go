package logSymbols

import (
	"github.com/fatih/color"
)

func Info() string {
	return color.New(color.FgBlue).SprintFunc()("i")
}

func Success() string {
	return color.New(color.FgGreen).SprintFunc()("√")
}

func Warning() string {
	return color.New(color.FgYellow).SprintFunc()("‼")
}

func Error() string {
	return color.New(color.FgRed).SprintFunc()("×")
}
