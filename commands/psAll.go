package commands

import (
	"fmt"
	"lxr-cli/client"
	"lxr-cli/response"
)

func PsAll() {

	cli := client.CreateClient()

	res, err := cli.Get("http://lxr/ps/all")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	containers, err := response.GetPsResponse(res)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(containers) == 0 {
		fmt.Println("No containers found")
		return
	}

	fmt.Printf(
		"%-15s %-15s %-15s %-10s %-10s %-15s %-10s\n",
		"CONTAINER ID",
		"NAME",
		"IMAGE",
		"PID",
		"STATUS",
		"IP",
		"PORT",
	)

	for _, con := range containers {

		fmt.Printf(
			"%-15s %-15s %-15s %-10d %-10s %-15s %-10d\n",
			con.ContainerID,
			con.ContainerName,
			con.Image,
			con.PID,
			con.Status,
			con.IPAddress,
			con.Port,
		)
	}
}
