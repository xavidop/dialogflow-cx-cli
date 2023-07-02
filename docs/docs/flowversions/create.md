# Create


## Usage

You can find the create functionality within the `cxcli flow-version create` subcommand. You can read the documentation about this command [here](/cmd/cxcli_flow-version_create).

## Example

Here is a simple example of using the `cxcli flow-version create` command:

```sh
cxcli flow-version create "v1.0.0" --start-flow "Test Flow" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 --description "My flow version"
```

The above command will give you output similar to the following:

```sh
$ cxcli flow-version create "v1.0.0" --start-flow "Test Flow" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 --description "My flow version"
INFO Version created with id: projects/test-cx-346408/locations/us-central1/agents/40278ea0-c0fc-4d9a-a4d4-caa68d86295f/flows/b23e0247-2dd4-4d86-a23a-a289569480d4/versions/6
```

!!! info "Is it taking longer than expected?"
    That is okay, depending on the size of the flow it takes longer.

## Useful Links

If you want to learn more about Dialogflow CX flow version creation, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/version).
