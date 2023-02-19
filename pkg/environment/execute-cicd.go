package environment

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func ExecutePipeline(envName string, locationID string, projectID string, agentName string) error {

	if global.Output != "json" {
		global.Log.Infof("Executing cicd for environment %s", envName)

	}

	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	environmentClient, err := cxpkg.CreateEnvironmentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer environmentClient.Close()

	env, err := cxpkg.GetEnvironmentIdByName(environmentClient, agent.GetName(), envName)
	if err != nil {
		return err
	}

	result, err := cxpkg.RunContinuousTest(environmentClient, env)
	if err != nil {
		return err
	}

	global.Log.Infof(result.GetResult().String())

	return nil
}
