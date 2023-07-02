# cxcli completion zsh

Generate the autocompletion script for zsh

## Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(cxcli completion zsh)

To load completions for every new session, execute once:

### Linux:

	cxcli completion zsh > "${fpath[1]}/_cxcli"

### macOS:

	cxcli completion zsh > $(brew --prefix)/share/zsh/site-functions/_cxcli

You will need to start a new shell for this setup to take effect.


```
cxcli completion zsh [flags]
```

## Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file path (optional)
  -o, --output-format string   Output Format. Options: text, json. Default: text (optional) (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command (optional)
  -v, --verbose                verbose error output (with stack trace) (optional)
```

## See also

* [cxcli completion](/cmd/cxcli_completion/)	 - Generate the autocompletion script for the specified shell

