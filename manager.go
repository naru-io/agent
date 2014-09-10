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

func (m *Manager) RegisterHostAndContainers() error {
	host := NewHost(m.Docker)
	serr := m.Storage.AddHost(host)
	if serr != nil {
		log.Fatal("agent: ", serr)
	}

	apiContainers, err := m.Docker.ListContainers(dockerapi.ListContainersOptions{})
	if err != nil {
		log.Fatal("agent: ", err)
	}

	containers := make([]*Container, 0)

	for _, c := range apiContainers {
		container := NewContainer(&c)
		containers = append(containers, container)
		m.Storage.AddContainer(container)
	}
	log.Printf("%#v", containers)

	return nil
}
