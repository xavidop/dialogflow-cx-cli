# cxcli webhook update

update a webhook

```
cxcli webhook update [name] [flags]
```

## Options

```
  -a, --agent-name string    Dialogflow CX Agent Name
  -e, --environment string   Optional. Environment where you want to set the webhook url. Default: global (default "global")
  -h, --help                 help for update
  -l, --location-id string   Dialogflow CX Location ID of the Project
  -p, --project-id string    Dialogflow CX Project ID
  -r, --url string           Webhook URL
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file
  -o, --output-format string   Output Format (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command
  -v, --verbose                verbose error output (with stack trace)
```

## See also

* [cxcli webhook](/cmd/cxcli_webhook/)	 - Actions on webhook commands

