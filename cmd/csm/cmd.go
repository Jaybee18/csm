/*
Copyright © 2024 Jan Bessler jbessler.business@gmail.com
*/
package cmd

import (
	"github.com/Jaybee18/csm/cmd/csm/commands/add"
	"github.com/Jaybee18/csm/cmd/csm/commands/config"
	"github.com/Jaybee18/csm/cmd/csm/commands/list"
	"github.com/Jaybee18/csm/cmd/csm/commands/remove"
	"github.com/Jaybee18/csm/cmd/csm/commands/run"
)

func InitCommand() {
	rootCmd.AddCommand(add.NewCommand())
	rootCmd.AddCommand(config.NewCommand())
	rootCmd.AddCommand(list.NewCommand())
	rootCmd.AddCommand(remove.NewCommand())
	rootCmd.AddCommand(run.NewCommand())
}
