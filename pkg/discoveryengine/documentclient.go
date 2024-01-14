package discoveryengine

import (
	"context"
	"fmt"

	discoveryengine "cloud.google.com/go/discoveryengine/apiv1alpha"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"google.golang.org/api/option"
)

func CreateDocumentRESTClient(locationId string) (*discoveryengine.DocumentClient, error) {
	ctx := context.Background()
	endpointString := "discoveryengine.googleapis.com:443"

	if locationId != "global" {
		endpointString = fmt.Sprintf("%s-discoveryengine.googleapis.com:443", locationId)

	}
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return discoveryengine.NewDocumentRESTClient(ctx, credentials, endpoint)
	} else {
		return discoveryengine.NewDocumentRESTClient(ctx, endpoint)
	}

}

func CreateDocumentGRPCClient(locationId string) (*discoveryengine.DocumentClient, error) {
	ctx := context.Background()

	endpointString := "discoveryengine.googleapis.com:443"

	if locationId != "global" {
		endpointString = fmt.Sprintf("%s-discoveryengine.googleapis.com:443", locationId)
	}

	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return discoveryengine.NewDocumentClient(ctx, credentials, endpoint)
	} else {
		return discoveryengine.NewDocumentClient(ctx, endpoint)
	}
}

func GetDocumentIdByName(client *discoveryengine.DocumentClient) {

}
