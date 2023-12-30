package datastore

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	discoveryenginepkg "github.com/xavidop/dialogflow-cx-cli/pkg/discoveryengine"
)

func Search(name, locationID, projectID, query string) error {
	searchClient, err := discoveryenginepkg.CreateSearchRESTClient(locationID)
	if err != nil {
		return err
	}
	defer searchClient.Close()

	if err := discoveryenginepkg.Search(searchClient, projectID, locationID, name, query); err != nil {
		return err
	}

	global.Log.Infof("Datastore search finished")

	return nil
}
