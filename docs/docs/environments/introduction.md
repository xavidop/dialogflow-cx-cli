# Environments

## What is this?

In software, it is a common pattern (and a best practice) to have different environments where developers can deploy different versions of their software. Each environment has its own configurations.

In Dialogflow CX we have the same concept, you can create a version of your agent and then, deploy it to an environment. Same with the webhook, you can deploy a webhook version and use that version in an environment.

With the `cxcli` you can interact easily with the environments of your Dialogflow CX agents.

All the commands that you have available in the `cxcli` to interact with your environments are located down the `cxcli environment` command.

## CICD

In Dialogflow CX, while you are testing your agents, you can save those tests and associate them to a specific environment.

The `cxcli` has a command that allows you to run these cicd pipelines from your terminal or from your CI processes. You can find the whole explanation [here](/environments/cicd)

## Useful Links

If you want to check the full usage of the `cxcli environment` command, please refer to this [page](/cmd/cxcli_environment).

If you want to learn more about Dialogflow CX environments, check the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/version).