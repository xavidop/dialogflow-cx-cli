package environment

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Update(name, description, startFlowName, locationID, projectID, agentName string) error {
	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	versionClient, err := cxpkg.CreateVersionGRPCClient(locationID)
	if err != nil {
		return err
	}
	defer versionClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	flowClient, err := cxpkg.CreateFlowRESTClient(locationID)
	if err != nil {
		return err
	}
	defer flowClient.Close()

	flow, err := cxpkg.GetFlowIdByName(flowClient, agent, startFlowName)
	if err != nil {
		return err
	}

	_, err = cxpkg.UpdateVersion(versionClient, flow, description, name)
	if err != nil {
		return err
	}
	global.Log.Infof("Version updated")

	return nil
}
