# Create


## Usage

You can find the create functionality within the `cxcli agent create` subcommand. You can read the documentation about this command [here](/cmd/cxcli_agent_create).

The argument to `--timezone` is a value from the [IANA time zone database](https://www.iana.org/time-zones). This is an example:
```
America/New_York
Europe/Madrid
```

## Example

Here is a simple example of the `cxcli agent create` command:

```sh
cxcli agent create test-agent --project-id test-cx-346408 --location-id us-central1 --locale en --timezone Europe/Madrid
```

The above command will give you output similar to the following:

```sh
$ cxcli agent create test-agent --project-id test-cx-346408 --location-id us-central1 --locale en --timezone Europe/Madrid
INFO Agent created with id: projects/test-cx-346408/locations/us-central1/agents/e2ae2503-f38c-46a1-a3bf-22e57617baf0
```

## Useful Links

If you want to learn more about Dialogflow CX Agent creation, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/agent).
