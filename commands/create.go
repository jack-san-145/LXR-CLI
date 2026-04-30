package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lxr-cli/client"
	"lxr-cli/models"
	"lxr-cli/response"
)

func Create(args []string) {
	if len(args) != 4 || args[0] != "--name" || args[2] != "--image" {
		fmt.Println("Usage: lxr create --name <container_name> --image <image_name>")
		return
	}

	var con models.NewContainer

	con.ContainerName = args[1]
	con.Image = args[3]

	jsonData, err := json.Marshal(con)
	if err != nil {
		fmt.Println("Error Marshaling: ", err)
		return
	}

	cli := client.CreateClient()
	res, err := cli.Post("http://lxr/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Request Error (create): ", err)
	}

	response, err := response.GetContainerCreationResponse(res)
	if err != nil || response.IsCreated == false {
		fmt.Println("Container creation Failured")
		return
	}

	fmt.Printf("Container Created: Name = %v , ID = %v", response.ContainerName, response.ContainerId)
	fmt.Println("")
}
