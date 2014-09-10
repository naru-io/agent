package agent

import (
	"encoding/json"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"log"
	"net/url"
)

const (
	HOST_STORAGE_URL = "/hosts"
)

type EtcdStorage struct {
	Client *etcd.Client
	path   string
}

func NewEtcdStorage(uri *url.URL) Storage {
	urls := make([]string, 0)
	if uri.Host != "" {
		urls = append(urls, "http://"+uri.Host)
	}

	return &EtcdStorage{Client: etcd.NewClient(urls), path: uri.Path}
}

func (e *EtcdStorage) AddHost(host *Host) error {
	hostId := host.ID
	path := fmt.Sprintf("%s/%s/_host", HOST_STORAGE_URL, hostId)

	jsonstr, jerr := json.Marshal(host)
	if jerr != nil {
		return jerr
	}

	_, err := e.Client.SetDir(fmt.Sprintf("%s/%s", HOST_STORAGE_URL, hostId), 0)
	if err != nil {
		log.Println(err)
	}

	_, err = e.Client.Set(path, string(jsonstr), 0)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (e *EtcdStorage) AddContainer(container *Container) error {
	return nil
}

func (e *EtcdStorage) RemoveContainer(container *Container) error {
	return nil
}

func (e *EtcdStorage) UpdateContainer(container *Container) error {
	return nil
}

func (e *EtcdStorage) AddListener(name string, listener Listener) {

}
