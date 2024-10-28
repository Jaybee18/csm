package utils

import (
	"runtime"

	"github.com/spf13/viper"
)

type Alias struct {
	Name        string
	Command     string
	Description string
}

func (a *Alias) IsEmpty() bool {
	return a.Name == "" && a.Command == "" && a.Description == ""
}

func GetAliases() []*Alias {
	var aliases []*Alias

	configAliases := viper.GetStringMap("aliases")
	for name, alias := range configAliases {
		if alias, ok := alias.(map[string]interface{}); ok {

			command, ok := alias["command"].(string)
			if !ok {
				command = ""
			}

			description, ok := alias["description"].(string)
			if !ok {
				description = ""
			}

			aliases = append(aliases, &Alias{
				Name:        name,
				Command:     command,
				Description: description,
			})
		}
	}

	return aliases
}

func GetAliasWithName(name string) *Alias {
	aliases := GetAliases()

	for _, alias := range aliases {
		if alias.Name == name {
			return alias
		}
	}

	return &Alias{}
}

func GetShell() []string {
	shell := viper.GetStringSlice("shell")

	if shell != nil {
		return shell
	}

	if runtime.GOOS == "windows" {
		shell = []string{"cmd", "/c"}
	} else {
		shell = []string{"/bin/sh", "-c"}
	}

	return shell
}
