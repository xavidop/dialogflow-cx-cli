# Update


## Usage

You can find the update functionality within the `cxcli environment update` subcommand. You can read the documentation about this command [here](/cmd/cxcli_environment_update).

The `--flow-versions` parameter is a list of flows and their versions to include in the environment, comma separated. For each `flow` used in this environment, it is required to add the `version` with this format: `@flow-version`. This is the full format:

```
flow@flow-version
```

An example of flows with their versions:

```bash 
Default Start Flow@production_v2,Test Flow@v1.0.0
```

## Example

Here is a simple example of the `cxcli entity-type update` command:

```sh
cxcli environment update my-env --flow-versions "Default Start Flow@production_v3,Test Flow@v1.1.0"  --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 --description "This is a test environment"
```

The above command will give you output similar to the following:

```sh
$ cxcli environment update my-env --flow-versions "Default Start Flow@production_v3,Test Flow@v1.1.0"  --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 --description "This is a test environment"
INFO Environment updated
```

## Useful Links

If you want to learn more about Dialogflow CX environment update, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/version#manage-environments).