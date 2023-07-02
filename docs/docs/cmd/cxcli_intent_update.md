# cxcli intent update

Updates an intent in an agent

```
cxcli intent update [intent-name] [flags]
```

## Options

```
  -a, --agent-name string          Dialogflow CX Agent Name (required)
  -d, --description string         Description for this intent (optional)
  -h, --help                       help for update
  -e, --locale string              Locale of the intent. Default: agent locale (optional)
  -l, --location-id string         Dialogflow CX Location ID of the Project (required)
  -p, --project-id string          Dialogflow CX Project ID (required)
  -t, --training-phrases strings   List of the training phrases for this intent, comma separated. Entities, add @entity-type to the word: word@entity-type in the training phrase. Example: hello,hi how are you today@sys.date,morning!. (Optional)
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file path (optional)
  -o, --output-format string   Output Format. Options: text, json. Default: text (optional) (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command (optional)
  -v, --verbose                verbose error output (with stack trace) (optional)
```

## See also

* [cxcli intent](/cmd/cxcli_intent/)	 - Actions on intent commands

