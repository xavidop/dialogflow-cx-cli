# Intent

## What is this?

Before we discuss intents, it is important to first understand what is natural language understanding (NLU). Natural language understanding is a subset of natural language processing (NLP). It helps a "machine" to be able to understand human language.

In Dialogflow CX, this is an important concept, since it will help predict the user's intention and allow us to act in a "smarter" way, and avoid the all-too-common default fallback intent: "I did not understand you, could you repeat?".

We refer to these intentions, proposals, or user requests that machine must classify as "intents". Each intent has training phrases. For example, the default `welcome_intent` intent can contain these three training phrases:

1. Hi
2. Hello
3. Whats up!

As you can see in the example above, the purpose of the `welcome_intent` intent is to start a conversation when a user says any of these training phrases. An intent can have multiple entities.

Whenever you create, modify, or delete an intent, it is important to re-train your Dialogflow CX flows. This will re-train your NLU model. By doing this, your bot will "understand you", including your latest changes.

With `cxcli`, you can easily interact with your Dialogflow CX intents.

All of the commands that you have available in `cxcli` to interact with your intents are located within the `cxcli intents` subcommand.

## Create

The `cxcli` has a command that allows you to create an intent. You can find a more detailed explanation [here](/intents/create)


## Delete

The `cxcli` has a command that allows you to delete your intents. You can find a more detailed explanation [here](/intents/delete)


## Useful Links

If you want to see the full usage of the `cxcli intents` command, please refer to this [page](/cmd/cxcli_intent).

If you want to learn more about Dialogflow CX intents, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/intent).
