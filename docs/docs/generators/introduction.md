# Generators

## What is this?

Generators use Google's latest generative large language models (LLMs), and prompts that you provide, to generate agent behavior and responses at runtime. The available models are provided by Vertex AI.

You can pass parameters to the generators using placeholders in the prompts. The parameters can be extracted from the user's input, or from the agent's session parameters.

This is an example of a generator prompt:

```
Give a warm welcome to the user
```

And this is an example of a generator prompt with a placeholder:

```
Give a warm welcome to the user with name $name
```

With `cxcli`, you can easily interact with the generators of your Dialogflow CX agents.

All of the commands that you have available in `cxcli` to interact with your generators are located within the `cxcli generator` subcommand.

## Create

The `cxcli` has a command that allows you to create a generator. You can find a more detailed information [here](/generators/create).

## Update

The `cxcli` has a command that allows you to update a generator. You can find a more detailed information [here](/generators/update).

## Delete

The `cxcli` has a command that allows you to delete a generator. You can find a more detailed explanation [here](/generators/delete).

## Useful Links

If you want to see the full usage of the `cxcli generator` command, please refer to this [page](/cmd/cxcli_generator).

If you want to learn more about Dialogflow CX generators, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/generators).
