# Create


## Usage

You can find the create functionality within the `cxcli entity-type create` subcommand. You can read the documentation about this command [here](/cmd/cxcli_entity-type_create).


The argument to `--entities` is a list of the entities with their synonyms, comma separated. This parameter has the following format:
```
entity1@synonym1|synonym2,entity2@synonym1|synonym2
```

An example entity type with synonyms: `pikachu@25|pika,charmander@3`

## Example

Here is a simple example of the `cxcli entity-type create` command:

```sh
cxcli entity-type create pokemon --entities "pikachu@25|pika,charmander@3" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1
```

The above command will give you output similar to the following:

```sh
$ cxcli entity-type create pokemon --entities "pikachu@25|pika,charmander@3" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1
INFO Entity Type created with id: projects/test-cx-346408/locations/us-central1/agents/40278ea0-c0fc-4d9a-a4d4-caa68d86295f/entityTypes/457a451d-f5ce-47da-b8dc-16b17d874a5d
```

## Useful Links

If you want to learn more about Dialogflow CX entity types creation, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/entity).
