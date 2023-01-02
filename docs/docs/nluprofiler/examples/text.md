# Text input with parameters example

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
      type: text
      text: I want 3 pizzas
    validate:
      intent: order_intent
      parameters:
        - parameter: number
          value: 3
        - parameter: order_type
          value: pizza

  - id: my_second_check
    input: 
      type: text
      text: I want 2 cokes
    validate:
      intent: order_intent
      parameters:
        - parameter: number
          value: 2
        - parameter: order_type
          value: coke
```