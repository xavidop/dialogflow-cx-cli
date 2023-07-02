# cxcli flow train

Trains a flow in an agent. If flow name is set to "all", it will train all flows in an agent

```
cxcli flow train [flow-name] [flags]
```

## Options

```
  -a, --agent-name string    Dialogflow CX Agent Name (required)
  -h, --help                 help for train
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

* [cxcli flow](/cmd/cxcli_flow/)	 - Actions on flow commands

