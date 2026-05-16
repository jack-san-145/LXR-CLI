package commands

import (
	"fmt"
	"lxr-cli/client"
	"lxr-cli/response"
	"net/http"
)

func Kill(containerName string) {

	cli := client.CreateClient()

	req, _ := http.NewRequest(
		"DELETE",
		"http://lxr/kill?container_name="+containerName,
		nil,
	)

	res, err := cli.Do(req)
	if err != nil {
		fmt.Println("Kill Request Error: ", err)
		return
	}

	response, err := response.GetKillResponse(res)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(response)
	defer res.Body.Close()

}
