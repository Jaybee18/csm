# gocsm
Go CLI script management tool

## Installation
```sh
go install github.com/Jaybee18/gocsm@latest
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
$ ./bin/gocsm
```

## Configure autocompletion
Powershell
```powershell
# Only for the current shell:
PS> gocsm completion powershell | Out-String | Invoke-Expression

# For all shells, source this file from the powershell profile:
PS> gocsm completion powershell > gocsm.ps1
```

Bash
```sh
# Only for the current shell
$ source <(gocsm completion bash)

# For all shells:
# Linux:
$ gocsm completion bash > /etc/bash_completion.d/gocsm
# macOS:
$ gocsm completion bash > $(brew --prefix)/etc/bash_completion.d/gocsm
```

For more information, see the `completion` subcommand and its available subcommands:
```sh
$ gocsm completion --help
# then
$ gocsm completion <command> --help
```
