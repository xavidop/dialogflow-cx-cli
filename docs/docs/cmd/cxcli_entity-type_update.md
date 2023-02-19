# cxcli entity-type update

Updates an entity type in an agent

```
cxcli entity-type update [entity-type-name] [flags]
```

## Options

```
  -a, --agent-name string    Dialogflow CX Agent Name
  -n, --entities strings     List of the entities for this entity type, comma separated. Format: entity1@synonym1|synonym2,entity2@synonym1|synonym2. Example: pikachu@25|pika,charmander@3
  -h, --help                 help for update
  -e, --locale string        Optional. Locale of the intent. Default: agent locale
  -l, --location-id string   Dialogflow CX Location ID of the Project
  -p, --project-id string    Dialogflow CX Project ID
  -r, --redacted             Optional. Set the entity type as redacted
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file
  -o, --output-format string   Output Format (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command
  -v, --verbose                verbose error output (with stack trace)
```

## See also

* [cxcli entity-type](/cmd/cxcli_entity-type/)	 - Actions on entity type commands

