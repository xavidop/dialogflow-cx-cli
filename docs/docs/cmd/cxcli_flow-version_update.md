# cxcli flow-version update

update a version

```
cxcli flow-version update [name] [flags]
```

## Options

```
  -a, --agent-name string    Dialogflow CX Agent Name (required)
  -d, --description string   Description for this version (optional)
  -h, --help                 help for update
  -l, --location-id string   Dialogflow CX Location ID of the Project (required)
  -p, --project-id string    Dialogflow CX Project ID (required)
  -s, --start-flow string    Start Flow name to create the version (required)
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file path (optional)
  -o, --output-format string   Output Format. Options: text, json. Default: text (optional) (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command (optional)
  -v, --verbose                verbose error output (with stack trace) (optional)
```

## See also

* [cxcli flow-version](/cmd/cxcli_flow-version/)	 - Actions on flow versions commands

