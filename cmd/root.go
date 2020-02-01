package cmd

import (
	"github.com/schlund/go-dusty/pkg/rpcserver"

	"github.com/spf13/cobra"
)

type cmd struct {
	clt     Parameters
	rootCmd *cobra.Command
}

func NewCommandTree() (*cobra.Command, error) {
	var rootCmd = &cobra.Command{
		Use:   "go-dusty",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			runDaemon, err := cmd.Flags().GetBool("daemonized")
			if err != nil {
				return err
			}

			if runDaemon {
				rpcserver.RunServer()
			} else {
				err := cmd.Help()
				return err
			}

			return nil
		},
	}

	c := &cmd{
		rootCmd: rootCmd,
	}

	c.addAssetsCmd()
	c.addBundlesCmd()
	c.addConfigCmd()
	c.addCpCmd()
	c.addDiskCmd()
	c.addDoctorCmd()
	c.addDumpCmd()
	c.addEnvCmd()
	c.addLogsCmd()
	c.addReposCmd()
	c.addRestartCmd()
	c.addScriptsCmd()
	c.addSetupCmd()
	c.addShellCmd()
	c.addShutdownCmd()
	c.addStatusCmd()
	c.addStopCmd()
	c.addTestCmd()
	c.addUpCmd()
	c.addUpgradeCmd()
	c.addValidateCmd()

	rootCmd.Flags().BoolP("daemonized", "d", false, "Run dusty daemon")

	return rootCmd, nil
}
