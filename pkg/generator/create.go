package generator

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Create(name, locationID, projectID, agentName, prompt, localeID string) error {
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

	generator, err := cxpkg.CreateGenerator(generatorClient, agent, name, prompt, localeID)
	if err != nil {
		return err
	}

	global.Log.Infof("Generator created with id: %v\n", generator.GetName())

	return nil
}
