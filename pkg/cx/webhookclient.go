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

func CreateWebhook(webhookClient *cx.WebhooksClient, agent *cxpb.Agent, url, name, flexible, requestBody, parametersMapping string) (*cxpb.Webhook, error) {
	ctx := context.Background()
	webhookType := cxpb.Webhook_GenericWebService_STANDARD

	flexibleBool, err := utils.ParseBool(flexible)
	if err != nil {
		return nil, err
	}
	if flexibleBool {
		webhookType = cxpb.Webhook_GenericWebService_FLEXIBLE
	}

	parametersMappingMap, err := ParseParametersMapping(parametersMapping)
	if err != nil {
		return nil, err
	}

	webhook := &cxpb.Webhook{
		DisplayName: name,
		Webhook: &cxpb.Webhook_GenericWebService_{
			GenericWebService: &cxpb.Webhook_GenericWebService{
				Uri:              url,
				WebhookType:      webhookType,
				RequestBody:      requestBody,
				ParameterMapping: parametersMappingMap,
			},
		},
	}

	reqCreateWebhook := &cxpb.CreateWebhookRequest{
		Parent:  agent.GetName(),
		Webhook: webhook,
	}

	return webhookClient.CreateWebhook(ctx, reqCreateWebhook)
}

func ParseParametersMapping(parametersMapping string) (map[string]string, error) {
	parametersMappingMap := make(map[string]string)
	if parametersMapping != "" {
		mappings := strings.Split(parametersMapping, ",")
		for _, mapping := range mappings {
			mappingSplit := strings.Split(mapping, "@")
			if len(mappingSplit) != 2 {
				return nil, errors.New("invalid parameters mapping")
			}
			parametersMappingMap[strings.TrimSpace(mappingSplit[0])] = strings.TrimSpace(mappingSplit[1])
		}
		return parametersMappingMap, nil
	}
	return parametersMappingMap, nil
}

func UpdateWebhook(webhookClient *cx.WebhooksClient, agent *cxpb.Agent, url, name, flexible, requestBody, parametersMapping string) (*cxpb.Webhook, error) {
	ctx := context.Background()

	webhook, err := GetWebhookIdByName(webhookClient, agent, name)
	if err != nil {
		return nil, err
	}

	if flexible != "" {
		flexibelBool, err := utils.ParseBool(flexible)
		if err != nil {
			return nil, err
		}
		if flexibelBool {
			webhook.GetGenericWebService().WebhookType = cxpb.Webhook_GenericWebService_FLEXIBLE
		} else {
			webhook.GetGenericWebService().WebhookType = cxpb.Webhook_GenericWebService_STANDARD
		}
	}

	if requestBody != "" {
		webhook.GetGenericWebService().RequestBody = requestBody
	}

	if url != "" {
		webhook.GetGenericWebService().Uri = url
	}
	if parametersMapping != "" {
		parametersMappingMap, err := ParseParametersMapping(parametersMapping)
		if err != nil {
			return nil, err
		}
		webhook.GetGenericWebService().ParameterMapping = parametersMappingMap
	}

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

func DeleteWebhook(webhookClient *cx.WebhooksClient, agent *cxpb.Agent, name, force string) error {
	ctx := context.Background()

	webhook, err := GetWebhookIdByName(webhookClient, agent, name)
	if err != nil {
		return err
	}

	reqDeleteWebhook := &cxpb.DeleteWebhookRequest{
		Name: webhook.GetName(),
	}

	if force != "" {
		forceBool, err := utils.ParseBool(force)
		if err != nil {
			return err
		}
		reqDeleteWebhook.Force = forceBool
	}

	return webhookClient.DeleteWebhook(ctx, reqDeleteWebhook)
}
