package tts

import (
	"context"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"google.golang.org/api/option"
)

func CreateTTSClient() (*texttospeech.Client, error) {
	ctx := context.Background()

	credentials := option.WithCredentialsFile(global.Credentials)

	client, err := texttospeech.NewClient(ctx, credentials)
	if err != nil {
		return nil, err
	}

	return client, nil

}

func ExecuteSynthesize(ttsClient *texttospeech.Client, input string, locale string) (*texttospeechpb.SynthesizeSpeechResponse, error) {
	ctx := context.Background()

	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: input},
		},

		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: locale,
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding:   texttospeechpb.AudioEncoding_LINEAR16,
			SampleRateHertz: 16000,
		},
	}

	return ttsClient.SynthesizeSpeech(ctx, &req)

}
