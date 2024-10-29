/*
Copyright © 2024 Jan Bessler jbessler.business@gmail.com
*/
package config

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
