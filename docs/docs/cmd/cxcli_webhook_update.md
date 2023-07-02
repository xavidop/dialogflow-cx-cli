# cxcli webhook update

update a webhook

```
cxcli webhook update [name] [flags]
```

## Options

```
  -a, --agent-name string    Dialogflow CX Agent Name (required)
  -e, --environment string   Environment where you want to set the webhook url. Default: global (optional) (default "global")
  -h, --help                 help for update
  -l, --location-id string   Dialogflow CX Location ID of the Project (required)
  -p, --project-id string    Dialogflow CX Project ID (required)
  -r, --url string           Webhook URL (required)
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

