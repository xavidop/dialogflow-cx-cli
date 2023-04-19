# Speech-to-text

<p align="center">
  <img alt="Google Cloud TTS Logo" src="/images/stt.png" style="height:256px;width:256px" />
</p>

The `cxcli` tool has various commands that allow you to interact with Google Cloud's Speech to Text service using the `Cloud Speech-to-text API`!

!!! info "Is this your first time using this feature?"
    Before you start using this functionality, please, read the [authentication](/overview/authentication) page.

## Usage

You can find the speech-to-text functionality within the `cxcli stt` subcommand. You can read the documentation about this command [here](/cmd/cxcli_stt).

The `cxcli stt` command has a `recognize` subcommand. You can find the usage of this command [here](/cmd/cxcli_stt_recognize).

### Parameters

These are the relevant parameters that you can use to interact with Google Cloud STT:

1. `locale`: this parameter accepts all of the locales that are available in the Google Cloud `Speech-to-text API`. You can find all the locales available [here](https://cloud.google.com/speech-to-text/docs/speech-to-text-supported-languages).

### Audio input file

It is important to know that the input audio needs to be in the following format:

1. A sample rate of 16000 Hertz
2. The audio encoding has to be Linear16. Linear16 is a 16-bit linear pulse-code modulation (PCM) encoding.

If you don't have a file with this format, you can create it by yourself using the `cxcli tts` command! All of the relevant information is located [here](/tts).

## Example

Here is a simple example of the `cxcli stt recognize` command:

```sh
cxcli stt recognize hi.mp3  --locale en-US
```

The above command will give you output similar to the following:

```sh
$ cxcli stt recognize hi.mp3 --locale en-US --verbose
INFO Duration time: 570 miliseconds
INFO Detections: 1
INFO 1. Text detected: hi
INFO 1. Confidence: 79.276474%
```

!!! info "Are you running this command in a CI/CD pipeline?"
    If this is the case, we recommend that you set the `--output-format` parameter to `json`.
