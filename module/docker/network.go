package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
)

// 创建网络

func (dc *DockerClient) NetworkCreate(name string, rq *types.NetworkCreate) (string, error) {

	ctx := context.Background()

	network, err := dc.Client.NetworkCreate(ctx, name, *rq)
	if err != nil {
		return "", err
	}

	return network.ID, nil

}

// 网络详情

func (dc *DockerClient) NetworkInspect(networkID string) (types.NetworkResource, error) {

	ctx := context.Background()

	network, err := dc.Client.NetworkInspect(ctx, networkID, types.NetworkInspectOptions{})
	return network, err

}

// 列出网络

func (dc *DockerClient) NetworkList() ([]types.NetworkResource, error) {

	ctx := context.Background()

	networks, err := dc.Client.NetworkList(ctx, types.NetworkListOptions{})
	return networks, err

}

// 删除网络

func (dc *DockerClient) NetworkRemove(networkID string) error {

	ctx := context.Background()

	return dc.Client.NetworkRemove(ctx, networkID)

}

// 查找网络

func (dc *DockerClient) NetworkSearch(networkName string) ([]types.NetworkResource, error) {

	ctx := context.Background()

	networks, err := dc.Client.NetworkList(ctx, types.NetworkListOptions{
		Filters: filters.NewArgs(
			filters.Arg("name", networkName),
		),
	})
	return networks, err

}

// 网络连接

func (dc *DockerClient) NetworkConnect(networkID, containerID string) error {

	ctx := context.Background()

	return dc.Client.NetworkConnect(ctx, networkID, containerID, nil)

}

// 网络断开

func (dc *DockerClient) NetworkDisconnect(networkID, containerID string) error {

	ctx := context.Background()

	return dc.Client.NetworkDisconnect(ctx, networkID, containerID, true)

}

// 网络查找容器

func (dc *DockerClient) NetworkSearchContainer(networkID string) ([]types.Container, error) {

	ctx := context.Background()

	containers, err := dc.Client.ContainerList(ctx, container.ListOptions{
		Filters: filters.NewArgs(
			filters.Arg("network", networkID),
		),
	})
	return containers, err

}
