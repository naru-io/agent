package agent

import (
	"log"
	"net/url"
)

type Listener func(key string, value string)

type Storage interface {
	AddHost(host *Host) error
	AddContainer(container *Container) error
	RemoveContainer(container *Container) error
	UpdateContainer(container *Container) error
	GetContainerIdsByHost(host *Host) ([]string, error)
	AddListener(name string, listener Listener)

	Set(key string, value string) error
	Get(key string) string
}

func NewStorage(uri *url.URL) Storage {
	factory := map[string]func(*url.URL) Storage{
		"etcd": NewEtcdStorage,
	}[uri.Scheme]
	if factory == nil {
		log.Fatal("agent: unrecognized storage backend: ", uri.Scheme)
	}

	log.Println("agent: Using ", uri.Scheme, " agent storage backend at", uri)
	return factory(uri)
}
