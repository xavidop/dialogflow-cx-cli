# cxcli flow update

Updates a flow in an agent

```
cxcli flow update [flow-name] [flags]
```

## Options

```
  -a, --agent-name string                     Dialogflow CX Agent Name
  -d, --description string                    Optional. Description for this flow
  -h, --help                                  help for update
  -e, --locale string                         Optional. Locale of the flow. Default: agent locale
  -l, --location-id string                    Dialogflow CX Location ID of the Project
  -s, --nlu-classification-threshold string   Optional. NLU Classification Threshold. From 0.0 (completely uncertain) to 1.0 (completely certain)
  -t, --nlu-model-training-mode string        Optional. NLU Model training mode. Possible values: automatic or manual
  -m, --nlu-model-type string                 Optional. NLU Model Type. Possible values: advanced or standard
  -p, --project-id string                     Dialogflow CX Project ID
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file
  -o, --output-format string   Output Format (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command
  -v, --verbose                verbose error output (with stack trace)
```

## See also

* [cxcli flow](/cmd/cxcli_flow/)	 - Actions on flow commands

