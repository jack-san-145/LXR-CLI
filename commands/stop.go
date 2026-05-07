package commands

import (
	"fmt"
	"lxr-cli/client"
	"lxr-cli/response"
)

func Stop(container_name string) {

	cli := client.CreateClient()

	res, _ := cli.Get("http://lxr/stop?container_name=" + container_name)

	response, err := response.GetStopResponse(res)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(response)
}
