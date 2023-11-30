package worker

import (
	"time"

	"github.com/forgoer/openssl"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/psutil"
	"github.com/opentdp/go-helper/socket"

	"tdp-cloud/cmd/args"
)

func (pod *SendPod) Register() (uint, error) {

	var err error

	logman.Info("register:send")

	stat := psutil.Summary(true)
	cloudId := psutil.CloudInstanceId()
	workerId := openssl.Md5ToString(stat.HostId)

	err = pod.WriteJson(&socket.PlainData{
		Method: "Register",
		TaskId: 0,
		Payload: &map[string]any{
			"CloudId":       cloudId,
			"WorkerId":      workerId,
			"WorkerMeta":    stat,
			"WorkerVersion": args.Version,
		},
	})

	return 0, err

}

func (pod *RespPod) Register(rs *socket.PlainData) {

	logman.Info("register:resp", "payload", rs.Payload)

	go KeepAlive(&SendPod{pod.WsConn})

}

//// 持续报送状态

func KeepAlive(pod *SendPod) error {

	for {
		time.Sleep(25 * time.Second)

		if _, err := pod.Ping(); err != nil {
			logman.Error("ping:fail", "error", err)
			return err
		}
	}

}
