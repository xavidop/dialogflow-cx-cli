# Environments CICD

## Usage

You can find the cicd command usage down the `cxcli environment execute-cicd` command. You can read the documentation about this command [here](/cmd/cxcli_environment_execute-cicd).

## Example

This a simple example of the `cxcli environment execute-cicd` command:

```sh
cxcli environment execute-cicd cicd-env --project-id test-cx-346408 --location-id us-central1 --agent-name test-agent
```

The command above will give you an output like this one:

```sh
$ cxcli environment execute-cicd cicd-env --project-id test-cx-346408 --location-id us-central1 --agent-name test-agent
INFO Executing cicd for environment cicd-env      
INFO PASSED                     
```

!!! info "are you running this command in a CICD pipeline?"
    If this is the case, we recommend you to execute with the `--output-format` parameter set to `json`.


## Useful Links

If you want to learn more about Dialogflow CX cicd executions, check the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/continuous-tests).