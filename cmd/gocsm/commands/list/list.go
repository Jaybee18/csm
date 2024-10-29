/*
Copyright © 2024 Jan Bessler jbessler.business@gmail.com
*/
package list

import (
	"fmt"

	"github.com/Jaybee18/gocsm/internal/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func listAliases(cmd *cobra.Command, args []string) error {
	aliases, err := utils.GetAliases()
	if err != nil {
		return err
	}

	t := table.NewWriter()

	t.AppendHeader(table.Row{"Alias", "Command", "Description"})

	t.SetColumnConfigs([]table.ColumnConfig{
		{Name: "Description", WidthMax: 20},
	})

	for _, alias := range aliases {
		t.AppendRow(table.Row{
			alias.Name,
			alias.Command,
			alias.Description,
		})
	}

	fmt.Println(t.Render())

	return nil
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "list",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		RunE: listAliases,
		Args: cobra.NoArgs,
	}

	return command
}
