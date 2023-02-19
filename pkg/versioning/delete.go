package versioning

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Delete(name, startFlowName, locationID, projectID, agentName string) error {
	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	versionClient, err := cxpkg.CreateVersionRESTClient(locationID)
	if err != nil {
		return err
	}
	defer versionClient.Close()

	flowClient, err := cxpkg.CreateFlowRESTClient(locationID)
	if err != nil {
		return err
	}
	defer flowClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	flow, err := cxpkg.GetFlowIdByName(flowClient, agent, startFlowName)
	if err != nil {
		return err
	}

	if err := cxpkg.DeleteVersion(versionClient, flow, name); err != nil {
		return err
	}

	global.Log.Infof("Version deleted")

	return nil
}
