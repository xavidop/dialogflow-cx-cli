# Create


## Usage

You can find the create command usage down the `cxcli intent create` command. You can read the documentation about this command [here](/cmd/cxcli_intent_create).


The `--training-phrases` parameter is a list of the training phrases for this intent, comma separated. For the entities used in this intent, add `@entity-type` to the word in the training phrase. This is the format: 

```
word@entity-type
```

Here you have an example: `hello, hi how are you today@sys.date, morning!`

## Example

This a simple example of the `cxcli intent create` command:

```sh
cxcli intent create test_intent --training-phrases "hello, hi how are you today@sys.date, morning"  --agent-name test-agent --project-id test-cx-346408 --location-id us-central1
```

The command above will give you an output like this one:

```sh
$ cxcli intent create test_intent --training-phrases "hello, hi how are you today@sys.date, morning"  --agent-name test-agent --project-id test-cx-346408 --location-id us-central1
INFO Intent created with id: projects/test-cx-346408/locations/us-central1/agents/40278ea0-c0fc-4d9a-a4d4-caa68d86295f/intents/a7870357-e942-43dd-99d2-4de8c81a3c09 
```

## Useful Links

If you want to learn more about Dialogflow CX Intent creation, check the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/intent).