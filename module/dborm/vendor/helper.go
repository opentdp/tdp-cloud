package vendor

import (
	"regexp"
	"tdp-cloud/module/dborm"
)

// 密钥掩码

func SecretMask(lst []*dborm.Vendor) {

	re, _ := regexp.Compile(`^(\w{8}).+(\w{8})$`)

	for k, v := range lst {
		lst[k].SecretId = re.ReplaceAllString(v.SecretId, "$1*******$2")
		lst[k].SecretKey = re.ReplaceAllString(v.SecretKey, "$1******$2")
	}

}
