# Interactive Dialog

<p align="center">
  <script async id="asciicast-594396" src="https://asciinema.org/a/594396.js"></script>
</p>

The `cxcli` tool has command that allows you to interactively play with your agent from your terminal!


## Usage

You can find the dialog functionality within the `cxcli dialog` subcommand. You can read the documentation about this command [here](/cmd/cxcli_dialog).

### Parameters

These are the relevant parameters that you can use to interact with your agent interactively from your terminal:

1. `locale`: this parameter accepts all of the locales that are available in Dialogflow CX.

## Example

Here is a simple example of the `cxcli dialog` command:

```sh
cxcli dialog --agent-name test-agent --project-id test-cx-346408 --location-id us-central1
```

The above command will give you output similar to the following:

```sh
$ cxcli dialog --agent-name test-agent --project-id test-cx-346408 --location-id us-central1
INFO Please press Ctrl+C whenever you want to stop the interaction. 
User> hi
Agent> Hi, this is Information Finder. Which product information do you need?
User> sorry, can you repeat
Agent> I didn't get that.
Agent> Hi, I am a chatbot that can help you find information about CorpX products: Intelligent Customer Insight, Order Fulfillment, Order Handling, Order Management Suite, Sales Engine, or Supercharge.
Agent> Do you want to continue looking up information for one of these products? Say yes or no.
User> no
Agent> Goodbye
```

!!! info "Important"
    Once the session is finished, you will get the terminal prompt again.