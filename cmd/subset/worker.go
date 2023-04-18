package subset

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-cloud/cmd/args"
	"tdp-cloud/service"
)

var workerAct string

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Manage the worker",
	Long:  "TDP Cloud Worker Management",
	Run: func(cmd *cobra.Command, rq []string) {
		args.SubCommand.Name = cmd.Name()
		args.SubCommand.Action = workerAct
		service.Control(args.SubCommand.Name, workerAct)
	},
}

func WithWorker() *cobra.Command {

	workerCmd.Flags().StringVarP(&workerAct, "service", "s", "", "management worker service")
	workerCmd.Flags().StringP("remote", "r", "", "register URL (e.g. ws://{domain}/wsi/{appid}/worker)")

	viper.BindPFlag("worker.remote", workerCmd.Flags().Lookup("remote"))

	return workerCmd

}
