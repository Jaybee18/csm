/*
Copyright © 2024 Jan Bessler jbessler.business@gmail.com
*/
package config

import (
	"github.com/spf13/cobra"
)

func configCommand(cmd *cobra.Command, args []string) {
	return
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "config",
		Short: "Configure the config file for csm",
		Run:   configCommand,
		Args:  cobra.ExactArgs(1),
	}

	editCommand := &cobra.Command{
		Use:   "edit",
		Short: "Edit the csm config file with the default editor in $EDITOR",
		Run:   configEditCommand,
	}

	command.AddCommand(editCommand)

	initCommand := &cobra.Command{
		Use:   "init",
		Short: "Init an empty config file",
		Run:   configInitCommand,
	}

	command.AddCommand(initCommand)

	return command
}
