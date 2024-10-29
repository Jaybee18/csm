/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func addCommand(cmd *cobra.Command, args []string) error {
	configAliases, ok := viper.Get("aliases").(map[string][]string)
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
