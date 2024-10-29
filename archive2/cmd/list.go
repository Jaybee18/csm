/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Jaybee18/gocsm/internal/config"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func listAliases(config *config.Config) error {
	if config == nil {
		return nil
	}

	aliasTable := table.NewWriter()
	aliasTable.AppendHeader(table.Row{"Alias", "Command", "Description"})
	aliasTable.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:     "Description",
			WidthMax: 30,
		},
	})

	for _, alias := range config.Aliases {
		aliasTable.AppendRow(table.Row{
			alias.Name,
			alias.Command,
			alias.Description,
		})
	}

	fmt.Println(aliasTable.Render())
	return nil
}

func cmdRun(cmd *cobra.Command, args []string) error {
	config := config.GetConfig()
	return listAliases(config)
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: cmdRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
