package lxrcli

import (
	"log"
	"lxr-cli/commands"
	"os"
)

func Execute() {

	if len(os.Args) < 2 {
		log.Fatal("Usage: lxr <command>")
	}

	switch os.Args[1] {
	case "run":
		//do nothing
	case "ping":
		commands.Ping()
	default:
		log.Fatal("Unknown command")
	}
}
