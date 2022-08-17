package json

import (
	"encoding/json"
)

// 序列化

func ToString(r any) string {

	b, _ := json.Marshal(r)

	return string(b)

}
