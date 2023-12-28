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
# A interactions is the test itself: given an input, you will validate the agent response returned by Dialogflow CX
# You can have multiple interactions defined
interactions:
  # The ID of the interactions
  - id: test_1
    user:
      # the input type
      # it could be text, audio or prompt
      type: text
      # The input itself in text format. For type: audio, you have to specify the audio file.
      text: I want 3 pizzas
    agent:
      validate:
        # String validation to check if the response returned by Dialogflow CX is correct
        - type: contains
          value: pizza

  - id: test_2
    user: 
      type: text
      text: hi
    agent:
      # example with a string similarity validation
      validate:
        - type: similarity
          algorithm: levenshtein
          threshold: 0.4
          value: hello
          configuration-levenshtein:
            casesensitive: false

  - id: test_3
    user: 
      type: audio
      audio: ./audio/hi.mp3
    agent:
      # example with a regexp validation
      validate:
        - type: regexp
          value: '/my-regex/'
```

## Input types

### Text input

The input text is the simplest one. You just have to specify the text you want to send to Dialogflow CX. Make sure that the text is in the language you specified in the `localeId` field. to use this type you have to set the `type` field to `text` and the `text` field to the text you want to send.

### Audio input

The audio input is a little bit more complex. You have to specify the path to the audio file you want to send to Dialogflow CX. Make sure that the audio file is in the language you specified in the `localeId` field. To use this type you have to set the `type` field to `audio` and the `file` field to the path to the audio file.

It is important to know that the input audio needs to have the following format:

1. A sample rate of 16000 Hertz
2. The audio encoding has to be Linear16. Linear16 is a 16-bit linear pulse-code modulation (PCM) encoding.

If you don't have a file with this format, you can create it by yourself using the `cxcli tts` command! All the information is located [here](/tts)

### Prompt input

The prompt input is the most complex one. You have to specify the input prompt you want to send to Dialogflow CX. To use this type you have to set the `type` field to `prompt` and the `prompt` field to the prompt you want to send.

It is important to know that the input prompt needs Vertex AI API to be enabled in your project and the proper permissions granted.  You can find more information on the [Authentication](/overview/authentication) page.

## Validation types

### Contains

The contains validation type is the simplest one. It just checks if the response returned by the Dialogflow CX agent contains the value specified in the `value` field. To use this type you have to set the `type` field to `contains` and the `value` field to the value you want to check:

```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: contains
    value: pizza
```

The `contains` validation has its own options:

```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: contains
    value: pizza
    configuration-contains:
      casesensitive: true
```

If you set the `casesensitive` field to `true`, the validation will be case sensitive. By default, it is set to `false`.

### Equals

The equals validation type is a little bit more complex. It checks if the response returned by the Dialogflow CX agent is equal to the value specified in the `value` field. To use this type you have to set the `type` field to `equals` and the `value` field to the value you want to check:

```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: equals
    value: Here you have 3 pizzas
```

The `equals` validation has its own options:

```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: equals
    value: Here you have 3 pizzas
    configuration-equals:
      casesensitive: true
```

If you set the `casesensitive` field to `true`, the validation will be case sensitive. By default, it is set to `false`.

### Regexp

The regexp validation type is the most complex one. It checks if the response returned by the Dialogflow CX agent matches the regexp specified in the `value` field. To use this type you have to set the `type` field to `regexp` and the `value` field to the regular expression you want to check:

```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: regexp
    value: '/Here you have \d pizzas/'
```

The `regexp` validation has its own options:

```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: regexp
    value: '/Here you have \d pizzas/'
    configuration-regexp:
      findinsubmatches: true
```

If you set the `findinsubmatches` field to `true`, the validation will check if the regexp matches any of the submatches. By default, it is set to `false`.

### String Similarity Algorithms

The string similarity validation type is the most complex one. It checks if the response returned by the Dialogflow CX agent is similar to the value specified in the `value` field using a `threshold` to check if the similarity is enough. To use this type you have to set the `type` field to `similarity` and the `algorithm` field to the algorithm you want to use and the `value` field to the value you want to check and the `threshold` field to the threshold you want to use. The available algorithms are:

1. `levenshtein`: The Levenshtein distance is a string metric for measuring the difference between two sequences. Informally, the Levenshtein distance between two words is the minimum number of single-character edits (insertions, deletions or substitutions) required to change one word into the other.
2. `jaro`: The Jaro distance is a measure of similarity between two strings. The higher the Jaro distance for two strings is, the more similar the strings are. The score is normalized such that 0 equates to no similarity and 1 is an exact match.
3. `jaro-winkler`: The Jaro–Winkler distance is a string metric for measuring the edit distance between two sequences. Informally, the Jaro–Winkler distance is the edit distance between two strings with the twist that higher scores are returned for strings that match from the beginning for a set prefix length.
4. `smith-waterman-gotoh`: The Smith–Waterman algorithm performs local sequence alignment; that is, for determining similar regions between two strings of nucleic acid sequences or protein sequences. Instead of looking at the entire sequence, the Smith–Waterman algorithm compares segments of all possible lengths and optimizes the similarity measure.
5. `sorensen-dice`: The Dice distance is a measure of similarity between two strings. The higher the Dice distance for two strings is, the more similar the strings are. The score is normalized such that 0 equates to no similarity and 1 is an exact match.
6. `jaccard`: The Jaccard distance is a measure of similarity between two strings. The higher the Jaccard distance for two strings is, the more similar the strings are. The score is normalized such that 0 equates to no similarity and 1 is an exact match.
7. `overlap-coefficient`: The Overlap distance (or Szymkiewicz-Simpson distance) is a measure of similarity between two strings. The higher the Overlap distance for two strings is, the more similar the strings are. The score is normalized such that 0 equates to no similarity and 1 is an exact match.
8.  `hamming`: The Hamming measures the minimum number of substitutions required to change one string into the other, or equivalently, the minimum number of errors that could have transformed one string into the other. The higher the Hamming distance for two strings is, the more similar the strings are. The score is normalized such that 0 equates to no similarity and 1 is an exact match.

#### Levenshtein

The `levenshtein` algorithm has the following configuration:

```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: levenshtein
    value: hello
    threshold: 0.4
