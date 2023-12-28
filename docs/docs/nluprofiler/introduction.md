# NLU profiler

## What is this?

<p align="center">
  <img alt="Flow" src="/images/nlu.png" />
</p>

Use the NLU Profiler to test user utterances and improve your agent's interaction model.

With the NLU Profiler, you can see how utterances resolve to intents and slots in your interaction model. When an utterance doesn't resolve to the right intent or slot, you can update the interaction model and try again. With the `cxcli`, you can see which intents Dialogflow CX considered and discarded. Then, you can determine how to use additional samples to train your model to resolve utterances to their desired intents and slots.

Every suite is executed in the Dialogflow CX session so you can test not only your NLU but also a conversation itself.

All of the commands that are available in `cxcli` to execute the NLU profiler are located within the [`cxcli profile-nlu` subcommand](/cmd/cxcli_profile-nlu).

## Reference

It is important to know which [suites](/nluprofiler/suites) and [tests](/nluprofiler/tests) you can build. Because of that, you can find the entire reference on the [Reference](/nluprofiler/suites) page. Suites and test are defined as `yaml` files.

The `cxcli` has a command that allows you to run these suites from your terminal or from your CI pipelines.

To execute a suite, you can run the `cxcli profile-nlu execute` command. For the usage, please refer to this [page](/cmd/cxcli_profile-nlu_execute).

## Examples

You can find some useful examples on our [GitHub repo](https://github.com/xavidop/dialogflow-cx-cli/tree/master/examples/profilenlu) and the [Examples](/nluprofiler/examples) page.


## Execution Example

Here is a simple example of the `cxcli profile-nlu execute` command:

```sh
cxcli profile-nlu execute examples/profilenlu/suite.yaml
```

The above command will give you output similar to the following:

```sh
$ cxcli profile-nlu execute suite.yaml
[INFO] Running suite: Example NLU Profiler Suite
[INFO][test-file:test_1][check:test_1_1][input:prompt] User> hi (auto-generated from prompt: "give me a one line hello without exclamation mark using only 2 characters and lowercase")
[INFO][test-file:test_1][check:test_1_1][input:prompt] Agent> Hi! How are you doing?
[INFO][test-file:test_1][check:test_1_1][validation:hi_intent] Intent Detected: hi_intent
[INFO][test-file:test_1][check:test_1_2][input:text] User> hello
[INFO][test-file:test_1][check:test_1_2][input:text] Agent> Hi! How are you doing?
[INFO][test-file:test_1][check:test_1_2][validation:hi_intent] Intent Detected: hi_intent
[INFO][test-file:test_1][check:test_1_3][input:audio] User> ./audio/hi.mp3
[INFO][test-file:test_1][check:test_1_3][input:audio] Agent> Hi! How are you doing?
[INFO][test-file:test_1][check:test_1_3][validation:hi_intent] Intent Detected: hi_intent
[INFO][test-file:test_2][check:test_2_1][input:text] User> I want 3 pizzas
[INFO][test-file:test_2][check:test_2_1][input:text] Agent> 
[INFO][test-file:test_2][check:test_2_1][validation:order_intent] Intent Detected: order_intent
[INFO][test-file:test_2][check:test_2_1][validation:order_intent] Param order_type: pizza 
[INFO][test-file:test_2][check:test_2_1][validation:order_intent] Param number: 3 
[INFO][test-file:test_2][check:test_2_2][input:text] User> I want 2 cokes
[INFO][test-file:test_2][check:test_2_2][input:text] Agent> 
[INFO][test-file:test_2][check:test_2_2][validation:order_intent] Intent Detected: order_intent
[INFO][test-file:test_2][check:test_2_2][validation:order_intent] Param number: 2 
[INFO][test-file:test_2][check:test_2_2][validation:order_intent] Param order_type: coke 
```

!!! info "Are you running this command in a CI/CD pipeline?"
    If this is the case, we recommend that you set the `--output-format` parameter to `json`.
