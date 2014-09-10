package agent

import (
	"encoding/json"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"log"
	"net/url"
	"strings"
	"time"
)

const (
	HOST_STORAGE_PATH      = "hosts"
	CONTAINER_STORAGE_PATH = "containers"
)

type EtcdStorage struct {
	Client         *etcd.Client
	Path           string
	HostPath       string
	ContainersPath string
}

func NewEtcdStorage(uri *url.URL) Storage {
	urls := make([]string, 0)
	if uri.Host != "" {
		urls = append(urls, "http://"+uri.Host)
	}

	return &EtcdStorage{Client: etcd.NewClient(urls), Path: uri.Path}
}

func (e *EtcdStorage) containerPath(containerId string) string {
	return e.ContainersPath + "/" + containerId

}

func (e *EtcdStorage) AddHost(host *Host) error {
	hostId := host.ID
	e.HostPath = fmt.Sprintf("%s%s/%s", e.Path, HOST_STORAGE_PATH, hostId)
	e.ContainersPath = e.HostPath + "/" + CONTAINER_STORAGE_PATH

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
	key := e.containerPath(container.ID)
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

//TODO:Deprecated
func (e *EtcdStorage) GetContainerIdsByHost(host *Host) ([]string, error) {
	containerIds := make([]string, 0)

	path := e.ContainersPath
	resp, err := e.Client.Get(path, false, true)
	if err != nil {
		return nil, err
	}

	for _, node := range resp.Node.Nodes {
		cid := strings.Replace(node.Key, e.ContainersPath+"/", "", 1)
		containerIds = append(containerIds, cid)
	}

	return containerIds, nil
}

func (e *EtcdStorage) Set(key string, value string) error {
	_, err := e.Client.Set(key, value, 0)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (e *EtcdStorage) Get(key string) string {
	resp, err := e.Client.Get(key, false, false)
	if err != nil {
		return ""
	}
	return resp.Node.Value
}

//TODO Need better naming
func (e *EtcdStorage) AddListener(name string, listener Listener) {
	path := fmt.Sprintf("%s/_future/%s", e.ContainersPath, name)

	watch := func() {
		log.Println("S:Watch", name)
		watchChannel := make(chan *etcd.Response)
		stopChannel := make(chan bool)
		go e.Client.Watch(path, 0, true, watchChannel, nil)
		for {
			resp, ok := <-watchChannel
			if !ok {
				break
			}
			log.Println("Listen", name)
			listener(resp.Node.Key, resp.Node.Value)
		}
		stopChannel <- true
		close(watchChannel)
		log.Println("E:Watch", name)
	}

	go func() {
		for {
			watch()
			time.Sleep(500 * time.Millisecond)
		}
	}()
}
