# Create


## Usage

You can find the create functionality within the `cxcli webhook create` subcommand. You can read the documentation about this command [here](/cmd/cxcli_webhook_create).

### Flexible Webhook

If you want to create a flexible webhook, you have to set the `--flexible` parameter. When you set these parameters, you have to provide a `--request-body` and a `--parameters-mapping` parameter: 
    
1. The `--request-body` parameter is a JSON string that will be sent to the webhook.
2. The `--parameters-mapping` parameter is a comma-separated list of key-value pairs. The key is the name of the parameter that will be sent to the webhook, and the value is a JSON path that will be used to extract the value from the `--request-body` parameter. This parameter has the following format:
```
parameter@json-path,paramter2@json-path2
```
An example of a parameter mapping: `my-param@$.fully.qualified.path.to.field`

## Example

### Standard Webhook

Here is a simple example of using the `cxcli webhook create` command:

```sh
cxcli webhook create my-webhook --url "https://my-webhook.com" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 
```

The above command will give you output similar to the following:

```sh
$ cxcli webhook create my-webhook --url "https://my-webhook.com" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 
INFO Webhook created with id: projects/test-cx-346408/locations/us-central1/agents/40278ea0-c0fc-4d9a-a4d4-caa68d86295f/webhooks/55f56aeb-be30-40a2-8bd6-cbbd6b9cc041
```

### Flexible Webhook

```sh
cxcli webhook create my-webhook --url "https://my-webhook.com" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 --flexible true --request-body "{\"hello\": true}" --parameters-mapping "my-param@$.fully.qualified.path.to.field, my-param2@$.fully.qualified.path.to.field2"
```

The above command will give you output similar to the following:

```sh
$ cxcli webhook create my-webhook --url "https://my-webhook.com" --agent-name test-agent --project-id test-cx-346408 --location-id us-central1 --flexible true --request-body "{\"hello\": true}" --parameters-mapping "my-param@$.fully.qualified.path.to.field, my-param2@$.fully.qualified.path.to.field2"
INFO Webhook created with id: projects/test-cx-346408/locations/us-central1/agents/40278ea0-c0fc-4d9a-a4d4-caa68d86295f/webhooks/13df6f13-6848-4fab-8cda-752b4f9819fa 
```

## Useful Links

If you want to learn more about Dialogflow CX webhook creation, refer to the [official documentation](https://cloud.google.com/dialogflow/cx/docs/concept/webhook).
