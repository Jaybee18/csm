/*
Copyright © 2024 Jan Bessler jbessler.business@gmail.com
*/
package list

import (
	"fmt"

	"github.com/Jaybee18/csm/internal/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func listAliases(cmd *cobra.Command, args []string) {
	aliases, err := utils.GetAliases()
	if err != nil {
		fmt.Println(err.Error())
		return
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
	return
}

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all stored aliases",
		Run:     listAliases,
		Args:    cobra.NoArgs,
	}

	return command
}
