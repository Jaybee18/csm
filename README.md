# csm
Go CLI script management tool

## Installation
```sh
go install github.com/Jaybee18/csm@latest
```
After installing you have to either manually create the config file at `$HOME/.csm.yaml` or use the config init command:
```
$ csm config init
```

## Usage
Adding a new command:
```sh
$ csm add my-alias "echo this is my test command" -d "an optional description"
```
Using a command:
```
$ csm run my-alias
this is my test command
```
To see the list of configured commands:
```
$ csm ls
╭──────────┬─────────────────────────────┬─────────────────────────╮
│ ALIAS    │ COMMAND                     │ DESCRIPTION             │
├──────────┼─────────────────────────────┼─────────────────────────┤
│ my-alias │ lsof -i -P -n | grep LISTEN │ an optional description │
╰──────────┴─────────────────────────────┴─────────────────────────╯
```
To view help:
```
$ csm --help
A script/command management tool

Usage:
  csm [command]

Available Commands:
  add         Add a command under the given alias
  completion  Generate the autocompletion script for the specified shell
  config      Configure the config file for csm
  help        Help about any command
  list        List all stored aliases
  remove      Remove the entry with the given alias
  run         Run the command stored under the given alias

Flags:
      --config string   config file (default is $HOME/.csm.yaml)
  -h, --help            help for csm
      --shell strings   shell for commands (default [/bin/sh,-c])
  -t, --toggle          Help message for toggle
  -v, --version         version for csm

Use "csm [command] --help" for more information about a command.
```

## Development
Build with
```sh
$ make build
```
Run with
```sh
$ go run main.go

# or after building
$ ./bin/csm
```

## Configure autocompletion
Powershell
```powershell
# Only for the current shell:
PS> csm completion powershell | Out-String | Invoke-Expression

# For all shells, source this file from the powershell profile:
PS> csm completion powershell > csm.ps1
```

Bash
```sh
# Only for the current shell
$ source <(csm completion bash)

# For all shells:
# Linux:
$ csm completion bash > /etc/bash_completion.d/csm
# macOS:
$ csm completion bash > $(brew --prefix)/etc/bash_completion.d/csm
```

For more information, see the `completion` subcommand and its available subcommands:
```sh
$ csm completion --help
# then
$ csm completion <command> --help
```
