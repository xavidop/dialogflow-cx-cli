# Update


## Usage

You can find the update functionality within the `cxcli generator update` subcommand. You can read the documentation about this command [here](/cmd/cxcli_generator_update).

If you want to update a generator, you have to set the `--prompt` parameter. This parameter is the prompt that will be used to generate the responses. You can use placeholders in the prompt to pass parameters to the generator. The parameters can be extracted from the user's input, or from the agent's session parameters. 

The placeholders have the following format: `$parameter-name`. For example:
```   
Give a warm welcome to the user with name $name
```

    

## Example

Here is a simple example of the `cxcli generator update` command:

```sh
cxcli generator update my-generator --prompt "Give a warm farewell to the user with name $name" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 
```

The above command will give you output similar to the following:

```sh
$ cxcli generator update my-generator --prompt "Give a warm farewell to the user with name $name" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 
INFO Generator updated
```

## Useful Links

If you want to learn more about Dialogflow CX generator update, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/generators).