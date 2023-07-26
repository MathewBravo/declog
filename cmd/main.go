package main

import (
	"os"

	"github.com/MathewBravo/declog/pkg/parser"
)

func main() {
	args := os.Args
	parser.ParseCommands(args[1:])
}
