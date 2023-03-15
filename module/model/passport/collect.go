package passport

import (
	"tdp-cloud/module/model/certjob"
	"tdp-cloud/module/model/domain"
	"tdp-cloud/module/model/keypair"
	"tdp-cloud/module/model/machine"
	"tdp-cloud/module/model/script"
	"tdp-cloud/module/model/vendor"
)

// 统计信息

func Summary(userId uint) map[string]any {

	certjobCount, _ := certjob.Count(&certjob.FetchAllParam{UserId: userId})
	domainCount, _ := domain.Count(&domain.FetchAllParam{UserId: userId})
	keypairCount, _ := keypair.Count(&keypair.FetchAllParam{UserId: userId})
	machineCount, _ := machine.Count(&machine.FetchAllParam{UserId: userId})
	scriptCount, _ := script.Count(&script.FetchAllParam{UserId: userId})
	vendorCount, _ := vendor.Count(&vendor.FetchAllParam{UserId: userId})

	return map[string]any{
		"Certjob": certjobCount,
		"Domain":  domainCount,
		"Keypair": keypairCount,
		"Machine": machineCount,
		"Script":  scriptCount,
		"Vendor":  vendorCount,
	}

}
