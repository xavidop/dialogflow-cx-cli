package cx

import (
	"context"
	"errors"
	"fmt"
	"strings"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"

	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func CreateIntentRESTClient(locationId string) (*cx.IntentsClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewIntentsRESTClient(ctx, credentials, endpoint)
	} else {
		return cx.NewIntentsRESTClient(ctx, endpoint)
	}

}

func CreateIntentGRPCClient(locationId string) (*cx.IntentsClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewIntentsClient(ctx, credentials, endpoint)
	} else {
		return cx.NewIntentsClient(ctx, endpoint)
	}

}

func CreateIntent(intentClient *cx.IntentsClient, agent *cxpb.Agent, intentName, description, localeId string, trainingPhrases []string, entityTypesClient *cx.EntityTypesClient) (*cxpb.Intent, error) {
	ctx := context.Background()
	localeToUse := agent.GetDefaultLanguageCode()
	if localeId != "" {
		localeToUse = localeId
	}

	intentTrainingPhrases, intentTrainingParameters, err := CreateIntentTrainingPhrases(trainingPhrases, entityTypesClient, agent)
	if err != nil {
		return nil, err
	}

	reqCreateIntent := &cxpb.CreateIntentRequest{
		Parent:       agent.GetName(),
		LanguageCode: localeToUse,
		Intent: &cxpb.Intent{
			DisplayName:     intentName,
			Description:     description,
			TrainingPhrases: intentTrainingPhrases,
			Parameters:      intentTrainingParameters,
		},
	}

	return intentClient.CreateIntent(ctx, reqCreateIntent)
}

func UpdateIntent(intentClient *cx.IntentsClient, agent *cxpb.Agent, intentName, description, localeId string, trainingPhrases []string, entityTypesClient *cx.EntityTypesClient) (*cxpb.Intent, error) {
	ctx := context.Background()
	paths := []string{}

	localeToUse := agent.GetDefaultLanguageCode()
	if localeId != "" {
		localeToUse = localeId
	}

	intentTrainingPhrases, intentTrainingParameters, err := CreateIntentTrainingPhrases(trainingPhrases, entityTypesClient, agent)
	if err != nil {
		return nil, err
	}

	intent, err := GetIntentIdByName(intentClient, agent, intentName)
	if err != nil {
		return nil, err
	}

	if description != "" {
		intent.Description = description
		paths = append(paths, "description")
	}

	if len(intentTrainingPhrases) > 0 {
		intent.TrainingPhrases = intentTrainingPhrases
		paths = append(paths, "training_phrases")
	}

	if len(intentTrainingParameters) > 0 {
		intent.Parameters = intentTrainingParameters
		paths = append(paths, "parameters")
	}

	reqUpdateIntent := &cxpb.UpdateIntentRequest{
		LanguageCode: localeToUse,
		Intent:       intent,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: paths,
		},
	}

	return intentClient.UpdateIntent(ctx, reqUpdateIntent)
}

func CreateIntentTrainingPhrases(trainingPhrases []string, entityTypesClient *cx.EntityTypesClient, agent *cxpb.Agent) ([]*cxpb.Intent_TrainingPhrase, []*cxpb.Intent_Parameter, error) {
	intentTrainingPhrases := []*cxpb.Intent_TrainingPhrase{}
	intentTrainingParameters := []*cxpb.Intent_Parameter{}

	for _, trainingPhrase := range trainingPhrases {
		trainingPhrase = strings.TrimSpace(trainingPhrase)
		var intentTrainingPhrase *cxpb.Intent_TrainingPhrase

		// If the training phrases contains an entity
		if strings.Contains(trainingPhrase, "@") {

			intentTrainingPhrase = &cxpb.Intent_TrainingPhrase{
				RepeatCount: 1,
			}

			intentTrainingPhraseParts := []*cxpb.Intent_TrainingPhrase_Part{}

			for _, word := range strings.Split(trainingPhrase, " ") {
				token := word
				entity := ""
				paramId := ""
				entityType := ""
				if strings.Contains(word, "@") {
					token = strings.Split(word, "@")[0]
					entity = strings.Split(word, "@")[1]
				}

				// Word with entity associated
				if entity != "" {

					// System entity
					if strings.Contains(entity, "sys.") {
						entityType = "projects/-/locations/-/agents/-/entityTypes/" + entity
						paramId = strings.ReplaceAll(entity, "sys.", "")
					} else {
						// User created entity
						paramId = entity
						ent, err := GetEntityTypeIdByName(entityTypesClient, agent, entity)
						if err != nil {
							return nil, nil, err
						}
						entityType = ent.GetName()
					}
					intentParameter := &cxpb.Intent_Parameter{
						Id:         paramId,
						EntityType: entityType,
					}
					intentTrainingParameters = append(intentTrainingParameters, intentParameter)
				}

				intentTrainingPhrasePart := &cxpb.Intent_TrainingPhrase_Part{
					Text:        token + " ",
					ParameterId: paramId,
				}

				intentTrainingPhraseParts = append(intentTrainingPhraseParts, intentTrainingPhrasePart)
			}

			intentTrainingPhrase.Parts = intentTrainingPhraseParts

		} else {
			// If there is not an entity associated, we just add the training phrase as text
			intentTrainingPhrase = &cxpb.Intent_TrainingPhrase{
				Parts: []*cxpb.Intent_TrainingPhrase_Part{
					{
						Text: trainingPhrase,
					},
				},
				RepeatCount: 1,
			}
		}

		intentTrainingPhrases = append(intentTrainingPhrases, intentTrainingPhrase)
	}
	return intentTrainingPhrases, intentTrainingParameters, nil
}

func DeleteIntent(intentClient *cx.IntentsClient, agent *cxpb.Agent, intentName string) error {
	ctx := context.Background()

	intent, err := GetIntentIdByName(intentClient, agent, intentName)
	if err != nil {
		return err
	}

	reqDeleteIntent := &cxpb.DeleteIntentRequest{
		Name: intent.GetName(),
	}
	return intentClient.DeleteIntent(ctx, reqDeleteIntent)
}

func GetIntentIdByName(intentClient *cx.IntentsClient, agent *cxpb.Agent, intentName string) (*cxpb.Intent, error) {
	ctx := context.Background()

	reqIntentList := &cxpb.ListIntentsRequest{
		Parent: agent.GetName(),
	}

	intents := intentClient.ListIntents(ctx, reqIntentList)

	for intent, err := intents.Next(); err == nil; {
		if intent.DisplayName == intentName {
			return intent, nil
		}
		intent, err = intents.Next()
		if err != nil {
			return nil, err
		}

	}

	return nil, errors.New("intent not found")

}
