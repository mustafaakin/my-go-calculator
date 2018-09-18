package main

import (
	"github.com/fatih/color"
	"os"
	"strconv"
)

func calculate(a, b int64, op string) int64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	default:
		panic("I'm sorry Dave, I'm afraid I cannot do that")
	}
}

func main() {
	if len(os.Args) != 4 {
		// os.Args[0] is the application name
		panic("Usage: <NUMBER> <OPERATION> <NUMBER>")
	}

	a, _ := strconv.ParseInt(os.Args[1], 10, 64)
	b, _ := strconv.ParseInt(os.Args[3], 10, 64)
	op := os.Args[2]
	result := calculate(a, b, op)
	color.Yellow("Result is = %d", result)
}
