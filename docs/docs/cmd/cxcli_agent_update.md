# cxcli agent update

Update an agent

```
cxcli agent update [agent-name] [flags]
```

## Options

```
  -r, --avatar-uri string            Avatar URI of the agent
  -d, --description string           Description of the agent
  -b, --enable-interaction-logging   Enable interaction logging for this agent
  -s, --enable-speech-adaptation     Enable speech adaptation for this agent
  -n, --enable-spell-correction      Enable spell correction for this agent
  -a, --enable-stackdriver-logging   Enable Google Stackdriver logging for this agent
  -h, --help                         help for update
  -l, --location-id string           Dialogflow CX Location ID of the Project
  -p, --project-id string            Dialogflow CX Project ID
  -x, --supported-locales strings    Supported locales of the agent, comma separated. Example: fr,es,de
  -t, --timezone string              Timezone of the agent from the time zone database https://www.iana.org/time-zones. Example: America/New_York, Europe/Madrid
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

