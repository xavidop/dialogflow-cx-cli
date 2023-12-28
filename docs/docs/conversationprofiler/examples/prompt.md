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
interactions:
  - id: test_1_1
    user: 
      type: prompt
      prompt: give me a one line hello
    agent:
      validate:
        - type: equals
          value: hello
```