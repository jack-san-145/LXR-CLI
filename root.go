package lxrcli

import (
	"fmt"
	"lxr-cli/commands"
	"os"
)

func Execute() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: lxr <command>")
	}

	switch os.Args[1] {
	case "run":
		commands.Run(os.Args[2])
	case "ping":
		commands.Ping()
	case "create":
		commands.Create(os.Args[2:])
	default:
		fmt.Println("Unknown command")
	}
}
