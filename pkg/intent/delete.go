package intent

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Delete(intentName, locationID, projectID, agentName string) error {

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

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	if err := cxpkg.DeleteIntent(intentClient, agent, intentName); err != nil {
		return err
	}

	global.Log.Infof("Intent deleted")

	return nil
}
