# Speech-to-text

<p align="center">
  <img alt="GCP TTS Logo" src="/images/stt.png" style="height:256px;width:256px" />
</p>



`cxcli` has some commands that allows you to interact with Google Cloud Text to Speech service using the `Cloud Speech-to-text API`!

!!! info "Is this your first time using this feature?"
    Before you start using this functionality, please, read the [authentication](/overview/authentication) page.

## Usage

You can find the speech-to-text commands usage down the `cxcli stt` command. You can read the documentation about this command [here](/cmd/cxcli_stt).

The `cxcli stt` root command has the `recognize` command. You can find the usage of this command [here](/cmd/cxcli_stt_recognize).

### Parameters

These are the relevant parameters that you can use to interact with Google Cloud stt:

1. `locale`: the locale accepts all the locales accepted by the Google `Cloud Speech-to-text API`. You can find all the locales available [here](https://cloud.google.com/speech-to-text/docs/speech-to-text-supported-languages)

### Audio input file

It is important to know that the input has to have this format:

1. A Sample Rate Hertz of 16000Hz
2. The audio encoding has to be be Linear16. Linear16 is a 16-bit linear pulse-code modulation (PCM) encoding.

If you don't have a file with this format, you can create it by yourself using the `cxcli tts` command! All the information is located [here](/tts)

## Example

This a simple example of the `cxcli stt recognize` command:

```sh
cxcli stt recognize hi.mp3  --locale en-US
```

The command above will give you an audio file like[this one:

```sh
$ cxcli stt recognize hi.mp3 --locale en-US --verbose
INFO Duration time: 570 miliseconds               
INFO Detections: 1                                
INFO 1. Text detected: hi                         
INFO 1. Confidence: 79.276474%                     
```

!!! info "are you running this command in a CICD pipeline?"
    If this is the case, we recommend you to execute with the `--output-format` parameter set to `json`.