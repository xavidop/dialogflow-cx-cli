package entitytype

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Create(entityTypeName, locationID, projectID, agentName, localeId string, entities []string, redacted string) error {

	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	entityTypeClient, err := cxpkg.CreateEntityTypeRESTClient(locationID)
	if err != nil {
		return err
	}
	defer entityTypeClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	entityType, err := cxpkg.CreateEntityType(entityTypeClient, agent, entityTypeName, localeId, entities, redacted)
	if err != nil {
		return err
	}

	global.Log.Infof("Entity Type created with id: %v\n", entityType.GetName())

	return nil
}
