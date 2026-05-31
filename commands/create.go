package commands

import (
	// "bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"lxr-cli/client"
	"lxr-cli/models"
	"os"
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

	defer res.Body.Close()

	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		fmt.Println("stream error:", err)
	}

}
