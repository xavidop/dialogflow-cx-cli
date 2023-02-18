package cx

import (
	"context"
	"fmt"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"google.golang.org/api/option"
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
