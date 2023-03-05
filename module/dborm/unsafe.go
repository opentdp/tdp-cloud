package dborm

import (
	"errors"
	"regexp"
	"strings"
)

var orderSafe = regexp.MustCompile(`^(\w+)( DESC)?$`)

func OrderSafe(data string) error {

	for _, v := range strings.Split(data, ",") {
		if !orderSafe.MatchString(v) {
			return errors.New("unsafe order")
		}
	}

	return nil

}
