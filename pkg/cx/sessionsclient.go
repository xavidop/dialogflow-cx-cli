package cx

import (
	"context"
	"fmt"
	"os"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"

	"github.com/google/uuid"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"google.golang.org/api/option"
)

func CreateSessionsClient(locationId string) (*cx.SessionsClient, error) {
	ctx := context.Background()

	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewSessionsRESTClient(ctx, credentials, endpoint)

	} else {
		return cx.NewSessionsRESTClient(ctx, endpoint)

	}
}

func DetectIntentFromText(sessionClient *cx.SessionsClient, agent *cxpb.Agent, localeId string, input string) (*cxpb.DetectIntentResponse, error) {

	textInput := cxpb.TextInput{Text: input}
	queryTextInput := cxpb.QueryInput_Text{Text: &textInput}
	queryInput := cxpb.QueryInput{
		Input:        &queryTextInput,
		LanguageCode: localeId,
	}

	response, err := DetectIntent(sessionClient, agent, queryInput)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func DetectIntentFromAudio(sessionClient *cx.SessionsClient, agent *cxpb.Agent, localeId string, audioFile string) (*cxpb.DetectIntentResponse, error) {

	dat, err := os.ReadFile(audioFile)
	if err != nil {
		return nil, err
	}

	audioInputConfig := cxpb.InputAudioConfig{
		AudioEncoding:   cxpb.AudioEncoding_AUDIO_ENCODING_LINEAR_16,
		SampleRateHertz: 16000,
	}

	audioInput := cxpb.AudioInput{
		Audio:  dat,
		Config: &audioInputConfig,
	}
	queryAudioInput := cxpb.QueryInput_Audio{
		Audio: &audioInput,
	}
	queryInput := cxpb.QueryInput{
		Input:        &queryAudioInput,
		LanguageCode: localeId,
	}

	response, err := DetectIntent(sessionClient, agent, queryInput)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func DetectIntent(sessionClient *cx.SessionsClient, agent *cxpb.Agent, queryinput cxpb.QueryInput) (*cxpb.DetectIntentResponse, error) {
	ctx := context.Background()

	sessionPath := fmt.Sprintf("%s/sessions/%s", agent.GetName(), uuid.NewString())

	request := cxpb.DetectIntentRequest{Session: sessionPath, QueryInput: &queryinput}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
