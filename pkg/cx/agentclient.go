package cx

import (
	"context"
	"errors"
	"fmt"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	"github.com/xavidop/dialogflow-cx-test-runner/internal/global"
	"google.golang.org/api/option"
)

func CreateAgentClient(locationId string) (*cx.AgentsClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewAgentsRESTClient(ctx, credentials, endpoint)
	} else {
		return cx.NewAgentsRESTClient(ctx, endpoint)
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
	}

	return nil, errors.New("agent not found")
}
