package passport

import (
	"tdp-cloud/model/certjob"
	"tdp-cloud/model/domain"
	"tdp-cloud/model/keypair"
	"tdp-cloud/model/machine"
	"tdp-cloud/model/script"
	"tdp-cloud/model/vendor"
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
