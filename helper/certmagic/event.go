package certmagic

import (
	"context"

	"tdp-cloud/helper/logman"
)

var evtlog = logman.Named("cert.event").Sugar()

func magicEvent(ctx context.Context, evt string, data map[string]any) error {

	evtlog.Infof("Certmagic Event: %s with data: %v\n", evt, data)
	return nil

}
