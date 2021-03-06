package cmd

import (
	"strings"

	"github.com/drud/ddev/pkg/ddevapp"
	"github.com/drud/ddev/pkg/util"
	"github.com/spf13/cobra"
)

// DdevSSHCmd represents the ssh command.
var DdevSSHCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Starts a shell session in the container for a service. Uses web service by default.",
	Long:  `Starts a shell session in the container for a service. Uses web service by default. To start a shell session for another service, run "ddev ssh --service <service>`,
	Run: func(cmd *cobra.Command, args []string) {
		app, err := ddevapp.GetActiveApp("")
		if err != nil {
			util.Failed("Failed to ssh: %v", err)
		}

		if strings.Contains(app.SiteStatus(), ddevapp.SiteNotFound) {
			util.Failed("App not currently running. Try 'ddev start'.")
		}

		if strings.Contains(app.SiteStatus(), ddevapp.SiteStopped) {
			util.Failed("App is stopped. Run 'ddev start' to start the environment.")
		}

		app.DockerEnv()

		err = app.ExecWithTty(serviceType, "bash")

		if err != nil {
			util.Failed("Failed to ssh %s: %s", app.GetName(), err)
		}
	},
}

func init() {
	DdevSSHCmd.Flags().StringVarP(&serviceType, "service", "s", "web", "Defines the service to connect to. [e.g. web, db]")
	RootCmd.AddCommand(DdevSSHCmd)
}
