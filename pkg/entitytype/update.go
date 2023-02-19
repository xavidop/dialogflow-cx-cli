package entitytype

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Update(entityTypeName, locationID, projectID, agentName, localeId string, entities []string, redacted bool) error {

	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	entityTypeClient, err := cxpkg.CreateEntityTypeGRPCClient(locationID)
	if err != nil {
		return err
	}
	defer entityTypeClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	_, err = cxpkg.UpdateEntityType(entityTypeClient, agent, entityTypeName, localeId, entities, redacted)
	if err != nil {
		return err
	}

	global.Log.Infof("Entity Type updated")

	return nil
}
