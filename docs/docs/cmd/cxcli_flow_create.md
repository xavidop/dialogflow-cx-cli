# cxcli flow create

Creates a flow in an agent

```
cxcli flow create [flow-name] [flags]
```

## Options

```
  -a, --agent-name string                     Dialogflow CX Agent Name (required)
  -d, --description string                    Description for this flow (optional)
  -h, --help                                  help for create
  -e, --locale string                         Locale of the flow. Default: agent locale (optional)
  -l, --location-id string                    Dialogflow CX Location ID of the Project (required)
  -s, --nlu-classification-threshold string   NLU Classification Threshold. From 0.0 (completely uncertain) to 1.0 (completely certain). Default: 0.3 (optional) (default "0.3")
  -t, --nlu-model-training-mode string        NLU Model training mode. Possible values: automatic or manual. Default: manual (optional) (default "manual")
  -m, --nlu-model-type string                 NLU Model Type. Possible values: advanced or standard. Default: standard (optional) (default "standard")
  -p, --project-id string                     Dialogflow CX Project ID (required)
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file path (optional)
  -o, --output-format string   Output Format. Options: text, json. Default: text (optional) (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command (optional)
  -v, --verbose                verbose error output (with stack trace) (optional)
```

## See also

* [cxcli flow](/cmd/cxcli_flow/)	 - Actions on flow commands

