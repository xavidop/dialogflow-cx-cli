package agent

import (
	"github.com/xavidop/dialogflow-cx-cli/cmd/agent/types"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Create(agentName, locationID, projectID string, createInput *types.CreateUpdateAgent) error {

	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	agent, err := cxpkg.CreateAgent(agentClient, agentName, locationID, projectID, createInput)
	if err != nil {
		return err
	}

	global.Log.Infof("Agent created with id: %v\n", agent.GetName())

	return nil
}
