package test

import (
	a "bitbucket.org/naru-io/agent"

	dockerapi "github.com/fsouza/go-dockerclient"

	"net/url"
	"os"
	"testing"
)

const (
	ETCD_TEST_URL = "etcd://etcd_1:4001/"
)

func TestAddHostToStorage(t *testing.T) {
	docker, err := dockerapi.NewClient(os.Getenv("DOCKER_HOST"))
	if err != nil {
		t.Error(err)
	}

	host := a.NewHost(docker)
	etcdURL, _ := url.Parse(ETCD_TEST_URL)
	storage := a.NewStorage(etcdURL)

	err = storage.AddHost(host)
	if err != nil {
		t.Error(err)
	}

}
