# Create


## Usage

You can find the create functionality within the `cxcli generator create` subcommand. You can read the documentation about this command [here](/cmd/cxcli_generator_create).

If you want to create a generator, you have to set the `--prompt` parameter. This parameter is the prompt that will be used to generate the responses. You can use placeholders in the prompt to pass parameters to the generator. The parameters can be extracted from the user's input, or from the agent's session parameters.

The placeholders have the following format: `$parameter-name`. For example:
```   
Give a warm welcome to the user with name $name
```

## Example

Here is a simple example of using the `cxcli generator create` command:

```sh
cxcli generator create my-generator --prompt "Give a warm welcome to the user with name $name" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 
```

The above command will give you output similar to the following:

```sh
$ cxcli generator create my-generator --prompt "Give a warm welcome to the user with name $name" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 
[INFO] Generator created with id: projects/test-cx-346408/locations/us-central1/agents/40278ea0-c0fc-4d9a-a4d4-caa68d86295f/generators/03b0a452-997c-40dc-afa2-6bef917d224f
```

## Useful Links

If you want to learn more about Dialogflow CX generator creation, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/generators).
