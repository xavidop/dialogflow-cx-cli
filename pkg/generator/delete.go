package generator

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Delete(name, locationID, projectID, agentName, force string) error {
	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	generatorClient, err := cxpkg.CreateGeneratorRESTClient(locationID)
	if err != nil {
		return err
	}
	defer generatorClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	if err := cxpkg.DeleteGenerator(generatorClient, agent, name, force); err != nil {
		return err
	}

	global.Log.Infof("Generator deleted")

	return nil
}
