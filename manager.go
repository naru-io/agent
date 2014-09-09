package agent

import (
	"log"

	dockerapi "github.com/fsouza/go-dockerclient"
)

type Manager struct {
	Docker  *dockerapi.Client
	Storage Storage
}

func (m *Manager) RegisterHostAndContainers() error {
	containers, err := m.Docker.ListContainers(dockerapi.ListContainersOptions{})
	if err != nil {
		log.Fatal("agent: ", err)
	}

	for _, container := range containers {
		log.Println(container)
	}

	return nil
}
