/*
Copyright © 2024 Jan Bessler jbessler.business@gmail.com
*/
package run

import (
	"fmt"
	"os/exec"

	"github.com/Jaybee18/gocsm/internal/utils"
	"github.com/spf13/cobra"
)

func runCommand(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Too many args")
	}

	alias := utils.GetAliasWithName(args[0])
	if alias.IsEmpty() {
		return fmt.Errorf("Alias not found")
	}

	shell := utils.GetShell()

	command := exec.Command(shell[0], shell[1], alias.Command)
	stdout, err := command.Output()
	if err != nil {
		return err
	}

	fmt.Print(string(stdout[:]))

	return nil
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "run",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		RunE: runCommand,
		Args: cobra.ExactArgs(1),
	}

	return command
}
