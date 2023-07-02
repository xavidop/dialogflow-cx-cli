# cxcli entity-type delete

Deletes an entity type in an agent

```
cxcli entity-type delete [entity-type-name] [flags]
```

## Options

```
  -a, --agent-name string    Dialogflow CX Agent Name (required)
  -f, --force string         Forces to delete the Entity type. NOTE: it will delete all any references to the entity type. Possible values: true or false (optional)
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

* [cxcli entity-type](/cmd/cxcli_entity-type/)	 - Actions on entity type commands

