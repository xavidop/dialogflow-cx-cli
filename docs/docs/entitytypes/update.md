# Update


## Usage

You can find the update functionality within the `cxcli entity-type update` subcommand. You can read the documentation about this command [here](/cmd/cxcli_entity-type_update).


The argument to `--entities` is a list of the entities with their synonyms, comma separated. This parameter has the following format:
```
entity1@synonym1|synonym2,entity2@synonym1|synonym2
```

An example entity type with synonyms: `pikachu@25|pika,charmander@3`

## Example

Here is a simple example of the `cxcli entity-type update` command:

```sh
cxcli entity-type update pokemon --entities "pikachu@25|pika,charmander@3" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1
```

The above command will give you output similar to the following:

```sh
$ cxcli entity-type update pokemon --entities "pikachu@25|pika,charmander@3" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1
INFO Entity Type updated
```

## Useful Links

If you want to learn more about Dialogflow CX entity types update, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/entity).
