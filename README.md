[![Goreleaser](https://github.com/xavidop/dialogflow-cx-test-runner/actions/workflows/release_build.yml/badge.svg)](https://github.com/xavidop/dialogflow-cx-test-runner/actions/workflows/release_build.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/xavidop/dialogflow-cx-test-runner)](https://goreportcard.com/report/github.com/xavidop/dialogflow-cx-test-runner)

# CX CLI
Dialogflow utility to test your Dialogflow CX Project

<!-- TOC -->

- [CX CLI](#cx-cli)
  - [Installation](#installation)
    - [Homebrew](#homebrew)
- [Usage](#usage)
  - [Profile NLU Test Suites](#profile-nlu-test-suites)
  - [CICD for your Dialogflow CX environments](#cicd-for-your-dialogflow-cx-environments)
  - [Text-to-Speech](#text-to-speech)

<!-- /TOC -->

## Installation

You can download the latest release from [here](https://github.com/xavidop/dialogflow-cx-test-runner/releases)

### Homebrew

If you use the package manager Homebrew, you can install this utility following these steps:

1. Add my Hombre tab:
```bash
brew tap xavidop/tap git@github.com:xavidop/homebrew-tap.git
brew update
```
1. Install the CLI:
```bash
brew install cxcli
```
# Usage

```bash
Usage:
  cxcli [flags]
  cxcli [command]

Available Commands:
  cicd        Actions on CICD testings
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  profile-nlu Actions on testing
  tts         Actions on text-to-speech commands
  version     Get cxcli version

Flags:
  -c, --credentials string   verbose error output (with stack trace)
  -h, --help                 help for cxcli
  -o, --output string        Output Format (default text)
  -u, --skip-update-check    Skip the check for updates check run before every command
  -v, --verbose              verbose error output (with stack trace)
```

## Profile NLU Test Suites

To run test suites, you can execute the following command

```bash
Usage:
  cxcli profile-nlu execute [suite-file] [flags]

Aliases:
  execute, execute, e, exe, exec

Flags:
  -h, --help   help for execute

Global Flags:
  -c, --credentials string   verbose error output (with stack trace)
  -o, --output string        Output Format (default "text")
  -u, --skip-update-check    Skip the check for updates check run before every command
  -v, --verbose              verbose error output (with stack trace)
```

Example: `cxcli profile-nlu execute "examples/suite.yaml" --credentials credentials.json`

You can find a suite example [here](/examples/)

## CICD for your Dialogflow CX environments

To execute a CICD pipeline you can execute the following command

```bash
Usage:
  cxcli cicd execute [environment] [flags]

Aliases:
  execute, execute, e, exe, exec

Flags:
  -a, --agent-name string    Dialogflow CX Agent Name
  -h, --help                 help for execute
  -l, --location-id string   Dialogflow CX Location ID of the Project
  -p, --project-id string    Dialogflow CX Project ID

Global Flags:
  -c, --credentials string   verbose error output (with stack trace)
  -o, --output string        Output Format (default text)
  -u, --skip-update-check    Skip the check for updates check run before every command
  -v, --verbose              verbose error output (with stack trace)
```

Example: `cxcli cicd execute "<your-env>" --project-id <your-project-id> --location-id <your-location-id> --agent-name <your-agent-name> --credentials, credentials.json`

## Text-to-Speech

To transform your text to speech, you can execute the following command:

```bash
Usage:
  cxcli tts synthesize [input] [flags]

Aliases:
  synthesize, synth, s

Flags:
  -h, --help            help for synthesize
  -i, --input string    Input text to synthesize
  -l, --locale string   Input locale
  -o, --output string   Output file name (default "output.mp3")

Global Flags:
  -c, --credentials string   verbose error output (with stack trace)
  -u, --skip-update-check    Skip the check for updates check run before every command
  -v, --verbose              verbose error output (with stack trace)
```

Example: `cxcli tts synthesize "hi" -l en-US -o output.mp3 --credentials credentials.json`