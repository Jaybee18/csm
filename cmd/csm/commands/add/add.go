/*
Copyright © 2024 Jan Bessler jbessler.business@gmail.com
*/
package add

import (
	"fmt"

	"github.com/Jaybee18/csm/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func addCommand(cmd *cobra.Command, args []string) {
	name := args[0]
	command := args[1]
	description, _ := cmd.Flags().GetString("description")

	aliases, err := utils.GetAliases()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if !utils.GetAliasWithName(name).IsEmpty() {
		fmt.Println(fmt.Errorf("Command with alias \"%s\" already exists", name))
	}

	alias := &utils.Alias{
		Name:        name,
		Command:     command,
		Description: description,
	}

	aliases = append(aliases, alias)

	viper.Set("aliases", aliases)
	viper.WriteConfig()
	return
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "add <alias> <command>",
		Short: "Add a command under the given alias",
		Run:   addCommand,
		Args:  cobra.ExactArgs(2),
	}

	command.Flags().StringP("description", "d", "", "Description for the command")

	return command
}
