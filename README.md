# csm
Go CLI script management tool

## Installation
```sh
go install github.com/Jaybee18/csm@latest
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
