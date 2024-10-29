/*
Copyright © 2024 Jan Bessler jbessler.business@gmail.com
*/
package config

import (
	"github.com/spf13/cobra"
)

func configCommand(cmd *cobra.Command, args []string) error {
	return nil
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "config",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		RunE: configCommand,
	}

	editCommand := &cobra.Command{
		Use:   "edit",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		RunE: configEditCommand,
	}

	command.AddCommand(editCommand)

	return command
}
