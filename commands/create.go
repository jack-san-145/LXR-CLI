package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"lxr-cli/client"
	"lxr-cli/models"
	"lxr-cli/response"
)

func Create(args []string) {
	if len(args) != 3 || args[0] != "--name" {
		fmt.Println("Usage: lxr create --name <container_name> <image_name>")
		return
	}

	var con models.NewContainer

	con.ContainerName = args[1]
	con.Image = args[2]

	jsonData, err := json.Marshal(con)
	if err != nil {
		fmt.Println("Error Marshaling: ", err)
		return
	}

	cli := client.CreateClient()
	res, err := cli.Post("http://lxr/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Request Error (create): ", err)
	}

	response, err := response.GetJsonResponse(res)
	if err != nil || response.IsCreated == false {
		fmt.Println("Container creation Failured")
		return
	}

	fmt.Printf("Container Created: Name = %v , ID = %v", response.ContainerName, response.ContainerId)
}
