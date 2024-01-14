package discoveryengine

import (
	"context"
	"errors"
	"fmt"

	discoveryengine "cloud.google.com/go/discoveryengine/apiv1alpha"
	discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1alpha/discoveryenginepb"
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

func GetDataStoreIdByName(client *discoveryengine.DataStoreClient, dataStoreName string, projectId string, locationId string) (*discoveryenginepb.DataStore, error) {

	ctx := context.Background()

	parentPath := fmt.Sprintf("projects/%s/locations/%s/collections/default_collection", projectId, locationId)

	reqAgentList := &discoveryenginepb.ListDataStoresRequest{
		Parent: parentPath,
	}

	datastores := client.ListDataStores(ctx, reqAgentList)

	for datastore, err := datastores.Next(); err == nil; {
		if datastore.DisplayName == dataStoreName {
			return datastore, nil
		}
		datastore, err = datastores.Next()
		if err != nil {
			return nil, err
		}

	}

	return nil, errors.New("datastore not found")
}

func DeleteDataStore(dataStoreClient *discoveryengine.DataStoreClient, dataStoreId, projectId, locationId string) error {
	ctx := context.Background()

	reqDeleteDataStore := &discoveryenginepb.DeleteDataStoreRequest{
		Name: dataStoreId,
	}
	longRunningOperation, err := dataStoreClient.DeleteDataStore(ctx, reqDeleteDataStore)
	if err != nil {
		return err
	}

	return longRunningOperation.Wait(ctx)
}
