package cmd

import (
	"errors"

	liferay "github.com/mdelapenya/lpn/liferay"

	"github.com/spf13/cobra"
)

var enableDebug bool
var debugPort int
var httpPort int

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a Liferay Portal instance",
	Long: `Runs a Liferay Portal instance, obtained from the unofficial repositories: ` + liferay.Releases +
		` or ` + liferay.Nightlies + `.
		For that, please run this command adding "release" or "nightly" subcommands.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("run requires zero or one argument representing the image tag to be run")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		SubCommandInfo()
	},
}
