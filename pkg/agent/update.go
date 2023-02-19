package agent

import (
	"github.com/xavidop/dialogflow-cx-cli/cmd/agent/types"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Update(agentName, locationID, projectID string, updateInput *types.CreateUpdateAgent) error {

	agentClient, err := cxpkg.CreateAgentGRPCClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	_, err = cxpkg.UpdateAgent(agentClient, agentName, locationID, projectID, updateInput)
	if err != nil {
		return err
	}

	global.Log.Infof("Agent updated\n")

	return nil
}
