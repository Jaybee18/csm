/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/Jaybee18/gcsm/internal/utils"
	"github.com/spf13/cobra"
)

func runCommand(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Too many args")
	}

	alias := utils.GetAliasWithName(args[0])
	if alias.IsEmpty() {
		return fmt.Errorf("Alias not found")
	}

	shell := utils.GetShell()

	command := exec.Command(shell[0], shell[1], alias.Command)
	stdout, err := command.Output()
	if err != nil {
		return err
	}

	fmt.Print(string(stdout[:]))

	return nil
}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: runCommand,
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
