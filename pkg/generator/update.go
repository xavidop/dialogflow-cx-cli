package generator

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Update(name, locationID, projectID, agentName, prompt, localeID string) error {
	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	generatorClient, err := cxpkg.CreateGeneratorGRPCClient(locationID)
	if err != nil {
		return err
	}
	defer generatorClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	_, err = cxpkg.UpdateGenerator(generatorClient, agent, name, prompt, localeID)
	if err != nil {
		return err
	}
	global.Log.Infof("Generator updated")

	return nil
}
