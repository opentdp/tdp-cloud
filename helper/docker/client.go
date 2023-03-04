package docker

import (
	"github.com/docker/docker/client"
)

type DockerClient struct {
	*client.Client
}

func New(ops ...client.Opt) (*DockerClient, error) {

	client, err := client.NewClientWithOpts(ops...)
	return &DockerClient{client}, err

}