```

The options are:
```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: levenshtein
    value: hello
    threshold: 0.4
    configuration-levenshtein:
      casesensitive: true
      insertcost: 1
      deletecost: 1
      replacecost: 1
```
Let's explain each one of them:
1. The `casesensitive` field is used to check if the algorithm is case sensitive. By default, it is set to `true`.
2. the `insertcost` field is used to set the cost of an insert operation. By default, it is set to `1`.
3. the `deletecost` field is used to set the cost of a delete operation. By default, it is set to `1`.
4. the `replacecost` field is used to set the cost of a replace operation. By default, it is set to `1`.

#### Jaro

The `Jaro` algorithm has the following configuration:

```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: jaro
    value: hello
    threshold: 0.4
```

The options are:
```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: jaro
    value: hello
    threshold: 0.4
    configuration-jaro:
      casesensitive: true
```

The `casesensitive` field is used to check if the algorithm is case sensitive. By default, it is set to `true`.

#### Jaro-Winkler

The `Jaro-Winkler` algorithm has the following configuration:

```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: jaro-winkler
    value: hello
    threshold: 0.4
```

The options are:
```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: jaro-winkler
    value: hello
    threshold: 0.4
    configuration-jaro-winkler:
      casesensitive: true
```

The `casesensitive` field is used to check if the algorithm is case sensitive. By default, it is set to `true`.

#### Smith-Waterman-Gotoh

The `Smith-Waterman-Gotoh` algorithm has the following configuration:

```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: smith-waterman-gotoh
    value: hello
    threshold: 0.4
```

The options are:
```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: smith-waterman-gotoh
    value: hello
    threshold: 0.4
    configuration-smith-waterman-gotoh:
      casesensitive: true
      gappenalty: -0.5
```

Let's explain each one of them:
1. The `casesensitive` field is used to check if the algorithm is case sensitive. By default, it is set to `true`.
2. the `gappenalty` defines a score penalty for character insertions or deletions. For relevant results, the gap penalty should be a non-positive number. By default, it is set to `-0.5`.

#### Sorensen-Dice

The `Sorensen-Dice` algorithm has the following configuration:

```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: sorensen-dice
    value: hello
    threshold: 0.4
```

The options are:
```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: sorensen-dice
    value: hello
    threshold: 0.4
    configuration-sorensen-dice:
      casesensitive: true
      NgramSize: 2
```

Let's explain each one of them:
1. The `casesensitive` field is used to check if the algorithm is case sensitive. By default, it is set to `true`.
2. The `NgramSize` represents the size (in characters) of the tokens generated when comparing the input sequences. By default, it is set to `2`.

#### Jaccard

The `Jaccard` algorithm has the following configuration:

```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: jaccard
    value: hello
    threshold: 0.4
```

The options are:
```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: jaccard
    value: hello
    threshold: 0.4
    configuration-jaccard:
      casesensitive: true
      NgramSize: 2
```

Let's explain each one of them:
1. The `casesensitive` field is used to check if the algorithm is case sensitive. By default, it is set to `true`.
2. The `NgramSize` represents the size (in characters) of the tokens generated when comparing the input sequences. By default, it is set to `2`.

#### Overlap Coefficient

The `Overlap Coefficient` algorithm has the following configuration:

```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: overlap-coefficient
    value: hello
    threshold: 0.4
```

The options are:
```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: overlap-coefficient
    value: hello
    threshold: 0.4
    configuration-overlap-coefficient:
      casesensitive: true
      NgramSize: 2
```

Let's explain each one of them:
1. The `casesensitive` field is used to check if the algorithm is case sensitive. By default, it is set to `true`.
2. The `NgramSize` represents the size (in characters) of the tokens generated when comparing the input sequences. By default, it is set to `2`.

#### Hamming

The `Hamming` algorithm has the following configuration:

```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: hamming
    value: hello
    threshold: 0.4
```

The options are:
```yaml
validate:
  # String validation to check if the response returned by Dialogflow CX is correct
  - type: similarity
    algorithm: hamming
    value: hello
    threshold: 0.4
    configuration-hamming:
      casesensitive: true
```

The `casesensitive` field is used to check if the algorithm is case sensitive. By default, it is set to `true`.

## JSON Schema

`cxcli` also has a [jsonschema](http://json-schema.org/draft/2020-12/json-schema-validation.html) file, which you can use to have better
editor support:

```sh
https://cxcli.xavidop.me/static/conversationtest.json
```

You can also specify it in your `yml` config files by adding a
comment like the following:
```yaml
# yaml-language-server: $schema=https://cxcli.xavidop.me/static/conversationtest.json
```
