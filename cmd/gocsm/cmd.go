/*
Copyright © 2024 Jan Bessler jbessler.business@gmail.com
*/
package cmd

import (
	"github.com/Jaybee18/gocsm/cmd/gocsm/commands/add"
	"github.com/Jaybee18/gocsm/cmd/gocsm/commands/config"
	"github.com/Jaybee18/gocsm/cmd/gocsm/commands/list"
	"github.com/Jaybee18/gocsm/cmd/gocsm/commands/run"
)

func InitCommand() {
	rootCmd.AddCommand(add.NewCommand())
	rootCmd.AddCommand(config.NewCommand())
	rootCmd.AddCommand(list.NewCommand())
	rootCmd.AddCommand(run.NewCommand())
}
