# Intent

## What is this?

Before start talking about intents, it is important to understand what is NLU. The Natural Language Understanding (NLU) is a subset of Natural Language Processing (NLP). It helps a "machine" to be able to understand human language.
In Dialogflow CX this is an important part, since it will help predicting the user's intention and allowing us to act in a more "smart" way, and avoid the already typical: "I did not understand you, could you repeat it?". We call these intentions, proposals or user requests, which the machine must classify. These are the "Intents". Each intent has training phrases. For example the `welcome_intent` intent can have these 3 training phrases:

1. Hi
2. Hello
3. Whats up!

As you can see in the example above, our intention with the `welcome_intent` intent is to start a conversation when a user says any of these training phrases. An intent can have multiple entities.

Whenever you create, modify or deletes an intent it is important to re-train your Dialogflow CX flows. This will re-train your NLU. By doing this your bot will "understand you" including your latest changes.

With the `cxcli` you can interact easily with your Dialogflow CX intents.

All the commands that you have available in the `cxcli` to interact with your agents are located down the `cxcli intents` command.

## Create

The `cxcli` has a command that allows you to create an intents. You can find the whole explanation [here](/intents/create)


## Delete

The `cxcli` has a command that allows you to delete your intents. You can find the whole explanation [here](/intents/delete)


## Useful Links

If you want to check the full usage of the `cxcli intents` command, please refer to this [page](/cmd/cxcli_intent).

If you want to learn more about Dialogflow CX intents, check the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/intent).