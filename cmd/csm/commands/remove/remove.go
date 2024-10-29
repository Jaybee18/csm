/*
Copyright © 2024 Jan Bessler jbessler.business@gmail.com
*/
package remove

import (
	"fmt"

	"github.com/Jaybee18/csm/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func removeCommand(cmd *cobra.Command, args []string) {
	name := args[0]
	alias := utils.GetAliasWithName(name)

	if alias.IsEmpty() {
		fmt.Println(fmt.Errorf("Alias \"%s\" doesn't exist", name))
	}

	aliases, err := utils.GetAliases()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	aliases = utils.Remove(aliases, utils.GetAliasWithName(name))

	viper.Set("aliases", aliases)
	viper.WriteConfig()
	return
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:     "remove <alias>",
		Aliases: []string{"rm"},
		Short:   "Remove the entry with the given alias",
		Run:     removeCommand,
		Args:    cobra.ExactArgs(1),
	}

	return command
}
