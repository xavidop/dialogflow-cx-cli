# cxcli agent update

Update an agent

```
cxcli agent update [agent-name] [flags]
```

## Options

```
  -r, --avatar-uri string                   Avatar URI of the agent (optional)
  -d, --description string                  Description of the agent (optional)
  -b, --enable-interaction-logging string   Enable interaction logging for this agent. Possible values: true or false (optional)
  -s, --enable-speech-adaptation string     Enable speech adaptation for this agent. Possible values: true or false (optional)
  -n, --enable-spell-correction string      Enable spell correction for this agent. Possible values: true or false (optional)
  -a, --enable-stackdriver-logging string   Enable Google Stackdriver logging for this agent. Possible values: true or false (optional)
  -h, --help                                help for update
  -l, --location-id string                  Dialogflow CX Location ID of the Project (required)
  -p, --project-id string                   Dialogflow CX Project ID (required)
  -x, --supported-locales strings           Supported locales of the agent, comma separated. Example: fr,es,de (optional)
  -t, --timezone string                     Timezone of the agent from the time zone database https://www.iana.org/time-zones. Example: America/New_York, Europe/Madrid (optional)
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file path (optional)
  -o, --output-format string   Output Format. Options: text, json. Default: text (optional) (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command (optional)
  -v, --verbose                verbose error output (with stack trace) (optional)
```

## See also

* [cxcli agent](/cmd/cxcli_agent/)	 - Actions on agent commands

