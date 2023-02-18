package cx

import (
	"context"
	"errors"
	"fmt"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"google.golang.org/api/option"
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
