package discoveryengine

import (
	"context"
	"errors"
	"fmt"
	"strings"

	discoveryengine "cloud.google.com/go/discoveryengine/apiv1alpha"
	discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1alpha/discoveryenginepb"
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

func GetDocumentIdByName(client *discoveryengine.DocumentClient, dataStore *discoveryenginepb.DataStore, documentName string) (*discoveryenginepb.Document, error) {
	ctx := context.Background()

	parentPath := fmt.Sprintf("%s/branches/default_branch", dataStore.GetName())
	reqDocumentList := &discoveryenginepb.ListDocumentsRequest{
		Parent: parentPath,
	}

	documents := client.ListDocuments(ctx, reqDocumentList)

	for document, err := documents.Next(); err == nil; {
		// GCS Documents and websites
		if document.GetStructData().String() == "" {
			if strings.Contains(document.GetContent().GetUri(), documentName) {
				return document, nil
			}
		} else {
			// JSON Documents
			if strings.Contains(document.GetStructData().String(), documentName) {
				return document, nil
			}
		}
		document, err = documents.Next()
		if err != nil {
			return nil, err
		}

	}

	return nil, errors.New("document not found")
}
