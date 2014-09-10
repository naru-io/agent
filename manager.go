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

	m.Storage.AddListener(m.CreateContainerListener)
	return nil
}

func (m *Manager) CreateContainerListener(name string, container *Container) {
	action := container.Action
	if action == "" {
		log.Println("container's action empty")
		return
	}

	if action != "create" {
		log.Println("Currently only create container")
		return
	}

	imageName := container.Image
	images, err := m.Docker.ListImages(false)
	if err != nil {
		log.Println(err)
		return
	}

	ok := false
	for _, image := range images {
		for _, repoTag := range image.RepoTags {
			if imageName == repoTag {
				ok = true
				break
			}
		}
	}

	//Pull image
	if !ok {
		//Pull Image
		opts := dockerapi.PullImageOptions{Repository: imageName}
		auth := dockerapi.AuthConfiguration{}
		err = m.Docker.PullImage(opts, auth)
		if err != nil {
			log.Println(err)
			/*
				container.Status = "Failed Pull Image"
				value, _ := json.Marshal(container)
				//TODO: Use Storage's method!
				_, err = m.Storage.Client.Set(name, string(value), 0)
				if err != nil {
					log.Println(err)
					return
				}
			*/
		}
	}

	//Create container
	opts := dockerapi.CreateContainerOptions{Name: container.Names[0], Config: &container.Config}

	_container, err := m.Docker.CreateContainer(opts)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%#v", _container)
}
