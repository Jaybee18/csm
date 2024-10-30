<div align="center">

# csm
![GoVersion](https://img.shields.io/github/go-mod/go-version/Jaybee18/csm)
![Version](https://img.shields.io/github/v/release/Jaybee18/csm)
![License](https://img.shields.io/github/license/Jaybee18/csm)
![Stars](https://img.shields.io/github/stars/Jaybee18/csm)

Go CLI script management tool for linux, macOS and windows with auto-completion and command arguments
</br>
</br>
</div>

Heavily inspired by [pier](https://github.com/pier-cli/pier) but since I can't understand and generally dislike Rust, it was of no use for me. So I made a similar, but simpler, tool in Go.

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

# or with arguments
$ csm add second-alias "echo %0" -d "print the first argument"
```
Using a command:
```sh
$ csm run my-alias
this is my test command

# or with arguments
$ csm run second-alias printme
printme
```
To see the list of configured commands:
```
$ csm ls
╭──────────┬──────────────────────────────┬─────────────────────────╮
│ ALIAS    │ COMMAND                      │ DESCRIPTION             │
├──────────┼──────────────────────────────┼─────────────────────────┤
│ my-alias │ echo this is my test command │ an optional description │
╰──────────┴──────────────────────────────┴─────────────────────────╯
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
