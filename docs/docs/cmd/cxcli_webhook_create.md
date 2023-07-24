# cxcli webhook create

create a webhook

```
cxcli webhook create [name] [flags]
```

## Options

```
  -a, --agent-name string           Dialogflow CX Agent Name (required)
  -e, --environment string          Environment where you want to set the webhook url. Default: global (optional) (default "global")
  -f, --flexible string             Creates a flexible webhook. Possible values: true or false (optional) (default "false")
  -h, --help                        help for create
  -l, --location-id string          Dialogflow CX Location ID of the Project (required)
  -m, --parameters-mapping string   Creates a parameter mapping for flexible webhook, comma separated. The format is parameter@json-path,paramter2@json-path2. Example: my-param@$.fully.qualified.path.to.field (required only if flexible is true)
  -p, --project-id string           Dialogflow CX Project ID (required)
  -t, --request-body string         Creates a request body for flexible webhook. It has to be in JSON Format (required only if flexible is true)
  -r, --url string                  Webhook URL (required)
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

