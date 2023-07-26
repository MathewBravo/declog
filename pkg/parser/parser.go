package parser

import (
	"strings"

	"github.com/MathewBravo/declog/pkg/commands"
)

func ParseCommands(args []string) {
	switch strings.ToLower(args[0]) {
	case "init":
		commands.CommandExecutor(commands.Init)
	case "new-log", "newlog", "nl":
		commands.CommandExecutor(commands.NewLog)
	case "nd", "newdec", "new-dec":
		commands.CommandExecutor(commands.NewDecision)
	default:
		commands.CommandExecutor(commands.Unrecognized)
	}
}
