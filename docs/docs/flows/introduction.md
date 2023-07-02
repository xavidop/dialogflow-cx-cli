# Flows

## What is this?

In intricate conversations, various topics come into play. Take, for example, a flight book service, where a dialogue may cover areas like seat booking a meal reservation, obtaining customer information, and confirming the booking. Each of these topics requires multiple interactions between the agent and the end-user to gather the necessary information.

Flows are employed to delineate these topics and their respective conversational paths. Each agent is equipped with a default start flow, which serves as a foundation for simple agents.

With `cxcli`, you can easily interact with the flows of your Dialogflow CX agents.

All of the commands that you have available in `cxcli` to interact with your flows are located within the `cxcli flow` subcommand.

## Create

The `cxcli` has a command that allows you to create a flow. You can find a more detailed information [here](/flows/create).

## Update

The `cxcli` has a command that allows you to update a flow. You can find a more detailed information [here](/flows/update).


## Delete

The `cxcli` has a command that allows you to delete your flow. You can find a more detailed explanation [here](/flows/delete).

## Train

In Dialogflow CX, when you are adding training phrases to your intents and synonyms to your entity types, you will need to train your agent in order to have the latest changes ready for your end-users.

The `cxcli` tool has a command that allows you train your flow from your terminal or from your CI processes. You can find a more detailed explanation [here](/flows/train).

## Useful Links

If you want to see the full usage of the `cxcli flow` command, please refer to this [page](/cmd/cxcli_flow).

If you want to learn more about Dialogflow CX flows, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/flow).
