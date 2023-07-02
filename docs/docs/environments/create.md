# Create


## Usage

You can find the create functionality within the `cxcli environment create` subcommand. You can read the documentation about this command [here](/cmd/cxcli_environment_create).

The `--flow-versions` parameter is a list of flows and their versions to include in the environment, comma separated. For each `flow` used in this environment, it is required to add the `version` with this format: `@flow-version`. This is the full format:

```
flow@flow-version
```

An example of flows with their versions:

```bash 
Default Start Flow@production_v2,Test Flow@v1.0.0
```

## Example

Here is a simple example of using the `cxcli environment create` command:

```sh
cxcli environment create my-env --flow-versions "Default Start Flow@production_v2,Test Flow@v1.0.0"  --agent-name test-agent --project-id test-cx-346408 --location-id us-central1
```

The above command will give you output similar to the following:

```sh
$ cxcli environment create my-env --flow-versions "Default Start Flow@production_v2,Test Flow@v1.0.0"  --agent-name test-agent --project-id test-cx-346408 --location-id us-central1
INFO Environment created with id: projects/test-cx-346408/locations/us-central1/agents/40278ea0-c0fc-4d9a-a4d4-caa68d86295f/environments/9f9cf42a-9033-4d73-9407-613d041c9403
```

## Useful Links

If you want to learn more about Dialogflow CX environment creation, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/version#manage-environments).
