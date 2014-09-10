package agent

import (
	dockerapi "github.com/fsouza/go-dockerclient"
)

type Container struct {
	dockerapi.Container
	Tags  []string
	Attrs map[string]string
}

func NewContainer(apiContainer *dockerapi.Container) *Container {
	container := &Container{*apiContainer, nil, nil}
	return container
}
