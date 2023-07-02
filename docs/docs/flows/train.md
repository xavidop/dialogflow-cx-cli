# Train

## Usage

You can find the train functionality within the `cxcli flow train` subcommand. You can read the documentation about this command [here](/cmd/cxcli_flow_train).

## Example

Here is a simple example of the `cxcli flow train` command:

```sh
cxcli flow train my-flow --project-id test-cx-346408 --location-id us-central1 --agent-name test-agent
```

The above command will give you output similar to the following:

```sh
$ cxcli flow train my-flow --project-id test-cx-346408 --location-id us-central1 --agent-name test-agent
INFO Flow trained
```

!!! info "Is it taking longer than expected?"
    That is okay, depending on the amount of data to train it takes longer.

## Useful Links

If you want to learn more about Dialogflow CX flow train execution, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/flow).
