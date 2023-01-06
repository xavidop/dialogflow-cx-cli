package agent

import (
	"os"

	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Export(locationID, projectID, agentName, output string) error {

	agentClient, err := cxpkg.CreateAgentGRPCClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	responseData, err := cxpkg.ExportAgentById(agentClient, agent.GetName())
	if err != nil {
		return err
	}

	err = os.WriteFile(output, responseData.GetAgentContent(), 0644)
	if err != nil {
		return err
	}
	global.Log.Infof("Agent exported to file: %v\n", output)

	return nil
}
