package response

import (
	"bufio"
	"encoding/json"
	"lxr-cli/models"
	"net/http"
)

func GetResponse(res *http.Response) (string, error) {

	defer res.Body.Close()

	res_reader := bufio.NewReader(res.Body)
	content, err := res_reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return content, nil

}

func GetContainerCreationResponse(res *http.Response) (*models.CreationResponse, error) {

	defer res.Body.Close()

	var response models.CreationResponse

	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil

}
func GetContainerRunResponse(res *http.Response) (*models.ContainerRunResponse, error) {

	defer res.Body.Close()

	var response models.ContainerRunResponse

	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil

}
