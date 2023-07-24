# Webhooks

## What is this?

Webhooks serve as platforms for hosting your business logic or invoking other services. Within a session, webhooks enable you to utilize the data extracted through Dialogflow's natural language processing to generate dynamic responses, verify gathered data, or initiate actions on the backend.

There are two types of webhooks: standard webhooks and flexible webhooks. In the case of a standard webhook, the request and response fields are determined by Dialogflow. On the other hand, a flexible webhook allows you to specify the request and response fields according to your requirements.

With `cxcli`, you can easily interact with the webhooks of your Dialogflow CX agents.

All of the commands that you have available in `cxcli` to interact with your webhooks are located within the `cxcli webhook` subcommand.

You can create, update or delete standard and flexible webhooks with `cxcli` for a specific environment by setting the `--environment` parameter. If you do not specify an environment, the `cxcli` will create, update or delete this webhook for all environments.

## Create

The `cxcli` has a command that allows you to create a standard or flexible webhook. You can find a more detailed information [here](/webhooks/create).

## Update

The `cxcli` has a command that allows you to update a standard or flexible webhook. You can find a more detailed information [here](/webhooks/update).

## Delete

The `cxcli` has a command that allows you to delete your standard or flexible webhook. You can find a more detailed explanation [here](/webhooks/delete).

## Useful Links

If you want to see the full usage of the `cxcli webhook` command, please refer to this [page](/cmd/cxcli_webhook).

If you want to learn more about Dialogflow CX webhooks, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/webhook).
