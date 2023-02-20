# cxcli agent create

Creates an agent

```
cxcli agent create [agent-name] [flags]
```

## Options

```
  -r, --avatar-uri string                   Avatar URI of the agent
  -d, --description string                  Description of the agent
  -b, --enable-interaction-logging string   Enable interaction logging for this agent. Possible values: true or false
  -s, --enable-speech-adaptation string     Enable speech adaptation for this agent. Possible values: true or false
  -n, --enable-spell-correction string      Enable spell correction for this agent. Possible values: true or false
  -a, --enable-stackdriver-logging string   Enable Google Stackdriver logging for this agent. Possible values: true or false
  -h, --help                                help for create
  -e, --locale string                       Default locale of the agent
  -l, --location-id string                  Dialogflow CX Location ID of the Project
  -p, --project-id string                   Dialogflow CX Project ID
  -x, --supported-locales strings           Supported locales of the agent, comma separated. Example: fr,es,de
  -t, --timezone string                     Timezone of the agent from the time zone database https://www.iana.org/time-zones. Example: America/New_York, Europe/Madrid
```

## Options inherited from parent commands

```
  -c, --credentials string     Google Cloud credentials JSON file
  -o, --output-format string   Output Format (default "text")
  -u, --skip-update-check      Skip the check for updates check run before every command
  -v, --verbose                verbose error output (with stack trace)
```

## See also

* [cxcli agent](/cmd/cxcli_agent/)	 - Actions on agent commands

