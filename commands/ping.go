package commands

import (
	"fmt"
	"lxr-cli/client"
	"lxr-cli/response"
)

func Ping() {
	cli := client.CreateClient()
	res, err := cli.Get("http://lxr/ping")
	if err != nil {
		fmt.Println("Request Error: ", err)
		return
	}

	content, err := response.GetResponse(res)
	if err != nil {
		fmt.Println("Error :", err)
	}
	fmt.Print(content)
}
