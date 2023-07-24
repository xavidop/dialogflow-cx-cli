# cxcli webhook delete

delete a webhook

```
cxcli webhook delete [name] [flags]
```

## Options

```
  -a, --agent-name string    Dialogflow CX Agent Name (required)
  -f, --force string         Forces to delete the webhook and its references in environments and flows. Possible values: true or false (optional)
  -h, --help                 help for delete
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

* [cxcli webhook](/cmd/cxcli_webhook/)	 - Actions on webhook commands

