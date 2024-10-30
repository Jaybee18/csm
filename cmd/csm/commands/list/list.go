/*
Copyright © 2024 Jan Bessler jbessler.business@gmail.com
*/
package list

import (
	"fmt"
	"slices"

	"github.com/Jaybee18/csm/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

const CELL_PADDING = 1
const TABLE_BORDER_WIDTH = 1

/*
Returns the maximum width of the column at the given
index. Returns -1 when the column index is out of
range for at least one row.
*/
func GetColumnWidth(table [][]string, column int) int {
	maxWidth := 0
	for _, row := range table {
		if len(row) <= column {
			return -1
		}

		if len(row[column]) > maxWidth {
			maxWidth = len(row[column])
		}
	}
	return maxWidth + CELL_PADDING*2
}

func RowFromSlice(slice []string) table.Row {
	row := table.Row{}
	for _, element := range slice {
		row = append(row, element)
	}
	return row
}

/*
Prints the table of aliases with their commands and descriptions.
Shrinks the biggest column to fit the terminal width.
*/
func listAliases(cmd *cobra.Command, args []string) {
	aliases, err := utils.GetAliases()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	content := [][]string{{
		"Alias",
		"Command",
		"Description",
	}}
	for _, alias := range aliases {
		content = append(content, []string{
			alias.Name,
			alias.Command,
			alias.Description,
		})
	}

	columnWidths := []int{
		GetColumnWidth(content, 0),
		GetColumnWidth(content, 1),
		GetColumnWidth(content, 2),
	}

	noFormat, err := cmd.Flags().GetBool("no-format")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if noFormat {
		for _, row := range content {
			fmt.Println(fmt.Sprintf("%-*v%-*v%v", columnWidths[0], row[0], columnWidths[1], row[1], row[2]))
		}
		return
	}

	terminalWidth, _, err := term.GetSize(0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)
	t.AppendHeader(RowFromSlice(content[0]))
	for _, row := range content[1:] {
		t.AppendRow(RowFromSlice(row))
	}

	widestWidth := slices.Max(columnWidths)
	widestColumn := slices.Index(columnWidths, widestWidth)

	t.SetColumnConfigs([]table.ColumnConfig{
		{Name: content[0][widestColumn], WidthMax: terminalWidth - (utils.Sum(columnWidths) - widestWidth) - TABLE_BORDER_WIDTH*4 - CELL_PADDING*2},
	})

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

	command.Flags().BoolP("no-format", "n", false, "Disable fancy table formatting")

	return command
}
