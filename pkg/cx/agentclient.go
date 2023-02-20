package cx

import (
	"context"
	"errors"
	"fmt"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	"github.com/xavidop/dialogflow-cx-cli/cmd/agent/types"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/internal/utils"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func CreateAgentRESTClient(locationId string) (*cx.AgentsClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewAgentsRESTClient(ctx, credentials, endpoint)
	} else {
		return cx.NewAgentsRESTClient(ctx, endpoint)
	}

}

func CreateAgentGRPCClient(locationId string) (*cx.AgentsClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewAgentsClient(ctx, credentials, endpoint)
	} else {
		return cx.NewAgentsClient(ctx, endpoint)
	}
}

func GetAgentIdByName(agentClient *cx.AgentsClient, agentName string, projectId string, locationId string) (*cxpb.Agent, error) {
	ctx := context.Background()

	parentPath := fmt.Sprintf("projects/%s/locations/%s", projectId, locationId)

	reqAgentList := &cxpb.ListAgentsRequest{
		Parent: parentPath,
	}

	agents := agentClient.ListAgents(ctx, reqAgentList)

	for agent, err := agents.Next(); err == nil; {
		if agent.DisplayName == agentName {
			return agent, nil
		}
		agent, err = agents.Next()
		if err != nil {
			return nil, err
		}

	}

	return nil, errors.New("agent not found")
}

func ExportAgentById(agentClient *cx.AgentsClient, agentId, exportFormat string) (*cxpb.ExportAgentResponse, error) {
	ctx := context.Background()

	var format cxpb.ExportAgentRequest_DataFormat
	switch exportFormat {
	case "blob":
		format = cxpb.ExportAgentRequest_BLOB
	case "json":
		format = cxpb.ExportAgentRequest_JSON_PACKAGE
	default:
		format = cxpb.ExportAgentRequest_DATA_FORMAT_UNSPECIFIED
	}

	reqAgentExport := &cxpb.ExportAgentRequest{
		Name:       agentId,
		DataFormat: format,
	}

	longRunningOperation, err := agentClient.ExportAgent(ctx, reqAgentExport)
	if err != nil {
		return nil, err
	}

	return longRunningOperation.Wait(ctx)

}

func RestoreAgentById(agentClient *cx.AgentsClient, agentId string, agentContent []byte) error {
	ctx := context.Background()

	reqAgentRestore := &cxpb.RestoreAgentRequest{
		Name: agentId,
		Agent: &cxpb.RestoreAgentRequest_AgentContent{
			AgentContent: agentContent,
		},
	}

	longRunningOperation, err := agentClient.RestoreAgent(ctx, reqAgentRestore)
	if err != nil {
		return err
	}

	return longRunningOperation.Wait(ctx)

}

func CreateAgent(agentClient *cx.AgentsClient, agentName, locationID, projectID string, createInput *types.CreateUpdateAgent) (*cxpb.Agent, error) {
	ctx := context.Background()
	parentPath := fmt.Sprintf("projects/%s/locations/%s", projectID, locationID)

	advancedSettings := &cxpb.AdvancedSettings{
		LoggingSettings: &cxpb.AdvancedSettings_LoggingSettings{},
	}

	if createInput.EnableStackdriverLogging != "" {
		enableStackdriverLogging, err := utils.ParseBool(createInput.EnableStackdriverLogging)
		if err != nil {
			return nil, err
		}
		advancedSettings.GetLoggingSettings().EnableStackdriverLogging = enableStackdriverLogging
	}

	if createInput.EnableInteractionLogging != "" {
		enableInteractionLogging, err := utils.ParseBool(createInput.EnableInteractionLogging)
		if err != nil {
			return nil, err
		}
		advancedSettings.GetLoggingSettings().EnableInteractionLogging = enableInteractionLogging
	}

	speechToTextSettings := &cxpb.SpeechToTextSettings{}
	if createInput.EnableSpeechAdaptation != "" {
		enableSpeechAdaptation, err := utils.ParseBool(createInput.EnableSpeechAdaptation)
		if err != nil {
			return nil, err
		}
		speechToTextSettings.EnableSpeechAdaptation = enableSpeechAdaptation
	}

	agent := &cxpb.Agent{
		DisplayName:            agentName,
		DefaultLanguageCode:    createInput.DefaultLanguageCode,
		TimeZone:               createInput.TimeZone,
		Description:            createInput.Description,
		SupportedLanguageCodes: createInput.SupportedLanguageCodes,
		AvatarUri:              createInput.AvatarURI,
		AdvancedSettings:       advancedSettings,
		SpeechToTextSettings:   speechToTextSettings,
	}

	if createInput.EnableSpellCorrection != "" {
		enableSpellCorrection, err := utils.ParseBool(createInput.EnableSpellCorrection)
		if err != nil {
			return nil, err
		}
		agent.EnableSpellCorrection = enableSpellCorrection
	}

	reqCreateAgent := &cxpb.CreateAgentRequest{
		Parent: parentPath,
		Agent:  agent,
	}
	return agentClient.CreateAgent(ctx, reqCreateAgent)
}

