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

func main() {
	hostID := "testhost1"
	log.Println("main")

	etcdClient := etcd.NewClient([]string{S_URL})

	path := "/hosts/" + hostID + "/containers/_future/pull/"

	opts := &dockerapi.PullImageOptions{
		Repository: "busybox",
		Tag:        "latest",
	}

	value, _ := json.Marshal(opts)

	resp, err := etcdClient.CreateInOrder(path, string(value), 0)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}
