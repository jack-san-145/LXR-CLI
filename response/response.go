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

func GetImagePullResponse(res *http.Response) (*models.ImagePullResponse, error) {
	defer res.Body.Close()

	var response models.ImagePullResponse

	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func GetStartResponse(res *http.Response) (string, error) {
	defer res.Body.Close()

	var response models.StartResponse

	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.AlreadyActive {
		return "Already Running", nil
	}
	if response.Activated {
		return "Started", nil
	}

	if response.Failed {
		return "Failed to Start", nil
	}

	return "Not Found", nil
}

func GetStopResponse(res *http.Response) (string, error) {

	defer res.Body.Close()

	var response models.StopResponse

	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if !response.Exists {
		return "Container doesn't Exists", nil
	}
	return "Container Stopped", nil
}

func GetKillResponse(res *http.Response) (string, error) {

	defer res.Body.Close()

	var response models.KillResponse

	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if !response.Exists {
		return "Container doesn't Exists", nil
	}

	if response.Killed {
		return "Container Killed Successfully", nil
	}
	return "Kill Failured", nil
}
