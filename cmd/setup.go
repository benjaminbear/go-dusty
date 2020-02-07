package cmd

import (
	"fmt"

	"github.com/benjaminbear/go-dusty/pkg/commands"

	"github.com/spf13/cobra"
)

func (c *cmd) addSetupCmd() {
	// setupCmd represents the setup command
	var setupCmd = &cobra.Command{
		Use:   "setup",
		Short: "Configure Dusty after installation",
		Long: `Run this command once after installation to set up
configuration values tailored to your system.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("setup called")

			err := commands.SetupDustyConfig(c.clt.username, c.clt.specsRepo, c.clt.vmMemory, c.clt.sshKeyPath, !c.clt.noUpdate)
			if err != nil {
				return err
			}

			return nil
		},
	}

	c.rootCmd.AddCommand(setupCmd)

	setupCmd.Flags().StringVar(&c.clt.username, "username", "", "User name of the primary Dusty client user.")
	setupCmd.Flags().StringVar(&c.clt.specsRepo, "specs-repo", "", "Repo where your Dusty specs are located.")
	setupCmd.Flags().IntVar(&c.clt.vmMemory, "vm-memory", 0, "Memory to assign to the Docker VM, in megabytes.")
	setupCmd.Flags().StringVar(&c.clt.sshKeyPath, "ssh-key-path", "", "Path to ssh private key for the used git repos.")
	setupCmd.Flags().BoolVar(&c.clt.noUpdate, "no-update", false, "Skip pulling managed repos at conclusion of setup.")
}
