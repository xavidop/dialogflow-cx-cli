package datastore

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	discoveryenginepkg "github.com/xavidop/dialogflow-cx-cli/pkg/discoveryengine"
)

func Search(name, locationID, projectID, query string) error {

	dataStoreClient, err := discoveryenginepkg.CreateDataStoreRESTClient(locationID)
	if err != nil {
		return err
	}
	defer dataStoreClient.Close()

	searchClient, err := discoveryenginepkg.CreateSearchRESTClient(locationID)
	if err != nil {
		return err
	}
	defer searchClient.Close()

	datastore, err := discoveryenginepkg.GetDataStoreIdByName(dataStoreClient, name, projectID, locationID)
	if err != nil {
		return err
	}

	if err := discoveryenginepkg.Search(searchClient, projectID, locationID, query, datastore); err != nil {
		return err
	}

	global.Log.Infof("Datastore search finished")

	return nil
}
