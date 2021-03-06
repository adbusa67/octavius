//Package list provides cli command to list all the created job
package list

import (
	client "octavius/internal/cli/client/metadata"
	daemon "octavius/internal/cli/daemon/metadata"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/printer"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// NewCmd create a command for list
func NewCmd(octaviusDaemon daemon.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Get job list",
		Long:  `Get job list will give available jobs in octavius`,
		Args:  cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			client := &client.GrpcClient{}
			jobList, err := octaviusDaemon.List(client)
			if err != nil {
				log.Error(err, "error when getting job list")
				printer.Println("error when getting job list", color.FgRed)
				return
			}
			if len(jobList.Jobs) == 0 {
				printer.Println("No jobs available", color.FgGreen)
				return
			}
			for _, job := range jobList.Jobs {
				printer.Println(job, color.FgGreen)
			}
		},
	}
}
