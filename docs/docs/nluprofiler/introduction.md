# NLU profiler

## What is this?

<p align="center">
  <img alt="Flow" src="/images/flow.png" style="height:512px;width:512px" />
</p>

Use the NLU Profiler to test user utterances and improve your agent's interaction model.

With the NLU Profiler, you can see how utterances resolve to intents and slots in your interaction model. When an utterance doesn't resolve to the right intent or slot, you can update the interaction model and try again. With the `cxcli`, you can see what intents Dialogflow CX considered and discarded. Then, you can determine how to use additional samples to train your model to resolve utterances to their intended intents and slots.

Every suite is executed in the Dialogflow CX session so you can test not only your NLU but also a conversation itself.

All the commands that you have available in the `cxcli` to execute the NLU profiler are located down the [`cxcli profile-nlu` command](/cmd/cxcli_profile-nlu).

## Reference

It is important to know which [suites](/nluprofiler/suites) and [tests](/nluprofiler/tests) you can build. Because of that, you can find the entire reference in the [Reference](/nluprofiler/suites) page. Suites and test are defined as `yaml` files.

The `cxcli` has a command that allows you to run these suites from your terminal or from your CI pipelines.

To execute a suite, you have to run the `cxcli profile-nlu execute` command. For the usage, please refer to this [page](/cmd/cxcli_profile-nlu_execute).

## Examples

You can find some useful examples on our [Github repo](https://github.com/xavidop/dialogflow-cx-cli/tree/master/examples) and the [Examples](/nluprofiler/examples) page.


## Execution Example

This a simple example of the `cxcli profile-nlu execute` command:

```sh
cxcli profile-nlu execute examples/suite.yaml
```

The command above will give you an output similar to this one:

```sh
$ cxcli profile-nlu execute suite.yaml
INFO Suite Information: test-agent                
INFO Test ID: test_1                              
INFO Input: type: text, value: hi                 
INFO Intent Detected: hi_intent                   
INFO Input: type: text, value: hello              
INFO Intent Detected: hi_intent                   
INFO Input: type: audio, value: ./audio/hi.mp3    
INFO Intent Detected: hi_intent                   
INFO Test ID: test_2                              
INFO Input: type: text, value: I want 3 pizzas    
INFO Intent Detected: order_intent                
INFO Param order_type: pizza                      
INFO Param number: 3                              
INFO Input: type: text, value: I want 2 cokes     
INFO Intent Detected: order_intent                
INFO Param number: 2                              
INFO Param order_type: coke                        
```

!!! info "are you running this command in a CICD pipeline?"
    If this is the case, we recommend you to execute with the `--output-format` parameter set to `json`.