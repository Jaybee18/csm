/*
Copyright © 2024 Jan Bessler jbessler.business@gmail.com
*/
package config

import (
	"fmt"
	"os"

	"github.com/Jaybee18/csm/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func configInitCommand(cmd *cobra.Command, args []string) {
	configPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	path := configPath + "/.csm.yaml"
	if _, err := os.Stat(path); err == nil {
		fmt.Println("Config file already exists at " + path)
	}

	_, err = os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	viper.Set("aliases", []utils.Alias{})
	viper.WriteConfig()
	return
}
