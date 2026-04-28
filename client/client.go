package client

import (
	"context"
	"net"
	"net/http"
	"time"
)

func CreateClient() http.Client {
	lxr_sock := "/var/run/lxr.sock"

	return http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", lxr_sock)
			},
		},
	}
}
