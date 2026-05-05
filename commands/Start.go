package commands

import (
	"fmt"

	"lxr-cli/client"
	"lxr-cli/response"
)

func Start(container_name string) {

	cli := client.CreateClient()

	res, err := cli.Get("http://lxr/start?container_name=" + container_name)
	content, err := response.GetContainerRunResponse(res)
	if err != nil || !content.Active {
		//fmt.Println("Container Run Error:", err)
		fmt.Println("container not found")
		return
	}

	fmt.Println(container_name, ": Started")
}
