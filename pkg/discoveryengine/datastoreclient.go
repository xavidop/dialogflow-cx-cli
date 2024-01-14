package discoveryengine

import (
	"context"
	"fmt"

	discoveryengine "cloud.google.com/go/discoveryengine/apiv1alpha"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"google.golang.org/api/option"
)

func CreateDataStoreRESTClient(locationId string) (*discoveryengine.DataStoreClient, error) {
	ctx := context.Background()
	endpointString := "discoveryengine.googleapis.com:443"

	if locationId != "global" {
		endpointString = fmt.Sprintf("%s-discoveryengine.googleapis.com:443", locationId)

	}
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return discoveryengine.NewDataStoreRESTClient(ctx, credentials, endpoint)
	} else {
		return discoveryengine.NewDataStoreRESTClient(ctx, endpoint)
	}

}

func CreateDataStoreGRPCClient(locationId string) (*discoveryengine.DataStoreClient, error) {
	ctx := context.Background()

	endpointString := "discoveryengine.googleapis.com:443"

	if locationId != "global" {
		endpointString = fmt.Sprintf("%s-discoveryengine.googleapis.com:443", locationId)
	}

	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return discoveryengine.NewDataStoreClient(ctx, credentials, endpoint)
	} else {
		return discoveryengine.NewDataStoreClient(ctx, endpoint)
	}
}
