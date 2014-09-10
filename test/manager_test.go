package test

import (
	a "bitbucket.org/naru-io/agent"
	dockerapi "github.com/fsouza/go-dockerclient"

	"log"
	"net/url"
	"os"
	"testing"
)

func TestManagerRegisterHostAndContainers(t *testing.T) {
	docker, err := dockerapi.NewClient(os.Getenv("DOCKER_HOST"))
	if err != nil {
		log.Fatal("agent:", err)
	}

	etcdURL, _ := url.Parse("etcd://etcd_1:4001/")
	storage := a.NewStorage(etcdURL)

	manager := &a.Manager{
		Docker:  docker,
		Storage: storage,
	}

	err = manager.RegisterHostAndContainers()
	if err != nil {
		t.Error(err)
	}

}
