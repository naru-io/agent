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

	dockerContainers, err := m.Docker.ListContainers(dockerapi.ListContainersOptions{})
	if err != nil {
		log.Println("agent: ", err)
		return err
	}

	containers := make([]*Container, 0)

	for _, c := range dockerContainers {
		dockerContainer, _ := m.Docker.InspectContainer(c.ID)
		container := NewContainer(dockerContainer)
		containers = append(containers, container)
		m.Storage.AddContainer(container)
	}

	m.Storage.AddListener("create", m.CreateContainerListener)
	m.Storage.AddListener("stop", m.StopContainerListener)
	return nil
}

func (m *Manager) CreateContainerListener(key string, value string) {
	log.Println("CreateContainer:", key, value)
}

func (m *Manager) StopContainerListener(key string, value string) {
	log.Println("StopContainer:", key, value)
}
