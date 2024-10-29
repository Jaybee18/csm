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

func configEditCommand(cmd *cobra.Command, args []string) {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		fmt.Println(fmt.Errorf("No editor configured; set $EDITOR env variable"))
	}

	configFile := viper.ConfigFileUsed()
	if configFile == "" {
		fmt.Println(fmt.Errorf("No config file found"))
	}

	command := exec.Command(editor, configFile)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = command.Wait()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
