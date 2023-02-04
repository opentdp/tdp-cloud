package global

import (
	"github.com/spf13/cobra"

	"tdp-cloud/cmd/args"
)

var cli = &cobra.Command{
	Use:   "tdp-cloud",
	Short: "TDP Cloud",
	Long:  args.ReadmeText,
}

func WithCli() *cobra.Command {

	cli.PersistentFlags().StringVarP(&args.ConfigFile, "config", "c", args.ConfigFile, "config file")

	return cli

}
