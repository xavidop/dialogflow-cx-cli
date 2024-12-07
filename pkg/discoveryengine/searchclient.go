package discoveryengine

import (
	"context"
	"fmt"
	"strings"

	discoveryengine "cloud.google.com/go/discoveryengine/apiv1alpha"
	discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1alpha/discoveryenginepb"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"google.golang.org/api/option"
)

func CreateSearchRESTClient(locationId string) (*discoveryengine.SearchClient, error) {
	ctx := context.Background()
	endpointString := "discoveryengine.googleapis.com:443"

	if locationId != "global" {
		endpointString = fmt.Sprintf("%s-discoveryengine.googleapis.com:443", locationId)

	}
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

	endpointString := "discoveryengine.googleapis.com:443"

	if locationId != "global" {
		endpointString = fmt.Sprintf("%s-discoveryengine.googleapis.com:443", locationId)
	}

	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return discoveryengine.NewSearchClient(ctx, credentials, endpoint)
	} else {
		return discoveryengine.NewSearchClient(ctx, endpoint)
	}
}

func Search(searchClient *discoveryengine.SearchClient, projectId string, locationId string, query string, datastore *discoveryenginepb.DataStore) error {

	ctx := context.Background()

	// Full resource name of search engine serving config
	servingConfig := fmt.Sprintf("%s/servingConfigs/default_search", datastore.GetName())

	searchRequest := &discoveryenginepb.SearchRequest{
		ServingConfig: servingConfig,
		Query:         query,
		PageSize:      10,
		ContentSearchSpec: &discoveryenginepb.SearchRequest_ContentSearchSpec{
			SnippetSpec: &discoveryenginepb.SearchRequest_ContentSearchSpec_SnippetSpec{
				ReturnSnippet: true,
			},
		},
	}

	results := searchClient.Search(ctx, searchRequest)

	for result, err := results.Next(); err == nil; {
		fmt.Printf("%+v\n", result)

		result, err = results.Next()
		if err != nil && !strings.Contains(err.Error(), "no more items") {
			return err
		}

	}

	return nil
}
