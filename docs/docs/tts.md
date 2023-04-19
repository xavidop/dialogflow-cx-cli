# Text-to-speech

<p align="center">
  <img alt="Google Cloud TTS Logo" src="/images/tts.png" height="140" />
</p>


The `cxcli` tool has various commands that allow you to interact with Google Cloud's Text to Speech service using the `Cloud Text-to-Speech API`!

!!! info "Is this your first time using this feature?"
    Before you start using this functionality, please, read the [authentication](/overview/authentication) page.

## Usage

You can find the text-to-speech functionality within the `cxcli tts` subcommand. You can read the documentation about this command [here](/cmd/cxcli_tts).

The `cxcli tts` command has a `synthesize` subcommand. You can find the usage of this command [here](/cmd/cxcli_tts_synthesize).

### Parameters

These are the relevant parameters that you can use to interact with Google Cloud TTS:

1. `locale`: this parameter accepts all of the locales that are available in the Google Cloud `Text-to-speech API`. You can find all the locales available [here](https://cloud.google.com/text-to-speech/docs/voices).
2. `output-file`: MP3 audio file where we are going to output the synthesized text.

### Output

It is important to know that the output audio will have the following format:

1. A sample rate of 16000 Hertz
2. The audio encoding will be Linear16. Linear16 is a 16-bit linear pulse-code modulation (PCM) encoding.


## Example

Here is a simple example of the `cxcli tts synthesize` command:

```sh
cxcli tts synthesize hi --locale en-US --output-file hi.mp3
```

The above command will give you an audio file similar to the following:

<audio controls>
  <source src="/static/hi.mp3" type="audio/mpeg">
Your browser does not support the audio element.
</audio>

You can download the audio file [here](/static/hi.mp3).
