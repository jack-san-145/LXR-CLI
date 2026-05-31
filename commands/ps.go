package commands

import (
	"fmt"
	"lxr-cli/client"
	"lxr-cli/response"
	"strconv"
	"strings"
)

func Ps() {

	cli := client.CreateClient()

	res, err := cli.Get("http://lxr/ps")
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
		fmt.Printf(
			"%-15s %-15s %-15s %-10s %-10s %-15s %-20s\n",
			"CONTAINER ID",
			"NAME",
			"IMAGE",
			"PID",
			"STATUS",
			"IP",
			"PORTS",
		)

		return
	}

	fmt.Printf(
		"%-15s %-15s %-15s %-10s %-10s %-15s %-20s\n",
		"CONTAINER ID",
		"NAME",
		"IMAGE",
		"PID",
		"STATUS",
		"IP",
		"PORTS",
	)

	for _, con := range containers {

		var ports []string

		for _, p := range con.Ports {
			ports = append(ports, strconv.Itoa(p))
		}

		portString := strings.Join(ports, ",")

		fmt.Printf(
			"%-15s %-15s %-15s %-10d %-10s %-15s %-20s\n",
			con.ContainerID,
			con.ContainerName,
			con.Image,
			con.PID,
			con.Status,
			con.IPAddress,
			portString,
		)

	}
}
