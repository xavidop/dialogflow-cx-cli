# cxcli completion fish

Generate the autocompletion script for fish

## Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	cxcli completion fish | source

To load completions for every new session, execute once:

	cxcli completion fish > ~/.config/fish/completions/cxcli.fish

You will need to start a new shell for this setup to take effect.


```
cxcli completion fish [flags]
```

## Options

```
  -h, --help              help for fish
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

