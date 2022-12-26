# cxcli completion bash

Generate the autocompletion script for bash

## Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(cxcli completion bash)

To load completions for every new session, execute once:

### Linux:

	cxcli completion bash > /etc/bash_completion.d/cxcli

### macOS:

	cxcli completion bash > $(brew --prefix)/etc/bash_completion.d/cxcli

You will need to start a new shell for this setup to take effect.


```
cxcli completion bash
```

## Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file
  -o, --output-format string   Output Format (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command
  -v, --verbose                verbose error output (with stack trace)
```

## See also

* [cxcli completion](/cmd/cxcli_completion/)	 - Generate the autocompletion script for the specified shell

