package passport

import (
	"tdp-cloud/module/dborm/domain"
	"tdp-cloud/module/dborm/keypair"
	"tdp-cloud/module/dborm/machine"
	"tdp-cloud/module/dborm/script"
	"tdp-cloud/module/dborm/vendor"
)

// 统计信息

func Summary(userId uint) map[string]any {

	domainCount, _ := domain.Count(&domain.FetchAllParam{UserId: userId})
	keypairCount, _ := keypair.Count(&keypair.FetchAllParam{UserId: userId})
	machineCount, _ := machine.Count(&machine.FetchAllParam{UserId: userId})
	scriptCount, _ := script.Count(&script.FetchAllParam{UserId: userId})
	vendorCount, _ := vendor.Count(&vendor.FetchAllParam{UserId: userId})

	return map[string]any{
		"Domain":  domainCount,
		"Keypair": keypairCount,
		"Machine": machineCount,
		"Script":  scriptCount,
		"Vendor":  vendorCount,
	}

}
