package subset

import (
	"github.com/open-tdp/go-helper/logman"
	"github.com/open-tdp/go-helper/upgrade"
	"github.com/spf13/cobra"

	"tdp-cloud/cmd/args"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update assistant",
	Long:  "TDP Cloud Update Assistant",
	Run: func(cmd *cobra.Command, rq []string) {
		ExecUpdate()
	},
}

func WithUpdate() *cobra.Command {

	return updateCmd

}

func ExecUpdate() error {

	err := upgrade.Apply(&upgrade.RequesParam{
		Server:  args.UpdateUrl,
		Version: args.Version,
	})

	if err == nil {
		logman.Info("Update Success")
	}

	return err

}
