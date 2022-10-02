package cx

import (
	"context"
	"fmt"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	"github.com/xavidop/dialogflow-cx-test-runner/internal/global"
	"google.golang.org/api/option"
)

func CreateEnvironmentsClient(locationId string) (*cx.EnvironmentsClient, error) {
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
