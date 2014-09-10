package test

import (
	dockerapi "github.com/fsouza/go-dockerclient"
	"os"
	"testing"
)

func TestDockerAPIClient(t *testing.T) {
	docker, err := dockerapi.NewClient(os.Getenv("DOCKER_HOST"))
	if err != nil {
		t.Error(err)
	}
	if docker == nil {
		t.Error("docker is nil!")
	}

	env, derr := docker.Info()
	if derr != nil {
		t.Error(derr)
	}
	t.Logf("%#v", env)

	/*

	   containers, cerr := docker.ListContainers(dockerapi.ListContainersOptions{})
	   if cerr != nil {
	       t.Error(cerr)
	   }

	   for _, c := range containers {
	           t.Logf("%#v", c)
	   }
	*/

	images, ierr := docker.ListImages(false)
	if ierr != nil {
		t.Error(ierr)
	}

	for _, i := range images {
		t.Logf("%#v", i)
	}

}
