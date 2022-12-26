# cxcli completion powershell

Generate the autocompletion script for powershell

## Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	cxcli completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
cxcli completion powershell [flags]
```

## Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

## Options inherited from parent commands

```
  -c, --credentials string   verbose error output (with stack trace)
  -o, --output string        Output Format (default "text")
  -u, --skip-update-check    Skip the check for updates check run before every command
  -v, --verbose              verbose error output (with stack trace)
```

## See also

* [cxcli completion](/cmd/cxcli_completion/)	 - Generate the autocompletion script for the specified shell

