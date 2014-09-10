package agent

import (
	"log"

	dockerapi "github.com/fsouza/go-dockerclient"
)

type Manager struct {
	Docker     *dockerapi.Client
	Storage    Storage
	Host       *Host
	Containers []*Container
}

func (m *Manager) Init() error {
	host := NewHost(m.Docker)
	serr := m.Storage.AddHost(host)
	if serr != nil {
		log.Printf("agent: ", serr)
		return serr
	}

	apiContainers, err := m.Docker.ListContainers(dockerapi.ListContainersOptions{})
	if err != nil {
		log.Println("agent: ", err)
		return err
	}

	containers := make([]*Container, 0)

	for _, c := range apiContainers {
		container := NewContainer(&c)
		containers = append(containers, container)
		m.Storage.AddContainer(container)
	}

	return nil
}
