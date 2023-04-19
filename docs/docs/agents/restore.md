# Restore


## Usage

You can find the restore functionality within the `cxcli agent restore` subcommand. You can read the documentation about this command [here](/cmd/cxcli_agent_restore).


!!! info "File format to be restored"
    Now you can restore your agents in `blob` or `json` format. When you choose `json` format, the input file has to be a `zip` file with a proper [Dialogflow CX structure](https://cloud.google.com/dialogflow/cx/docs/reference/json-export).

## Example

Here is a simple example of the `cxcli agent restore` command:

```sh
cxcli agent restore test-agent --project-id test-cx-346408 --location-id us-central1 --input agent.blob
```

The above command will give you output similar to the following:

```sh
$ cxcli agent restore test-agent --project-id test-cx-346408 --location-id us-central1 --input agent.zip
INFO Agent restored
```
!!! info "Are you running this command in a CI/CD pipeline?"
    If this is the case, we recommend that you set the `--output-format` parameter to `json`.

## Useful Links

If you want to learn more about Dialogflow CX restores, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/agent#export).
