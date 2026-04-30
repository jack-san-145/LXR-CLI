package commands

import (
	"fmt"
	"io"
	"net"
	"os"
)

func Exec(containerName string) {
	conn, err := net.Dial("unix", "/var/run/lxr.sock")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer conn.Close()

	req := fmt.Sprintf(
		"GET /exec?container_name=%s HTTP/1.1\r\nHost: lxr\r\n\r\n",
		containerName,
	)

	conn.Write([]byte(req))

	go io.Copy(conn, os.Stdin) // user input → daemon
	io.Copy(os.Stdout, conn)   // daemon → user output

}
