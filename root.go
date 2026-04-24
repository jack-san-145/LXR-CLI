package lxrcli

import (
	"log"
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
		//do nothing
	default:
		log.Fatal("Unknown command")
	}
}
