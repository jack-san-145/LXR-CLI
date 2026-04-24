package client

import (
	"context"
	"net"
	"net/http"
)

func CreateClient() http.Client {
	lxr_sock := "/var/run/lxr.sock"

	return http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", lxr_sock)
			},
		},
	}
}
