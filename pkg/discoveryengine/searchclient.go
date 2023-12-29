package discoveryengine

import (
	"context"
	"fmt"

	discoveryengine "cloud.google.com/go/discoveryengine/apiv1beta"
	discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func CreateSearchRESTClient(locationId string) (*discoveryengine.SearchClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-discoveryengine.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return discoveryengine.NewSearchRESTClient(ctx, credentials, endpoint)
	} else {
		return discoveryengine.NewSearchRESTClient(ctx, endpoint)
	}

}

func CreateSearchGRPCClient(locationId string) (*discoveryengine.SearchClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-discoveryengine.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return discoveryengine.NewSearchClient(ctx, credentials, endpoint)
	} else {
		return discoveryengine.NewSearchClient(ctx, endpoint)
	}
}

func Search(searchClient *discoveryengine.SearchClient, projectId string, locationId string, dataStoreId string, query string) (string, error) {

	ctx := context.Background()

	// Full resource name of search engine serving config
	servingConfig := fmt.Sprintf("projects/%s/locations/%s/collections/default_collection/dataStores/%s/servingConfigs/default_serving_config",
		projectId, locationId, dataStoreId)

	searchRequest := &discoveryenginepb.SearchRequest{
		ServingConfig: servingConfig,
		Query:         query,
	}

	it := searchClient.Search(ctx, searchRequest)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return "", err
		}
		fmt.Printf("%+v\n", resp)
	}

	return "", nil
}
