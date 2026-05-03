package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lxr-cli/client"
	"lxr-cli/response"
	"strings"
	"time"
)

func Pull(image string) {
	image = strings.Split(image, ":")[0]
	cli := client.CreateClient()
	cli.Timeout = 10 * time.Minute
	jsonData, err := json.Marshal(map[string]string{"img_name": image})
	if err != nil {
		fmt.Println("Error Marshaling: ", err)
		return
	}

	res, err := cli.Post("http://lxr/pull_image", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error in Pull request: ", err)
		return
	}

	response, err := response.GetImagePullResponse(res)
	if err != nil {
		fmt.Println("Error in unmarshal image response:  ", err)
		return
	}
	fmt.Println(response.Status)

}
