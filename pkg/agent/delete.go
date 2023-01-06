package agent

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Delete(locationID, projectID, agentName string) error {

	agentClient, err := cxpkg.CreateAgentGRPCClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	err = cxpkg.DeleteAgent(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	global.Log.Infoln("Agent deleted")

	return nil
}
