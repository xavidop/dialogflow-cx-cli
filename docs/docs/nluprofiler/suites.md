# Suites

## Reference

A suite is a yaml file with the following structure:

```yaml
# suite.yaml

# Name of the suite.
name: Example Suite
# Brief description of the suite.
description: Suite used as an example
# Project ID on Google Cloud where is located your Dialogflow CX agent.
projectId: test-cx-346408
# Location where your Dialogflow CX agent is running. 
# More info here: https://cloud.google.com/dialogflow/cx/docs/concept/region
locationId: us-central1
# Agent name of your Dialogflow CX agent.
# Notice: it is the agent name, not the agent ID.
agentName: test-agent
# You can have multiple tests defined in separated files
tests:
  # ID of the test.
  - id: test_id
    # File where the test specification is located
    file: ./test.yaml
```