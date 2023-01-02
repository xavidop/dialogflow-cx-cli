# Tests

## Reference

A test is a yaml file with the following structure:

```yaml
# test.yaml

# Name of the test.
name: Example test
# Brief description of the test.
description: These are some tests
# Locale of the interaction model that is gonna be tested. 
# You can find the locales here: https://cloud.google.com/dialogflow/cx/docs/reference/language
localeId: en
# A check is the test itself: given an input, you will validate the intents and the parameters/entities detected by Dialogflow CX
# You can have multiple checks defined
checks:
  # The ID of the check
  - id: test
    input:
      # the input type
      # it could be text or audio
      type: text
      # The input itself in text format. For type: audio, you have to specify the audio tag.
      text: I want 3 pizzas
    validate:
      # Intent that is supposed to be detected
      intent: order_intent
      # You can have multiple parameters/intents
      # Notice: this could be empty if your intent does not have any entities/parameters.
      parameters:
        # Entity name that is supposed to be detected
        - parameter: number
          # Value that is supposed to be detected
          value: 3
```

## Audio input

It is important to know that the input has to have this format:

1. A Sample Rate Hertz of 16000Hz
2. The audio encoding has to be be Linear16. Linear16 is a 16-bit linear pulse-code modulation (PCM) encoding.

If you don't have a file with this format, you can create it by yourself using the `cxcli tts` command! All the information is located [here](/tts)

## JSON Schema

`cxcli` also has a [jsonschema](http://json-schema.org/draft/2020-12/json-schema-validation.html) file, which you can use to have better
editor support:

```sh
https://cxcli.xavidop.me/static/test.json
```

You can also specify it in your `yml` config files by adding a
comment like the following:
```yaml
# yaml-language-server: $schema=https://cxcli.xavidop.me/static/test.json
```