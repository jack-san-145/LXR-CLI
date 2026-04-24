package commands

import (
	"log"
	"lxr-cli/client"
	"lxr-cli/response"
)

func Ping() {
	cli := client.CreateClient()
	res, err := cli.Get("http://lxr/ping")
	if err != nil {
		log.Fatal("Request Error: ", err)
	}

	content, err := response.GetResponse(res)
	if err != nil {
		log.Fatal("Error :", err)
	}
	log.Println(content)
}
