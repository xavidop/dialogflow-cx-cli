# System entity detection example

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
  - id: my_check
    input: 
      type: text
      text: please, give me 3
    validate:
      intent: finish_order_intent
      parameters:
        - parameter: number
          value: 3
```