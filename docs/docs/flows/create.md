# Create


## Usage

You can find the create functionality within the `cxcli flow create` subcommand. You can read the documentation about this command [here](/cmd/cxcli_flow_create).

## Example

Here is a simple example of using the `cxcli flow create` command:

```sh
cxcli flow create my-flow --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 --description "My test flow"
```

The above command will give you output similar to the following:

```sh
$ cxcli flow create my-flow --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 --description "My test flow"
INFO Flow created with id: projects/test-cx-346408/locations/us-central1/agents/40278ea0-c0fc-4d9a-a4d4-caa68d86295f/flows/49a38c72-2a63-4d71-a266-2d722eb8360e
```

## Useful Links

If you want to learn more about Dialogflow CX flow creation, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/flow).
