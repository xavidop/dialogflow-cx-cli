package entitytype

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Delete(entityTypeName, locationID, projectID, agentName, force string) error {

	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	entityTypeClient, err := cxpkg.CreateEntityTypeRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	if err := cxpkg.DeleteEntityType(entityTypeClient, agent, entityTypeName, force); err != nil {
		return err
	}

	global.Log.Infof("Entity Type deleted")

	return nil
}
