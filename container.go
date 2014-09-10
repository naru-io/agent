package agent

import (
	dockerapi "github.com/fsouza/go-dockerclient"
)

type Container struct {
	dockerapi.APIContainers
	Tags  []string
	Attrs map[string]string
}

func NewContainer(apiContainer *dockerapi.APIContainers) *Container {
	container := &Container{*apiContainer, nil, nil}
	return container
}
