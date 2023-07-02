# Environments

## What is this?

In software, it is a common pattern (and a best practice) to have different environments where developers can deploy different versions of their software. Each environment has its own configuration.

Dialogflow CX follows the same concept in that you can create a version of your agent and then deploy it to a specific environment. This is also similar to a webhook in that you can deploy a specific webhook version and use that version within an environment.

With `cxcli`, you can easily interact with the environments of your Dialogflow CX agents.

All of the commands that you have available in `cxcli` to interact with your environments are located within the `cxcli environment` subcommand.

## Create

The `cxcli` has a command that allows you to create an environment. You can find a more detailed information [here](/environments/create).

## Update

The `cxcli` has a command that allows you to update an environment. You can find a more detailed information [here](/environments/update).


## Delete

The `cxcli` has a command that allows you to delete your environment. You can find a more detailed explanation [here](/environments/delete).

## CI/CD

In Dialogflow CX, when you are testing your agents, you can save those tests and associate them with a specific environment.

The `cxcli` tool has a command that allows you to run these CI/CD pipelines from your terminal or from your CI processes. You can find a more detailed explanation [here](/environments/cicd).

## Useful Links

If you want to see the full usage of the `cxcli environment` command, please refer to this [page](/cmd/cxcli_environment).

If you want to learn more about Dialogflow CX environments, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/version).
