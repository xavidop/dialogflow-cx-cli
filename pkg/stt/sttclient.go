package stt

import (
	"context"
	"os"

	speechtotext "cloud.google.com/go/speech/apiv1"
	tspeechtotextpb "cloud.google.com/go/speech/apiv1/speechpb"

	"github.com/xavidop/dialogflow-cx-test-runner/internal/global"
	"google.golang.org/api/option"
)

func CreateSTTClient() (*speechtotext.Client, error) {
	ctx := context.Background()

	credentials := option.WithCredentialsFile(global.Credentials)

	client, err := speechtotext.NewClient(ctx, credentials)
	if err != nil {
		return nil, err
	}

	return client, nil

}

func ExecuteRecognition(sttClient *speechtotext.Client, input string, locale string) (*tspeechtotextpb.RecognizeResponse, error) {
	ctx := context.Background()

	content, err := os.ReadFile(input)
	if err != nil {
		return nil, err
	}

	req := &tspeechtotextpb.RecognizeRequest{
		Config: &tspeechtotextpb.RecognitionConfig{
			Encoding:        tspeechtotextpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 16000,
			LanguageCode:    locale,
		},
		Audio: &tspeechtotextpb.RecognitionAudio{
			AudioSource: &tspeechtotextpb.RecognitionAudio_Content{Content: content},
		},
	}
	return sttClient.Recognize(ctx, req)

}
