package commands

import (
	"fmt"

	"lxr-cli/client"
	"lxr-cli/response"
)

func Start(container_name string) {

	cli := client.CreateClient()

	res, _ := cli.Get("http://lxr/start?container_name=" + container_name)
	response, err := response.GetStartResponse(res)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(container_name + " " + response)
}
