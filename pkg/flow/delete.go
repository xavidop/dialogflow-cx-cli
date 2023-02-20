package flow

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Delete(FlowtName, locationID, projectID, agentName string) error {

	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	flowClient, err := cxpkg.CreateFlowRESTClient(locationID)
	if err != nil {
		return err
	}
	defer flowClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	if err := cxpkg.DeleteFlow(flowClient, agent, FlowtName); err != nil {
		return err
	}

	global.Log.Infof("Flow deleted")

	return nil
}
