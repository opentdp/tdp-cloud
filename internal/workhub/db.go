package workhub

import (
	"tdp-cloud/helper/json"
	"tdp-cloud/internal/dborm/machine"
)

func createMachine(node *Worker) {

	machine.Create(&machine.CreateParam{
		UserId:      node.UserId,
		VendorId:    0,
		HostName:    node.SystemStat.HostName,
		IpAddress:   node.Conn.RemoteAddr().String(),
		Region:      "",
		Model:       "worker",
		CloudId:     node.HostId,
		CloudMeta:   json.ToString(node.SystemStat),
		Description: "",
		Status:      "{}",
	})

}

func deleteMachine(node *Worker) {
}
