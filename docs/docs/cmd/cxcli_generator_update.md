# cxcli generator update

update a generator

```
cxcli generator update [name] [flags]
```

## Options

```
  -a, --agent-name string    Dialogflow CX Agent Name (required)
  -h, --help                 help for update
  -e, --locale string        Locale of the intent. Default: agent locale (optional)
  -l, --location-id string   Dialogflow CX Location ID of the Project (required)
  -p, --project-id string    Dialogflow CX Project ID (required)
  -r, --prompt string        Prompt (required)
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file path (optional)
  -o, --output-format string   Output Format. Options: text, json. Default: text (optional) (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command (optional)
  -v, --verbose                verbose error output (with stack trace) (optional)
```

## See also

* [cxcli generator](/cmd/cxcli_generator/)	 - Actions on generator commands

