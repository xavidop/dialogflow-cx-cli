package cx

import (
	"context"
	"errors"
	"fmt"
	"strings"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/internal/utils"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func CreateGeneratorRESTClient(locationId string) (*cx.GeneratorsClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewGeneratorsRESTClient(ctx, credentials, endpoint)
	} else {
		return cx.NewGeneratorsRESTClient(ctx, endpoint)
	}

}

func CreateGeneratorGRPCClient(locationId string) (*cx.GeneratorsClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewGeneratorsClient(ctx, credentials, endpoint)
	} else {
		return cx.NewGeneratorsClient(ctx, endpoint)
	}

}

func CreateGenerator(generatorClient *cx.GeneratorsClient, agent *cxpb.Agent, name, prompt, localeID string) (*cxpb.Generator, error) {
	ctx := context.Background()

	placeholders, err := ParsePlaceHolders(prompt)
	if err != nil {
		return nil, err
	}

	generator := &cxpb.Generator{
		DisplayName: name,
		PromptText: &cxpb.Phrase{
			Text: prompt,
		},
		Placeholders: placeholders,
	}

	reqCreateGenerator := &cxpb.CreateGeneratorRequest{
		Parent:       agent.GetName(),
		Generator:    generator,
		LanguageCode: localeID,
	}

	return generatorClient.CreateGenerator(ctx, reqCreateGenerator)
}

func ParsePlaceHolders(prompt string) ([]*cxpb.Generator_Placeholder, error) {
	placeholders := []*cxpb.Generator_Placeholder{}
	words := strings.Split(prompt, " ")

	for _, word := range words {
		if strings.HasPrefix(word, "$") {
			finalWord := strings.ReplaceAll(word, "$", "")
			placeholder := &cxpb.Generator_Placeholder{
				Id:   finalWord,
				Name: finalWord,
			}
			placeholders = append(placeholders, placeholder)
		}
	}

	return placeholders, nil
}

func UpdateGenerator(generatorClient *cx.GeneratorsClient, agent *cxpb.Agent, name, prompt, localeID string) (*cxpb.Generator, error) {
	ctx := context.Background()

	generator, err := GetGeneratorIdByName(generatorClient, agent, name)
	if err != nil {
		return nil, err
	}

	placeholders, err := ParsePlaceHolders(prompt)
	if err != nil {
		return nil, err
	}
	paths := []string{}

	generator.PromptText.Text = prompt
	paths = append(paths, "prompt_text")

	generator.Placeholders = placeholders
	paths = append(paths, "placeholders")

	reqUpdateGenerator := &cxpb.UpdateGeneratorRequest{
		Generator:    generator,
		LanguageCode: localeID,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: paths,
		},
	}

	return generatorClient.UpdateGenerator(ctx, reqUpdateGenerator)
}

func GetGeneratorIdByName(generatorClient *cx.GeneratorsClient, agent *cxpb.Agent, name string) (*cxpb.Generator, error) {
	ctx := context.Background()

	reqGeneratorsList := &cxpb.ListGeneratorsRequest{
		Parent: agent.GetName(),
	}

	generators := generatorClient.ListGenerators(ctx, reqGeneratorsList)

	for generator, err := generators.Next(); err == nil; {
		if generator.DisplayName == name {
			return generator, nil
		}
		generator, err = generators.Next()
		if err != nil {
			return nil, err
		}

	}

	return nil, errors.New("generator not found")

}

func DeleteGenerator(generatorClient *cx.GeneratorsClient, agent *cxpb.Agent, name, force string) error {
	ctx := context.Background()

	generator, err := GetGeneratorIdByName(generatorClient, agent, name)
	if err != nil {
		return err
	}

	reqDeleteGenerator := &cxpb.DeleteGeneratorRequest{
		Name: generator.GetName(),
	}

	if force != "" {
		forceBool, err := utils.ParseBool(force)
		if err != nil {
			return err
		}
		reqDeleteGenerator.Force = forceBool
	}

	return generatorClient.DeleteGenerator(ctx, reqDeleteGenerator)
}
