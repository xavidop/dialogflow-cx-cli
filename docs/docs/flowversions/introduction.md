# Flow Versions

## What is this?

It is possible to generate various iterations of your agent flows and deploy them to distinct serving environments.

When making modifications to a flow, you are essentially modifying a draft version. At any given time, you have the option to save a draft flow as a flow version. A flow version represents an unchangeable snapshot of your flow data along with all the related agent data, such as intents, entities, webhooks, pages, route groups, and more.

Furthermore, You can deploy this flow versions to different environments, such as development, test, and production. You can also deploy the same flow version to multiple environments.

With `cxcli`, you can easily interact with the flow versions of your Dialogflow CX agents.

All of the commands that you have available in `cxcli` to interact with your flows are located within the `cxcli flow-version` subcommand.

## Create

The `cxcli` has a command that allows you to create a flow versions. You can find a more detailed information [here](/flowversions/create).

## Update

The `cxcli` has a command that allows you to update a flow versions. You can find a more detailed information [here](/flowversions/update).


## Delete

The `cxcli` has a command that allows you to delete your flow versions. You can find a more detailed explanation [here](/flowversions/delete).


## Useful Links

If you want to see the full usage of the `cxcli flow-version` command, please refer to this [page](/cmd/cxcli_flow-version).

If you want to learn more about Dialogflow CX flows, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/version).
