# Authentication

`cxcli` uses some Google cloud APIs. By default the tool uses the default configuration that uses the `gcloud` cli. If you want to use another authentication key you can provide a `json` file with the global `--credentials` parameter.

The `cxcli` source code is open source, you can check it out [here](https://github.com/xavidop/dialogflow-cx-cli) to learn more about the actions the tool performs.

Below you can find the roles and the APIs needed to use the tool.

## Roles needed

### Dialogflow CX

**Dialogflow API Admin**: Provides full access to create, update, query, detect intent, and delete the agent from the console or API. Click [here](https://cloud.google.com/dialogflow/cx/docs/concept/access-control) for more information.

We are using the Admin role because `cxcli` performs the [List agent](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3beta1/projects.locations.agents/list) action.

This role allows you to execute Speech-to-text and Text-to-speech actions

## APIs enabled needed

These APIs should be enabled on your Google Cloud project if you want to use these `cxcli` capabilities:

### Dialogflow CX

You will need to enable the `Dialogflow API` on your project. More information [here](https://cloud.google.com/dialogflow/cx/docs)

### Speech-to-text

You will need to enable the `Cloud Speech-to-Text API` on your project. More information [here](https://cloud.google.com/speech-to-text/docs/transcribe-api)

### Text-to-speech

You will need to enable the `Cloud Text-to-Speech API` on your project. More information [here](https://cloud.google.com/text-to-speech/docs/apis)

