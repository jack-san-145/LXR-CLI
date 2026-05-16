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
	case "start":
		commands.Start(os.Args[2])
	case "stop":
		commands.Stop(os.Args[2])
	case "ping":
		commands.Ping()
	case "create":
		commands.Create(os.Args[2:])
	case "exec":
		commands.Exec(os.Args[2])
	case "pull":
		commands.Pull(os.Args[2])
	case "kill":
		commands.Kill(os.Args[2])
	default:
		fmt.Println("Lxr: <Unknown command>")
	}
}
