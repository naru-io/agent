package agent

import (
	dockerapi "github.com/fsouza/go-dockerclient"
)

type Container struct {
	dockerapi.APIContainers
	Action string           `json:"action,omitempty`
	Config dockerapi.Config `json:"config,omitempty`
	Tags   []string
	Attrs  map[string]string
}

func NewContainer(apiContainer *dockerapi.APIContainers) *Container {
	container := &Container{*apiContainer, "", dockerapi.Config{}, nil, nil}
	return container
}
