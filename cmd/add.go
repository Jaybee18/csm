/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Jaybee18/gocsm/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func addCommand(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Too few arguments")
	}

	name := args[0]
	command := args[1]
	description, _ := cmd.Flags().GetString("description")

	aliases, err := utils.GetAliases()
	if err != nil {
		return err
	}

	if !utils.GetAliasWithName(name).IsEmpty() {
		return fmt.Errorf("Command with alias \"%s\" already exists", name)
	}

	alias := &utils.Alias{
		Name:        name,
		Command:     command,
		Description: description,
	}

	aliases = append(aliases, alias)

	viper.Set("aliases", aliases)
	viper.WriteConfig()
	return nil
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <alias> <command>",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: addCommand,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("description", "d", "", "Description for the command")
}
