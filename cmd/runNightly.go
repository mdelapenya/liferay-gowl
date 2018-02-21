package cmd

import (
	"errors"

	date "github.com/mdelapenya/lpn/date"
	docker "github.com/mdelapenya/lpn/docker"
	liferay "github.com/mdelapenya/lpn/liferay"

	"github.com/spf13/cobra"
)

func init() {
	runCmd.AddCommand(runNightlyCmd)

	runNightlyCmd.Flags().IntVarP(&httpPort, "httpPort", "p", 8080, "Sets the HTTP port of Liferay Portal's bundle.")
	runNightlyCmd.Flags().BoolVarP(&enableDebug, "debug", "d", false, "Enables debug mode. (default false)")
	runNightlyCmd.Flags().IntVarP(&debugPort, "debugPort", "D", 9000, "Sets the debug port of Liferay Portal's bundle. It only applies if debug mode is enabled")
}

var runNightlyCmd = &cobra.Command{
	Use:   "nightly",
	Short: "Runs a Liferay Portal instance from Nightly Builds",
	Long: `Runs a Liferay Portal instance, obtained from Nightly Builds repository: ` + liferay.Nightlies + `.
	If no image tag is passed to the command, the tag representing the current date [` + date.CurrentDate + `] will be used.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("run requires zero or one argument representing the image tag to be run")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var tag string

		if len(args) == 0 {
			tag = date.CurrentDate
		} else {
			tag = args[0]
		}

		nightly := liferay.Nightly{}

		docker.RunDockerImage(
			nightly.GetRepository()+":"+tag, httpPort, enableDebug, debugPort)
	},
}
