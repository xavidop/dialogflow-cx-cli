package versioning

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Create(name, description, startFlowName, locationID, projectID, agentName string) error {
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

	version, err := cxpkg.CreateVersion(versionClient, flow, description, name)
	if err != nil {
		return err
	}

	global.Log.Infof("Version created with id: %v\n", version.GetName())

	return nil
}
