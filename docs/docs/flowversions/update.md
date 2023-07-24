# Update


## Usage

You can find the update functionality within the `cxcli flow-version update` subcommand. You can read the documentation about this command [here](/cmd/cxcli_flow-version_update).

## Example

Here is a simple example of the `cxcli flow update` command:

```sh
cxcli flow-version update "v1.0.0" --start-flow "Test Flow" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 --description "My updated flow version"
```

The above command will give you output similar to the following:

```sh
$ cxcli flow-version update "v1.0.0" --start-flow "Test Flow" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 --description "My updated flow version"
INFO Version updated
```

## Useful Links

If you want to learn more about Dialogflow CX flow version update, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/version).