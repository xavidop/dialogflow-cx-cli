package vertexai

import (
	"context"

	"github.com/tmc/langchaingo/llms/googleai/palm"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
)

func CreatePalmClient(projectId string) (*palm.LLM, error) {

	if global.Credentials != "" {
		credentials := palm.WithCredentialsFile(global.Credentials)
		project := palm.WithProjectID(projectId)
		return palm.New(credentials, project)
	} else {
		return palm.New()
	}

}

func GenerateText(palmClient *palm.LLM, prompt string) (string, error) {
	ctx := context.Background()

	return palmClient.Call(ctx, prompt)
}
