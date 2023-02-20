package flow

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Train(flowName, locationID, projectID, agentName string) error {

	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	flowClient, err := cxpkg.CreateFlowGRPCClient(locationID)
	if err != nil {
		return err
	}
	defer flowClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	if flowName == "all" {
		if err := cxpkg.TrainAll(flowClient, agent); err != nil {
			return err
		}

		global.Log.Infof("All flows trained")

	} else {
		if err := cxpkg.Train(flowClient, agent, flowName, ""); err != nil {
			return err
		}
		global.Log.Infof("Flow trained")

	}

	return nil
}
