package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
)

type Alias struct {
	Name        string
	Command     string
	Description string
}

var aliases = []Alias{
	{
		Name:        "test",
		Command:     "echo test",
		Description: "Runs the given echo test command. This text is very long.",
	},
}

func print_test_table() error {
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

func run_command_temp(ctx *cli.Context) error {
	args := ctx.Args()
	if args.Len() != 1 {
		return fmt.Errorf("len args != 1; TODO")
	}
	arg := args.First()

	var alias Alias
	for _, a := range aliases {
		if a.Name == arg {
			alias = a
		}
	}

	fmt.Println(fmt.Sprintf("%+v\n", alias))

	return nil
}

func main() {
	app := &cli.App{
		Name:  "gcsm",
		Usage: "Manage scripts and commands",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "list configured aliases",
				Action: func(*cli.Context) error {
					return print_test_table()
				},
			},
			{
				Name:  "run",
				Usage: "run the command of the given alias",
				Action: func(ctx *cli.Context) error {
					return run_command_temp(ctx)
				},
				Aliases: []string{"r"},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
