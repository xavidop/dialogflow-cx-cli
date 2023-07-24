# Agents

## What is this?

An agent in Dialogflow CX is the entity that handles all the conversations that we have defined on the Dialogflow CX console with the end users.

An agent is basically an assistant that will manage the state of each user's conversation when the end users are interacting with the agent through text or audio in multiple channels.

With `cxcli`, you can easily interact with your Dialogflow CX agents.

All of the commands that are available in `cxcli` to interact with your agents are located within the `cxcli agent` subcommand.

## Create

You can create an agent using this tool.

The `cxcli` has a command that allows you to create an agent. You can find a more detailed explanation [here](/agents/create)

## Update

You can update an agent using this cli.

The `cxcli` has a command that allows you to update an agent. You can find a more detailed explanation [here](/agents/udpate)

## Delete

The `cxcli` has a command that allows you to delete your agent. You can find a more detailed explanation [here](/agents/delete)

## Restore

You can restore an agent using a `blob` or a `json-package` file. Right now the Dialogflow CX API, used by the `cxcli`, works with the `blob` and `json` format.

The `cxcli` has a command that allows you to restore an agent. You can find a more detailed explanation [here](/agents/restore)


## Export

An agent can be exported as a `blob` or a `json-package` file. Right now the Dialogflow CX API, used by the `cxcli`, works with the `blob` and `json` format.

The `cxcli` has a command that allows you to export your agent. You can find a more detailed explanation [here](/agents/export)

## Useful Links

If you want to see the full usage of the `cxcli agent` command, please refer to this [page](/cmd/cxcli_agent).

If you want to learn more about Dialogflow CX agents, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/agent).
