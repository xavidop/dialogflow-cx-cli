# cxcli environment execute-cicd

Executes a CI/CD pipeline for a specific environment

```
cxcli environment execute-cicd [environment] [flags]
```

## Options

```
  -a, --agent-name string    Dialogflow CX Agent Name (required)
  -h, --help                 help for execute-cicd
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

* [cxcli environment](/cmd/cxcli_environment/)	 - Actions on environment

