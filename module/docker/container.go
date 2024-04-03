package docker

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

// 创建容器

type ContainerCreateParam struct {
	Name  string
	Image string
}

func (dc *DockerClient) ContainerCreate(rq *ContainerCreateParam) (string, error) {

	ctx := context.Background()

	resp, err := dc.Client.ContainerCreate(
		ctx,
		&container.Config{
			Image: rq.Image,
		},
		&container.HostConfig{},
		&network.NetworkingConfig{},
		&specs.Platform{},
		rq.Name,
	)
	return resp.ID, err

}

// 列出容器

func (dc *DockerClient) ContainerList() ([]types.Container, error) {

	ctx := context.Background()

	containers, err := dc.Client.ContainerList(ctx, container.ListOptions{})
	return containers, err

}

// 查找容器

func (dc *DockerClient) ContainerSearch(containerName string) ([]types.Container, error) {

	ctx := context.Background()

	containers, err := dc.Client.ContainerList(ctx, container.ListOptions{
		Filters: filters.NewArgs(
			filters.Arg("name", containerName),
		),
	})
	return containers, err

}

// 删除容器

func (dc *DockerClient) ContainerRemove(containerID string) error {

	ctx := context.Background()

	return dc.Client.ContainerRemove(ctx, containerID, container.RemoveOptions{})

}

// 启动容器

func (dc *DockerClient) ContainerStart(containerID string) error {

	ctx := context.Background()

	return dc.Client.ContainerStart(ctx, containerID, container.StartOptions{})

}

// 停止容器

func (dc *DockerClient) ContainerStop(containerID string) error {

	ctx := context.Background()

	return dc.Client.ContainerStop(ctx, containerID, container.StopOptions{})

}

// 查看容器日志

func (dc *DockerClient) ContainerLogs(containerID string) (io.ReadCloser, error) {

	ctx := context.Background()

	return dc.Client.ContainerLogs(ctx, containerID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})

}

// 查看容器状态

func (dc *DockerClient) ContainerInspect(containerID string) (types.ContainerJSON, error) {

	ctx := context.Background()

	return dc.Client.ContainerInspect(ctx, containerID)

}

// 查看容器进程

func (dc *DockerClient) ContainerTop(containerID string) (container.ContainerTopOKBody, error) {

	ctx := context.Background()

	return dc.Client.ContainerTop(ctx, containerID, []string{})

}

// 查看容器统计信息

func (dc *DockerClient) ContainerStats(containerID string) (types.ContainerStats, error) {

	ctx := context.Background()

	return dc.Client.ContainerStats(ctx, containerID, false)

}

// 进入容器

func (dc *DockerClient) ContainerAttach(containerID string) (types.HijackedResponse, error) {

	ctx := context.Background()

	return dc.Client.ContainerAttach(ctx, containerID, container.AttachOptions{
		Stream: true,
		Stdin:  true,
		Stdout: true,
		Stderr: true,
	})

}

// 交互式执行容器命令

func (dc *DockerClient) ContainerExecInteractive(containerID string, cmd []string) (types.HijackedResponse, error) {

	ctx := context.Background()

	exec, err := dc.Client.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		Cmd:          cmd,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
	})
	if err != nil {
		return types.HijackedResponse{}, err
	}

	return dc.Client.ContainerExecAttach(ctx, exec.ID, types.ExecStartCheck{
		Tty: true,
	})

}
