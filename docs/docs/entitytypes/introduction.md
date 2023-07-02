# Entity types

## What is this?

One of the most important aspects of natural language understanding is entity types or entities. Entities contain the important factors within a message such as names, dates, products, organizations, places, or anything that we want to extract from the message. We call this concept “entities”. For example, let's take a look at the `order_intent` intent:

1. I want a pizza
2. I want 3 cokes
3. Give me two burgers

As you can see in the example above, we can extract two entity Types: `quantity` and `order_type`. We can represent the entities within the above utterances as:

1. I want `{quantity}` `{order_type}`
2. Give me `{quantity}` `{order_type}`

We can also think of entity types as variables.

With `cxcli`, you can easily interact with your Dialogflow CX entity types.

All of the commands that you have available in `cxcli` to interact with your entity types are located within the `cxcli entity-type` subcommand.

## Create

The `cxcli` has a command that allows you to create an entity type. You can find a more detailed [here](/entitytypes/create).

## Update

The `cxcli` has a command that allows you to update an entity type. You can find a more detailed [here](/entitytypes/update).


## Delete

The `cxcli` has a command that allows you to delete your entity type. You can find a more detailed explanation [here](/entitytypes/delete).


## Useful Links

If you want to see the full usage of the `cxcli entity-type` command, please refer to this [page](/cmd/cxcli_entity-type).

If you want to learn more about Dialogflow CX entity types, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/entity).
