package intent

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Update(intentName, description, locationID, projectID, agentName, localeId string, trainingPhrases []string) error {

	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	intentClient, err := cxpkg.CreateIntentGRPCClient(locationID)
	if err != nil {
		return err
	}
	defer intentClient.Close()

	entityTypeClient, err := cxpkg.CreateEntityTypeRESTClient(locationID)
	if err != nil {
		return err
	}
	defer entityTypeClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	_, err = cxpkg.UpdateIntent(intentClient, agent, intentName, description, localeId, trainingPhrases, entityTypeClient)
	if err != nil {
		return err
	}

	global.Log.Infof("Intent updated\n")

	return nil
}
