# Create


## Usage

You can find the create functionality within the `cxcli intent create` subcommand. You can read the documentation about this command [here](/cmd/cxcli_intent_create).


The `--training-phrases` parameter is a list of the training phrases for this intent, comma separated. For the entities used in this intent, add `@entity-type` to the word in the training phrase. This is the format:

```
word@entity-type
```

An example training phrase with a entity: `hello, hi how are you today@sys.date, morning!`

## Example

Here is a simple example of using the `cxcli intent create` command:

```sh
cxcli intent create test_intent --training-phrases "hello, hi how are you today@sys.date, morning"  --agent-name test-agent --project-id test-cx-346408 --location-id us-central1
```

The above command will give you output similar to the following:

```sh
$ cxcli intent create test_intent --training-phrases "hello, hi how are you today@sys.date, morning"  --agent-name test-agent --project-id test-cx-346408 --location-id us-central1
INFO Intent created with id: projects/test-cx-346408/locations/us-central1/agents/40278ea0-c0fc-4d9a-a4d4-caa68d86295f/intents/a7870357-e942-43dd-99d2-4de8c81a3c09
```

## Useful Links

If you want to learn more about Dialogflow CX intent creation, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/intent).
