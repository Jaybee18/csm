/*
Copyright © 2024 Jan Bessler jbessler.business@gmail.com
*/
package run

import (
	"fmt"
	"os/exec"

	"github.com/Jaybee18/csm/internal/utils"
	"github.com/spf13/cobra"
)

func runCommand(cmd *cobra.Command, args []string) {
	alias := utils.GetAliasWithName(args[0])
	if alias.IsEmpty() {
		fmt.Println(fmt.Errorf("Alias not found"))
	}

	shell := utils.GetShell()

	command := exec.Command(shell[0], shell[1], alias.Command)
	stdout, err := command.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout[:]))
	return
}

func commandCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	aliases, err := utils.GetAliases()
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveError
	}
	return utils.Map(aliases, func(alias *utils.Alias) string { return alias.Name }), cobra.ShellCompDirectiveNoFileComp
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:               "run",
		Short:             "Run the command stored under the given alias",
		Run:               runCommand,
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: commandCompletion,
	}

	return command
}
