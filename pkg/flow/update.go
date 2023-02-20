package flow

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Update(flowName, description, locationID, projectID, agentName, localeId, nluClassificationThreshold, nluModelType, nluModelTrainingMode string) error {

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

	_, err = cxpkg.UpdateFlow(flowClient, agent, flowName, description, localeId, nluClassificationThreshold, nluModelType, nluModelTrainingMode)
	if err != nil {
		return err
	}

	global.Log.Infof("Flow updated\n")

	return nil
}
