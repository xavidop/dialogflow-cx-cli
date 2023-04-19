# Export


## Usage

You can find the export functionality within the `cxcli agent export` subcommand. You can read the documentation about this command [here](/cmd/cxcli_agent_export).


!!! info "Exported file format"
    You can export your agent in `blob` or `json` format. When you choose `json` format, the output file is going to be a `zip` file.


## Example

Here is a simple example of the `cxcli agent export` command:

```sh
cxcli agent export test-agent --project-id test-cx-346408 --location-id us-central1 --export-format blob
```

The above command will give you output similar to the following:

```sh
$ cxcli agent export test-agent --project-id test-cx-346408 --location-id us-central1 --export-format json --output-file agent.zip
INFO Agent exported to file: agent.zip
```
!!! info "are you running this command in a CI/CD pipeline?"
    If this is the case, we recommend that you set the `--output-format` parameter to `json`.

## Useful Links

If you want to learn more about Dialogflow CX exports, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/agent#export).
