package agent

import (
	"github.com/coreos/go-etcd/etcd"
	"net/url"
)

type EtcdStorage struct {
	client *etcd.Client
	path   string
}

func NewEtcdStorage(uri *url.URL) Storage {
	urls := make([]string, 0)
	if uri.Host != "" {
		urls = append(urls, "http://"+uri.Host)
	}

	return &EtcdStorage{client: etcd.NewClient(urls), path: uri.Path}
}

func (e *EtcdStorage) Add(container *Container) error {
	return nil
}

func (e *EtcdStorage) Remove(container *Container) error {
	return nil
}

func (e *EtcdStorage) Update(container *Container) error {
	return nil
}

func (e *EtcdStorage) AddListener(name string, listener Listener) {

}
