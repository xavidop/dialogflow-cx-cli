package environment

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Delete(name, locationID, projectID, agentName string) error {
	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	environmentClient, err := cxpkg.CreateEnvironmentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer environmentClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	if err := cxpkg.DeleteEnvironment(environmentClient, agent, name); err != nil {
		return err
	}

	global.Log.Infof("Environment deleted")

	return nil
}
