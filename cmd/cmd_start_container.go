package main

import (
	"encoding/json"
	"github.com/coreos/go-etcd/etcd"
	dockerapi "github.com/fsouza/go-dockerclient"
	"log"
)

const (
	S_URL = "http://etcd_1:4001"
)

type StartContainerOptions struct {
	ID         string
	HostConfig *dockerapi.HostConfig
}

func main() {
	hostID := "testhost1"
	log.Println("main")

	etcdClient := etcd.NewClient([]string{S_URL})

	path := "/hosts/" + hostID + "/containers/_future/start/"

	config := &dockerapi.HostConfig{
		PublishAllPorts: true,
	}
	opts := StartContainerOptions{
		ID:         "bbaffb6866e7f3cbe4479d39656074adf68e81540cb6ac8bac8ac1fceac5d873",
		HostConfig: config,
	}

	value, _ := json.Marshal(opts)

	resp, err := etcdClient.CreateInOrder(path, string(value), 0)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}
