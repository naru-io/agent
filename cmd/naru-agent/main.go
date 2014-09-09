package main

import (
	"flag"
	dockerapi "github.com/fsouza/go-dockerclient"
	"log"
	"net/url"
	"os"

//    "bitbucket.org/naru-io/agent"
)

func getOpt(name, def string) string {
	if env := os.Getenv(name); env != "" {
		return env
	}
	return def
}

func main() {
	flag.Parse()

	uri, uri_err := url.Parse(flag.Arg(0))
	if uri_err != nil {
		log.Fatal("agent:", uri_err)
	}

	docker, err := dockerapi.NewClient(getOpt("DOCKER_HOST", "unix:///var/run/docker.sock"))
	if err != nil {
		log.Fatal("agent:", err)
	}

	events := make(chan *dockerapi.APIEvents)
	if docker.AddEventListener(events) != nil {
		log.Fatal("agent:", err)
	}

	log.Println("Starting agent")

	for msg := range events {
		switch msg.Status {
		case "start":
			log.Println("events: start", msg)
		case "die":
			log.Println("events: die", msg)

		}
	}

	log.Fatal("agent: docker event loop closed") //TODO: reconnect?
}
