package commands

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"lxr-cli/client"
	"lxr-cli/models"
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
		fmt.Println("Request Error (create): ", err)
	}

	reader := bufio.NewReader(res.Body)
	defer res.Body.Close()
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Print(data)

	}

}
