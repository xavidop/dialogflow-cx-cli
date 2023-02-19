# cxcli environment create

create an environment

```
cxcli environment create [name] [flags]
```

## Options

```
  -a, --agent-name string       Dialogflow CX Agent Name
  -d, --description string      Optional. Description for this environment
  -s, --flow-versions strings   List of Flow and its version to be added to this environment, comma separated. Format: flowName1@version1,flowName2|version2. Example: Default Start Flow@v1.0.0|Buy Flow@v1.0.1
  -h, --help                    help for create
  -l, --location-id string      Dialogflow CX Location ID of the Project
  -p, --project-id string       Dialogflow CX Project ID
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file
  -o, --output-format string   Output Format (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command
  -v, --verbose                verbose error output (with stack trace)
```

## See also

* [cxcli environment](/cmd/cxcli_environment/)	 - Actions on environment

