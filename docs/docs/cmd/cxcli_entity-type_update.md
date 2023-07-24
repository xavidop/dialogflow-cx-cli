# cxcli entity-type update

Updates an entity type in an agent

```
cxcli entity-type update [entity-type-name] [flags]
```

## Options

```
  -a, --agent-name string    Dialogflow CX Agent Name (required)
  -n, --entities strings     List of the entities for this entity type, comma separated. Format: entity1@synonym1|synonym2,entity2@synonym1|synonym2. Example: pikachu@25|pika,charmander@3 (required)
  -h, --help                 help for update
  -e, --locale string        Locale of the intent. Default: agent locale (optional)
  -l, --location-id string   Dialogflow CX Location ID of the Project (required)
  -p, --project-id string    Dialogflow CX Project ID (required)
  -r, --redacted string      Set the entity type as redacted. Possible values: true or false (optional)
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