func UpdateAgent(agentClient *cx.AgentsClient, agentName, locationID, projectID string, updateInput *types.CreateUpdateAgent) (*cxpb.Agent, error) {
	ctx := context.Background()
	paths := []string{}

	agent, err := GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return nil, err
	}
	if updateInput.TimeZone != "" {
		agent.TimeZone = updateInput.TimeZone
		paths = append(paths, "time_zone")
	}

	if updateInput.Description != "" {
		agent.Description = updateInput.Description
		paths = append(paths, "description")
	}

	if len(updateInput.SupportedLanguageCodes) > 0 {
		agent.SupportedLanguageCodes = updateInput.SupportedLanguageCodes
		paths = append(paths, "supported_language_codes")
	}

	if updateInput.AvatarURI != "" {
		agent.AvatarUri = updateInput.AvatarURI
		paths = append(paths, "avatar_uri")
	}

	if updateInput.EnableStackdriverLogging != "" {
		enableStackdriverLogging, err := utils.ParseBool(updateInput.EnableStackdriverLogging)
		if err != nil {
			return nil, err
		}
		agent.AdvancedSettings.LoggingSettings.EnableStackdriverLogging = enableStackdriverLogging
		paths = append(paths, "advanced_settings")
	}

	if updateInput.EnableInteractionLogging != "" {
		enableInteractionLogging, err := utils.ParseBool(updateInput.EnableInteractionLogging)
		if err != nil {
			return nil, err
		}
		agent.AdvancedSettings.LoggingSettings.EnableInteractionLogging = enableInteractionLogging
		paths = append(paths, "advanced_settings")
	}

	if updateInput.EnableSpeechAdaptation != "" {
		enableSpeechAdaptation, err := utils.ParseBool(updateInput.EnableSpeechAdaptation)
		if err != nil {
			return nil, err
		}
		agent.SpeechToTextSettings.EnableSpeechAdaptation = enableSpeechAdaptation
		paths = append(paths, "speech_to_text_settings")
	}

	if updateInput.EnableSpellCorrection != "" {
		enableSpellCorrection, err := utils.ParseBool(updateInput.EnableSpellCorrection)
		if err != nil {
			return nil, err
		}
		agent.EnableSpellCorrection = enableSpellCorrection
		paths = append(paths, "enable_spell_correction")
	}

	reqUpdateAgent := &cxpb.UpdateAgentRequest{
		Agent: agent,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: paths,
		},
	}
	return agentClient.UpdateAgent(ctx, reqUpdateAgent)
}

func DeleteAgent(agentClient *cx.AgentsClient, agentName, projectId, locationId string) error {
	ctx := context.Background()

	agent, err := GetAgentIdByName(agentClient, agentName, projectId, locationId)
	if err != nil {
		return err
	}

	reqDeleteAgent := &cxpb.DeleteAgentRequest{
		Name: agent.GetName(),
	}
	return agentClient.DeleteAgent(ctx, reqDeleteAgent)
}
