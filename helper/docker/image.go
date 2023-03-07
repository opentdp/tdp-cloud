package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/docker/docker/api/types"
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
		authConfig, _ := json.Marshal(types.AuthConfig{
			Username: rq.Username,
			Password: rq.Password,
		})
		option.RegistryAuth = base64.StdEncoding.EncodeToString(authConfig)
	}

	return dc.Client.ImagePull(ctx, rq.ImageName, option)

}
