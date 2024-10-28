package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func configEditCommand(cmd *cobra.Command, args []string) error {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		return fmt.Errorf("No editor configured; set $EDITOR env variable")
	}

	configFile := viper.ConfigFileUsed()
	if configFile == "" {
		return fmt.Errorf("No config file found")
	}

	command := exec.Command(editor, configFile)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Start()
	if err != nil {
		return err
	}
	err = command.Wait()
	if err != nil {
		return err
	}

	return nil
}

// configEditCmd represents the config edit command
var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: configEditCommand,
}

func init() {
	configCmd.AddCommand(configEditCmd)
}
