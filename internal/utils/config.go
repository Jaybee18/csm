package utils

import (
	"fmt"
	"reflect"
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

func GetAliases() ([]*Alias, error) {
	var aliases []*Alias

	rawAliases, ok := viper.Get("aliases").([]interface{})
	if !ok {
		return aliases, fmt.Errorf("Could not cast aliases; type of: %s", reflect.TypeOf(viper.Get("aliases")).String())
	}

	for _, rawAlias := range rawAliases {
		rawAlias, ok := rawAlias.(map[string]interface{})
		if !ok {
			continue
		}

		name, ok := rawAlias["name"].(string)
		if !ok {
			name = ""
		}

		command, ok := rawAlias["command"].(string)
		if !ok {
			command = ""
		}

		description, ok := rawAlias["description"].(string)
		if !ok {
			description = ""
		}

		aliases = append(aliases, &Alias{
			Name:        name,
			Command:     command,
			Description: description,
		})
	}

	return aliases, nil
}

func GetAliasWithName(name string) *Alias {
	aliases, err := GetAliases()
	if err != nil {
		return &Alias{}
	}

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
