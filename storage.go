package agent

import (
	"log"
	"net/url"
)

type Listener func(name string, container *Container)

type Storage interface {
	Add(container *Container) error
	Remove(container *Container) error
	Update(container *Container) error
	AddListener(name string, listener Listener)
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
