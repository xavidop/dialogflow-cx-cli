package datastore

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	discoveryenginepkg "github.com/xavidop/dialogflow-cx-cli/pkg/discoveryengine"
)

func Delete(name, locationID, projectID string) error {

	dataStoreClient, err := discoveryenginepkg.CreateDataStoreRESTClient(locationID)
	if err != nil {
		return err
	}
	defer dataStoreClient.Close()

	datastore, err := discoveryenginepkg.GetDataStoreIdByName(dataStoreClient, name, projectID, locationID)
	if err != nil {
		return err
	}

	if err := discoveryenginepkg.DeleteDataStore(dataStoreClient, datastore.GetName(), projectID, locationID); err != nil {
		return err
	}

	global.Log.Infof("Datastore deleted")

	return nil
}
