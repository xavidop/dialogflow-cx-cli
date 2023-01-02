# Agents

## What is this?

An agent in Dialogflow CX is the entity that handles all the conversations that we have defined on the Dialogflow CX console with the end users.

An agent is basically an assistant that will manage the state of each user's conversation when the end users are interacting with the agent through text or audio in multiple channels.

With the `cxcli` you can interact easily with your Dialogflow CX agents.

All the commands that you have available in the `cxcli` to interact with your agents are located down the `cxcli agent` command.

## Restore

You can restore an agent using a `blob` file. Right now the Dialogflow CX API, used by the `cxcli`, only works with the `blob` format.

The `cxcli` has a command that allows you to restore an agent. You can find the whole explanation [here](/agents/restore)


## Export

An agent can be exported as a `blob` file. Right now the Dialogflow CX API, used by the `cxcli`, only works with the `blob` format.

The `cxcli` has a command that allows you to export your agent. You can find the whole explanation [here](/agents/export)


## Useful Links

If you want to check the full usage of the `cxcli agent` command, please refer to this [page](/cmd/cxcli_agent).

If you want to learn more about Dialogflow CX agents, check the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/agent).