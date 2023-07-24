# Update


## Usage

You can find the update functionality within the `cxcli intent update` subcommand. You can read the documentation about this command [here](/cmd/cxcli_intent_update).


The `--training-phrases` parameter is a list of the training phrases for this intent, comma separated. For the entities used in this intent, add `@entity-type` to the word in the training phrase. This is the format:

```
word@entity-type
```

An example training phrase with a entity: `hello, hi how are you today@sys.date, morning!`

## Example

Here is a simple example of using the `cxcli intent update` command:

```sh
cxcli intent update test_intent --training-phrases "hello, hi how are you today@sys.date, morning"  --agent-name test-agent --project-id test-cx-346408 --location-id us-central1
```

The above command will give you output similar to the following:

```sh
$ cxcli intent update test_intent --training-phrases "hello, hi how are you today@sys.date, morning"  --agent-name test-agent --project-id test-cx-346408 --location-id us-central1
INFO Intent updated
```

## Useful Links

If you want to learn more about Dialogflow CX intent update, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/intent).
