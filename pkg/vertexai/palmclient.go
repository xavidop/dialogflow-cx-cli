package vertexai

import (
	"context"

	"github.com/tmc/langchaingo/llms/vertexai"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
)

func CreatePalmClient(projectId string) (*vertexai.LLM, error) {

	if global.Credentials != "" {
		credentials := vertexai.WithCredentialsFile(global.Credentials)
		project := vertexai.WithProjectID(projectId)
		return vertexai.New(credentials, project)
	} else {
		return vertexai.New()
	}

}

func GenerateText(palmClient *vertexai.LLM, prompt string) (string, error) {
	ctx := context.Background()

	return palmClient.Call(ctx, prompt)
}
