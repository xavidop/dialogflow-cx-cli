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

func CreateWebhookRESTClient(locationId string) (*cx.WebhooksClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewWebhooksRESTClient(ctx, credentials, endpoint)
	} else {
		return cx.NewWebhooksRESTClient(ctx, endpoint)
	}

}

func CreateWebhookGRPCClient(locationId string) (*cx.WebhooksClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewWebhooksClient(ctx, credentials, endpoint)
	} else {
		return cx.NewWebhooksClient(ctx, endpoint)
	}

}

func CreateWebhook(webhookClient *cx.WebhooksClient, agent *cxpb.Agent, url, name string) (*cxpb.Webhook, error) {
	ctx := context.Background()

	webhook := &cxpb.Webhook{
		DisplayName: name,
		Webhook: &cxpb.Webhook_GenericWebService_{
			GenericWebService: &cxpb.Webhook_GenericWebService{
				Uri: url,
			},
		},
	}

	reqCreateWebhook := &cxpb.CreateWebhookRequest{
		Parent:  agent.GetName(),
		Webhook: webhook,
	}

	return webhookClient.CreateWebhook(ctx, reqCreateWebhook)
}

func UpdateWebhook(webhookClient *cx.WebhooksClient, agent *cxpb.Agent, url, name string) (*cxpb.Webhook, error) {
	ctx := context.Background()

	webhook, err := GetWebhookIdByName(webhookClient, agent, name)
	if err != nil {
		return nil, err
	}

	webhook.GetGenericWebService().Uri = url

	reqUpdateWebhook := &cxpb.UpdateWebhookRequest{
		Webhook: webhook,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"generic_web_service"},
		},
	}

	return webhookClient.UpdateWebhook(ctx, reqUpdateWebhook)
}

func GetWebhookIdByName(webhookClient *cx.WebhooksClient, agent *cxpb.Agent, name string) (*cxpb.Webhook, error) {
	ctx := context.Background()

	reqWebhookList := &cxpb.ListWebhooksRequest{
		Parent: agent.GetName(),
	}

	webhooks := webhookClient.ListWebhooks(ctx, reqWebhookList)

	for webhook, err := webhooks.Next(); err == nil; {
		if webhook.DisplayName == name {
			return webhook, nil
		}
		webhook, err = webhooks.Next()
		if err != nil {
			return nil, err
		}

	}

	return nil, errors.New("webhook not found")

}

func DeleteWebhook(webhookClient *cx.WebhooksClient, agent *cxpb.Agent, name string, force bool) error {
	ctx := context.Background()

	webhook, err := GetWebhookIdByName(webhookClient, agent, name)
	if err != nil {
		return err
	}

	reqDeleteWebhook := &cxpb.DeleteWebhookRequest{
		Name:  webhook.GetName(),
		Force: force,
	}

	return webhookClient.DeleteWebhook(ctx, reqDeleteWebhook)
}
