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

func CreateFlowRESTClient(locationId string) (*cx.FlowsClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewFlowsRESTClient(ctx, credentials, endpoint)
	} else {
		return cx.NewFlowsRESTClient(ctx, endpoint)
	}

}

func CreateFlowGRPCClient(locationId string) (*cx.FlowsClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewFlowsClient(ctx, credentials, endpoint)
	} else {
		return cx.NewFlowsClient(ctx, endpoint)
	}

}

func GetFlowIdByName(flowClient *cx.FlowsClient, agent *cxpb.Agent, flowName string) (*cxpb.Flow, error) {
	ctx := context.Background()

	reqList := &cxpb.ListFlowsRequest{
		Parent: agent.GetName(),
	}

	flows := flowClient.ListFlows(ctx, reqList)

	for flow, err := flows.Next(); err == nil; {
		if flow.DisplayName == flowName {
			return flow, nil
		}
		flow, err = flows.Next()
		if err != nil {
			return nil, err
		}

	}

	return nil, errors.New("flow not found")
}
