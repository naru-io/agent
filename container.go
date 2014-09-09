package agent

import (
	dockerapi "github.com/fsouza/go-dockerclient"
)

type Container struct {
	dockerapi.Container
	Tags  []string
	Attrs map[string]string
}

func NewContainer() *Container {
	return &Container{}
}
