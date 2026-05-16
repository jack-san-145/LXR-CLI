package commands

import (
	"fmt"
	"lxr-cli/client"
	"lxr-cli/response"
)

func Kill(containerName string) {

	cli := client.CreateClient()

	res, _ := cli.Get("http://lxr/kill?container_name=" + containerName)

	response, err := response.GetKillResponse(res)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(response)
}
