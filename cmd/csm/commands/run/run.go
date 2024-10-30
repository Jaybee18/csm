/*
Copyright © 2024 Jan Bessler jbessler.business@gmail.com
*/
package run

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Jaybee18/csm/utils"
	"github.com/spf13/cobra"
)

const VAR_DELIMITER = "%"

func runCommand(cmd *cobra.Command, args []string) {
	alias := utils.GetAliasWithName(args[0])
	if alias.IsEmpty() {
		fmt.Println(fmt.Errorf("Alias not found"))
		return
	}

	commandString := alias.Command

	cmdArgs := args[1:]
	for index, arg := range cmdArgs {
		searchString := fmt.Sprintf("%s%d", VAR_DELIMITER, index)
		if strings.Index(commandString, searchString) == -1 {
			fmt.Println(fmt.Sprintf("No argument %q found in command:\n%s", searchString, alias.Command))
			return
		}
		commandString = strings.Replace(commandString, searchString, arg, 1)
	}

	shell := utils.GetShell()

	command := exec.Command(shell[0], shell[1], commandString)
	stdout, err := command.Output()
	if err != nil {
		fmt.Println("Error running command: " + err.Error())
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
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: commandCompletion,
	}

	return command
}
