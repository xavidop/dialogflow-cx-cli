# cxcli agent restore

Executes a restore action for a specific agent

```
cxcli agent restore [agent-name] [flags]
```

## Options

```
  -h, --help                 help for restore
  -i, --input string         Input file name. Default: agent.blob (optional) (default "agent.blob")
  -l, --location-id string   Dialogflow CX Location ID of the Project (required)
  -p, --project-id string    Dialogflow CX Project ID (required)
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file path (optional)
  -o, --output-format string   Output Format. Options: text, json. Default: text (optional) (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command (optional)
  -v, --verbose                verbose error output (with stack trace) (optional)
```

## See also

* [cxcli agent](/cmd/cxcli_agent/)	 - Actions on agent commands

