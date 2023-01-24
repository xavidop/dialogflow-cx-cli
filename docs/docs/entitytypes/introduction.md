# Entity types

## What is this?

One of the most important parts of NLU is the Entity types or entities. These are the key information in a text, such as names, dates, products, organizations, places, or anything we want to extract from the text. We call this concept “Entities”. For example, let's take a look at the `order_intent` intent:

1. I want a pizza
2. I want 3 cokes
3. Give me two burgers

As you can see in the example above, we can extract 2 Entity Types: `quantity` and `order_type`. We can extrapolate the utterances above like this:

1. I want `{quantity}` `{order_type}`
2. Give me `{quantity}` `{order_type}`

We can think about Entity Types as variables as well.

With the `cxcli` you can interact easily with your Dialogflow CX entity types.

All the commands that you have available in the `cxcli` to interact with your agents are located down the `cxcli entity-type` command.

## Create

The `cxcli` has a command that allows you to create an entity type. You can find the whole explanation [here](/entitytypes/create)


## Delete

The `cxcli` has a command that allows you to delete your entity type. You can find the whole explanation [here](/entitytypes/delete)


## Useful Links

If you want to check the full usage of the `cxcli entity-type` command, please refer to this [page](/cmd/cxcli_entity-type).

If you want to learn more about Dialogflow CX entity types, check the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/entity).