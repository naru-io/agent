package agent

import (
	"encoding/json"
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

	m.Storage.AddListener("pull", m.PullImageListener)
	m.Storage.AddListener("create", m.CreateContainerListener)
	m.Storage.AddListener("start", m.StartContainerListener)
	m.Storage.AddListener("stop", m.StopContainerListener)
	return nil
}

func (m *Manager) PullImageListener(key, value string) {
	log.Println("S:PullImageListener:", key, value)

	var opts dockerapi.PullImageOptions
	err := json.Unmarshal([]byte(value), &opts)

	err = m.Docker.PullImage(opts, dockerapi.AuthConfiguration{})
	if err != nil {
		log.Println(err)
	}
	log.Println("E:PullImageListener:", key, value)
}

func (m *Manager) CreateContainerListener(key string, value string) {
	log.Println("CreateContainer:", key, value)

	var opts dockerapi.CreateContainerOptions
	err := json.Unmarshal([]byte(value), &opts)
	if err != nil {
		log.Println("CreateContainer: Wrong value:", err)
		return
	}

	//TODO: Pull Image

	container, cerr := m.Docker.CreateContainer(opts)
	if cerr != nil {
		log.Println("CreateContainer: ", cerr)
	}

	err = m.Storage.AddContainer(NewContainer(container))
	if err != nil {
		log.Println(err)
	}
}

func (m *Manager) StartContainerListener(key string, value string) {
	log.Println("StartContainer:", key, value)

	var opts StartContainerOptions
	err := json.Unmarshal([]byte(value), &opts)
	if err != nil {
		log.Println("StartContainer: Wrong value ", err)
	}

	err = m.Docker.StartContainer(opts.ID, opts.HostConfig)
	if err != nil {
		log.Println(err)
	}

	//TODO _
	container, _ := m.Docker.InspectContainer(opts.ID)

	err = m.Storage.AddContainer(NewContainer(container))
	if err != nil {
		log.Println(err)
	}
}

func (m *Manager) StopContainerListener(key string, value string) {
	log.Println("StopContainer:", key, value)
}
