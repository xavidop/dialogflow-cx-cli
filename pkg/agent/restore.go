package agent

import (
	"os"

	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Restore(locationID, projectID, agentName, input string) error {

	agentClient, err := cxpkg.CreateAgentGRPCClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	content, err := os.ReadFile(input)
	if err != nil {
		return err
	}

	err = cxpkg.RestoreAgentByFullName(agentClient, agent.GetName(), projectID, locationID, content)
	if err != nil {
		return err
	}

	global.Log.Infoln("Agent restored")

	return nil
}
