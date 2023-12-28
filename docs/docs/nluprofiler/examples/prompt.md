# Prompt example

## Suite file

```yaml
# suite.yaml

name: Example Suite
description: Suite used as an example
projectId: test-cx-346408
locationId: us-central1
agentName: test-agent
tests:
  - id: test_id
    file: ./test.yaml
```

## Test file

```yaml
# test.yaml

name: Example test
description: These are some tests
localeId: en
checks:
  - id: my_first_check
    input: 
      type: prompt
      prompt: give me a one line hello without exclamation mark using only 2 characters and lowercase
    validate:
      intent: hi_intent

  - id: my_second_check
    input: 
      type: text
      text: hello
    validate:
      intent: hi_intent
```