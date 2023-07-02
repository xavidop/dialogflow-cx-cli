# Update


## Usage

You can find the update functionality within the `cxcli agent update` subcommand. You can read the documentation about this command [here](/cmd/cxcli_agent_update).

The argument to `--timezone` is a value from the [IANA time zone database](https://www.iana.org/time-zones). This is an example:
```
America/New_York
Europe/Madrid
```

## Example

Here is a simple example of the `cxcli agent update` command:

```sh
cxcli agent update test-agent --project-id test-cx-346408 --location-id us-central1 --locale en --timezone Europe/Madrid
```

The above command will give you output similar to the following:

```sh
$ cxcli agent update test-agent --project-id test-cx-346408 --location-id us-central1 --locale en --timezone Europe/Madrid
INFO Agent updated
```

## Useful Links

If you want to learn more about Dialogflow CX Agent update, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/agent).
