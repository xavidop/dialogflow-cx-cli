# Text-to-speech

<p align="center">
  <img alt="GCP TTS Logo" src="/images/tts.png" height="140" />
</p>


`cxcli` has some commands that allows you to interact with Google Cloud Text to Speech service using the `Cloud Text-to-Speech API`!

!!! info "Is this your first time using this feature?"
    Before you start using this functionality, please, read the [authentication](/overview/authentication) page.

## Usage

You can find the text-to-speech commands usage down the `cxcli tts` command. You can read the documentation about this command [here](/cmd/cxcli_tts).

The `cxcli tts` root command has the `synthesize` command. You can find the usage of this command [here](/cmd/cxcli_tts_synthesize).

### Parameters

These are the relevant parameters that you can use to interact with Google Cloud tts:

1. `locale`: the locale accepts all the locales accepted by the Google `Cloud Text-to-speech API`. You can find all the locales available [here](https://cloud.google.com/text-to-speech/docs/voices)
2. `output-file`: mp3 file where we are going to write the synthesize text

### Output

It is important to know that the output will have this format:

1. A Sample Rate Hertz of 16000Hz
2. The audio encoding will be Linear16. Linear16 is a 16-bit linear pulse-code modulation (PCM) encoding.


## Example

This a simple example of the `cxcli tts synthesize` command:

```sh
cxcli tts synthesize hi --locale en-US --output-file hi.mp3
```

The command above will give you an audio file like this one:

<audio controls>
  <source src="/static/hi.mp3" type="audio/mpeg">
Your browser does not support the audio element.
</audio>

You can download the audio file [here](/static/hi.mp3)