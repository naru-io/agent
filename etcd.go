package agent

import (
	"encoding/json"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"log"
	"net/url"
)

const (
	HOST_STORAGE_PATH      = "hosts"
	CONTAINER_STORAGE_PATH = "containers"
)

type EtcdStorage struct {
	Client   *etcd.Client
	Path     string
	HostPath string
}

func NewEtcdStorage(uri *url.URL) Storage {
	urls := make([]string, 0)
	if uri.Host != "" {
		urls = append(urls, "http://"+uri.Host)
	}

	return &EtcdStorage{Client: etcd.NewClient(urls), Path: uri.Path}
}

func (e *EtcdStorage) AddHost(host *Host) error {
	hostId := host.ID
	e.HostPath = fmt.Sprintf("%s/%s/%s", e.Path, HOST_STORAGE_PATH, hostId)

	jsonstr, jerr := json.Marshal(host)
	if jerr != nil {
		return jerr
	}

	//TODO: Check and then SetDir
	_, err := e.Client.SetDir(e.HostPath, 0)
	if err != nil {
		log.Println("AddHost", err)
	}

	path := fmt.Sprintf("%s/_host", e.HostPath)
	_, err = e.Client.Set(path, string(jsonstr), 0)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (e *EtcdStorage) AddContainer(container *Container) error {
	key := fmt.Sprintf("%s/%s/%s", e.HostPath, CONTAINER_STORAGE_PATH, container.ID)
	value, err := json.Marshal(container)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = e.Client.Set(key, string(value), 0)
	if err != nil {
		log.Println(err)
		return err
	}

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
