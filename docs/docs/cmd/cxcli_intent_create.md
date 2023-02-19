# cxcli intent create

Creates an intent in an agent

```
cxcli intent create [intent-name] [flags]
```

## Options

```
  -a, --agent-name string          Dialogflow CX Agent Name
  -d, --description string         Optional. Description for this intent
  -h, --help                       help for create
  -e, --locale string              Optional. Locale of the intent. Default: agent locale
  -l, --location-id string         Dialogflow CX Location ID of the Project
  -p, --project-id string          Dialogflow CX Project ID
  -t, --training-phrases strings   List of the training phrases for this intent, comma separated. Entities, add @entity-type to the word: word@entity-type in the training phrase. Example: hello,hi how are you today@sys.date,morning!
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file
  -o, --output-format string   Output Format (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command
  -v, --verbose                verbose error output (with stack trace)
```

## See also

* [cxcli intent](/cmd/cxcli_intent/)	 - Actions on intent commands

