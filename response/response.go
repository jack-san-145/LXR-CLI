package response

import (
	"bufio"
	"fmt"
	"net/http"
)

func GetResponse(res *http.Response) (string, error) {

	defer res.Body.Close()

	res_reader := bufio.NewReader(res.Body)
	content, err := res_reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Response from LXR daemon: %v", content), nil

}
