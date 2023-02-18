package cx

import (
	"context"
	"errors"
	"fmt"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func CreateEnvironmentRESTClient(locationId string) (*cx.EnvironmentsClient, error) {
	ctx := context.Background()

	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewEnvironmentsRESTClient(ctx, credentials, endpoint)
	} else {
		return cx.NewEnvironmentsRESTClient(ctx, endpoint)
	}

}

func CreateEnvironmentGRPCClient(locationId string) (*cx.EnvironmentsClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewEnvironmentsClient(ctx, credentials, endpoint)
	} else {
		return cx.NewEnvironmentsClient(ctx, endpoint)
	}

}

func GetEnvironmentIdByName(environmentClient *cx.EnvironmentsClient, agentID string, environmentName string) (*cxpb.Environment, error) {
	ctx := context.Background()

	reqList := &cxpb.ListEnvironmentsRequest{
		Parent: agentID,
	}

	envs := environmentClient.ListEnvironments(ctx, reqList)

	for env, err := envs.Next(); err == nil; {
		if env.DisplayName == environmentName {
			return env, nil
		}
		env, err = envs.Next()
		if err != nil {
			return nil, err
		}

	}

	return nil, errors.New("environment not found")
}

func UpdateWebhookConfig(environmentClient *cx.EnvironmentsClient, environment *cxpb.Environment, webhook *cxpb.Webhook) (string, error) {
	ctx := context.Background()

	if len(environment.GetWebhookConfig().GetWebhookOverrides()) == 0 {
		environment.GetWebhookConfig().WebhookOverrides = []*cxpb.Webhook{webhook}
	} else {
		for i, webhookOverrides := range environment.GetWebhookConfig().GetWebhookOverrides() {
			if webhookOverrides.Name == webhook.Name {
				environment.GetWebhookConfig().GetWebhookOverrides()[i] = webhook
			}
		}
	}

	reqUpdateEnvironment := &cxpb.UpdateEnvironmentRequest{
		Environment: environment,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"webhook_config"},
		},
	}

	op, err := environmentClient.UpdateEnvironment(ctx, reqUpdateEnvironment)
	if err != nil {
		return "", err
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		return "", err
	}

	return resp.GetUpdateTime().String(), err
}

func RunContinuousTest(environmentClient *cx.EnvironmentsClient, env *cxpb.Environment) (*cxpb.ContinuousTestResult, error) {
	ctx := context.Background()

	req := &cxpb.RunContinuousTestRequest{
		Environment: env.GetName(),
	}
	op, err := environmentClient.RunContinuousTest(ctx, req)
	if err != nil {
		return nil, err
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		return nil, err
	}

	return resp.ContinuousTestResult, nil
}
