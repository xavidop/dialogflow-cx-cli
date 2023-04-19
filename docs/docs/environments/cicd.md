# Environments CI/CD

## Usage

You can find the CI/CD functionality within the `cxcli environment execute-cicd` subcommand. You can read the documentation about this command [here](/cmd/cxcli_environment_execute-cicd).

## Example

Here is a simple example of the `cxcli environment execute-cicd` command:

```sh
cxcli environment execute-cicd cicd-env --project-id test-cx-346408 --location-id us-central1 --agent-name test-agent
```

The above command will give you output similar to the following:

```sh
$ cxcli environment execute-cicd cicd-env --project-id test-cx-346408 --location-id us-central1 --agent-name test-agent
INFO Executing cicd for environment cicd-env
INFO PASSED
```

!!! info "Are you running this command in a CI/CD pipeline?"
    If this is the case, we recommend that you set the `--output-format` parameter to `json`.


## Useful Links

If you want to learn more about Dialogflow CX CI/CD execution, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/continuous-tests).
