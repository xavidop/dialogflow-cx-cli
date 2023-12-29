# Update


## Usage

You can find the update functionality within the `cxcli webhook update` subcommand. You can read the documentation about this command [here](/cmd/cxcli_webhook_update).

### Flexible Webhook

If you want to update a flexible webhook, you have to set the `--flexible` parameter. When you set these parameters, you have to provide a `--request-body` and a `--parameters-mapping` parameter: 
    
1. The `--request-body` parameter is a JSON string that will be sent to the webhook.
2. The `--parameters-mapping` parameter is a comma-separated list of key-value pairs. The key is the name of the parameter that will be sent to the webhook, and the value is a JSON path that will be used to extract the value from the `--request-body` parameter. This parameter has the following format:
```
parameter@json-path,paramter2@json-path2
```
An example of a parameter mapping: `my-param@$.fully.qualified.path.to.field`

## Example

### Standard Webhook

Here is a simple example of the `cxcli webhook update` command:

```sh
cxcli webhook update my-webhook --url "https://my-webhook-updated.com" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 
```

The above command will give you output similar to the following:

```sh
$ cxcli webhook update my-webhook --url "https://my-webhook-updated.com" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 
INFO Webhook updated
```


### Flexible Webhook

Here is a simple example of the `cxcli webhook update` command:

```sh
cxcli webhook update my-webhook --url "https://my-webhook-updated.com" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 --flexible true --request-body "{\"hello\": false}" --parameters-mapping "my-param@$.fully.qualified.path.to.field"
```

The above command will give you output similar to the following:

```sh
$ cxcli webhook update my-webhook --url "https://my-webhook-updated.com" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 --flexible true --request-body "{\"hello\": false}" --parameters-mapping "my-param@$.fully.qualified.path.to.field"
INFO Webhook updated
```

## Useful Links

If you want to learn more about Dialogflow CX webhook update, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/webhook).