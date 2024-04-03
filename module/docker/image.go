package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
)

// 拉取镜像

type ImagePullParam struct {
	ImageName string
	Username  string
	Password  string
}

func (dc *DockerClient) ImagePull(rq *ImagePullParam) (io.ReadCloser, error) {

	ctx := context.Background()
	option := types.ImagePullOptions{}

	if rq.Username != "" {
		authConfig, _ := json.Marshal(registry.AuthConfig{
			Username: rq.Username,
			Password: rq.Password,
		})
		option.RegistryAuth = base64.StdEncoding.EncodeToString(authConfig)
	}

	return dc.Client.ImagePull(ctx, rq.ImageName, option)

}

// 列出镜像

func (dc *DockerClient) ImageList() ([]image.Summary, error) {

	ctx := context.Background()

	images, err := dc.Client.ImageList(ctx, types.ImageListOptions{})
	return images, err

}

// 查找镜像

func (dc *DockerClient) ImageSearch(imageName string) ([]registry.SearchResult, error) {

	ctx := context.Background()

	images, err := dc.Client.ImageSearch(ctx, imageName, types.ImageSearchOptions{})
	return images, err

}

// 获取镜像信息

func (dc *DockerClient) ImageInspect(imageName string) (types.ImageInspect, error) {

	ctx := context.Background()

	image, _, err := dc.Client.ImageInspectWithRaw(ctx, imageName)
	return image, err

}

// 获取镜像历史

func (dc *DockerClient) ImageHistory(imageName string) ([]image.HistoryResponseItem, error) {

	ctx := context.Background()

	history, err := dc.Client.ImageHistory(ctx, imageName)
	return history, err

}

// 获取镜像标签

func (dc *DockerClient) ImageTag(imageName, tag string) error {

	ctx := context.Background()

	err := dc.Client.ImageTag(ctx, imageName, tag)
	return err

}

// 删除镜像

func (dc *DockerClient) ImageRemove(imageName string) error {

	ctx := context.Background()

	_, err := dc.Client.ImageRemove(ctx, imageName, types.ImageRemoveOptions{})
	return err

}

// 清理镜像

func (dc *DockerClient) ImagePrune() error {

	ctx := context.Background()

	_, err := dc.Client.ImagesPrune(ctx, filters.Args{})
	return err

}
