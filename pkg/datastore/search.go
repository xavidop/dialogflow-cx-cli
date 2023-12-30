package datastore

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	discoveryenginepkg "github.com/xavidop/dialogflow-cx-cli/pkg/discoveryengine"
)

func Search(name, locationID, projectID, query string) error {
	searchClient, err := discoveryenginepkg.CreateSearchGRPCClient(locationID)
	if err != nil {
		return err
	}
	defer searchClient.Close()

	discoveryenginepkg.Search(searchClient, projectID, locationID, name, query)

	global.Log.Infof("Webhook created with id: %v\n", "")

	return nil
}
