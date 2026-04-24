package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"net/http"
)

func main() {
	lxr_sock := "/var/run/lxr.sock"

	cli := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return net.Dial("unix", lxr_sock)
			},
		},
	}

	res, err := cli.Get("http://unix/ping")
	if err != nil {
		log.Println("Request Error: ", err)
		return
	}
	res_reader := bufio.NewReader(res.Body)
	content, err := res_reader.ReadString('\n')
	if err != nil {
		log.Println("Error : ", err)
	} else {
		log.Println("Response from server: ", content)
	}
	defer res.Body.Close()

}
