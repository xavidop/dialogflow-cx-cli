# Export


## Usage

You can find the export command usage down the `cxcli agent export` command. You can read the documentation about this command [here](/cmd/cxcli_agent_export).


!!! info "Exported file format"
    You can export your agent in `blob` or `json` format. When you choose `json` format, the output file is going to be a `zip` file.


## Example

This a simple example of the `cxcli agent export` command:

```sh
cxcli agent export test-agent --project-id test-cx-346408 --location-id us-central1 --export-format blob
```

The command above will give you an output like this one:

```sh
$ cxcli agent export test-agent --project-id test-cx-346408 --location-id us-central1 --export-format json --output-file agent.zip
INFO Agent exported to file: agent.zip                    
```
!!! info "are you running this command in a CICD pipeline?"
    If this is the case, we recommend you to execute with the `--output-format` parameter set to `json`.

## Useful Links

If you want to learn more about Dialogflow CX exports, check the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/agent#export).