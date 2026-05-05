package models

type NewContainer struct {
	Image         string `json:"image_name"`
	ContainerName string `json:"container_name"`
}

type CreationResponse struct {
	ContainerId   string `json:"container_id"`
	ContainerName string `json:"container_name"`
	IsCreated     bool   `json:"is_created"`
	AlreadyExists bool   `json:"already_exists"`
}

type ContainerRunResponse struct {
	Active bool `json:"active"`
}

type ImagePullResponse struct {
	Status string `json:"status"`
}
