package webhook

import (
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
)

func Create(name, url, locationID, projectID, agentName, environment string) error {
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

	webhook, err := cxpkg.CreateWebhook(webhookClient, agent, url, name)
	if err != nil {
		return err
	}

	if environment != "global" {
		environmentClient, err := cxpkg.CreateEnvironmentGRPCClient(locationID)
		if err != nil {
			return err
		}
		defer environmentClient.Close()

		env, err := cxpkg.GetEnvironmentIdByName(environmentClient, agent.GetName(), environment)
		if err != nil {
			return err
		}

		if _, err := cxpkg.UpdateWebhookConfig(environmentClient, env, webhook); err != nil {
			return err
		}

	}
	global.Log.Infof("Webhook created with id: %v\n", webhook.GetName())

	return nil
}
