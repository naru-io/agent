package agent

import (
	dockerapi "github.com/fsouza/go-dockerclient"

	"log"
	"os"
)

const (
	HOST_STATE_RUNNING = "running"
	HOST_STATE_DOWN    = "down"
)

type Host struct {
	ID              string `json:"id"`
	Name            string `json:"name,omitempty"`
	OperatingSystem string `json:"operatingSystem,omitempty"`
	KernelVersion   string `json:kernelVersion,omitempty"`
	State           string `json:state`
}

func NewHost(docker *dockerapi.Client) *Host {
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("hostname can't read")
		hostname = "-" //TODO
	}

	operatingSystem := ""
	kernelVersion := ""

	env, envErr := docker.Info()
	if envErr != nil {
		log.Printf("docker error ", envErr)
	}

	operatingSystem = env.Get("OperatingSystem")
	kernelVersion = env.Get("KernelVersion")

	host := &Host{
		ID:              hostname,
		Name:            hostname,
		State:           HOST_STATE_RUNNING,
		KernelVersion:   kernelVersion,
		OperatingSystem: operatingSystem,
	} //TODO: more precise stable ID

	return host
}
