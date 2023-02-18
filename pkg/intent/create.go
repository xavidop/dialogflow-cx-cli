package intent

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Create(intentName, locationID, projectID, agentName, localeId string, trainingPhrases []string) error {

	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	intentClient, err := cxpkg.CreateIntentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer intentClient.Close()

	entityTypeClient, err := cxpkg.CreateEntityTypesRESTClient(locationID)
	if err != nil {
		return err
	}
	defer entityTypeClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	intent, err := cxpkg.CreateIntent(intentClient, agent, intentName, localeId, trainingPhrases, entityTypeClient)
	if err != nil {
		return err
	}

	global.Log.Infof("Intent created with id: %v\n", intent.GetName())

	return nil
}
