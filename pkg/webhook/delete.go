package webhook

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Delete(name, locationID, projectID, agentName, force string) error {
	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	webhookClient, err := cxpkg.CreateWebhookRESTClient(locationID)
	if err != nil {
		return err
	}
	defer webhookClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	if err := cxpkg.DeleteWebhook(webhookClient, agent, name, force); err != nil {
		return err
	}

	global.Log.Infof("Webhook deleted")

	return nil
}
