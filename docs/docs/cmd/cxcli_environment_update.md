# cxcli environment update

update an environment

```
cxcli environment update [name] [flags]
```

## Options

```
  -a, --agent-name string       Dialogflow CX Agent Name (required)
  -d, --description string      Description for this environment (optional)
  -s, --flow-versions strings   List of Flow and its version to be added to this environment, comma separated. Format: flowName1@version1,flowName2|version2. Example: Default Start Flow@v1.0.0|Buy Flow@v1.0.1 (required)
  -h, --help                    help for update
  -l, --location-id string      Dialogflow CX Location ID of the Project (required)
  -p, --project-id string       Dialogflow CX Project ID (required)
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file path (optional)
  -o, --output-format string   Output Format. Options: text, json. Default: text (optional) (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command (optional)
  -v, --verbose                verbose error output (with stack trace) (optional)
```

## See also

* [cxcli environment](/cmd/cxcli_environment/)	 - Actions on environment

