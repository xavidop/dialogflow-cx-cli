package profileconversation

import (
	"fmt"
	"strings"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	"cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/google/uuid"
	"github.com/tmc/langchaingo/llms/vertexai"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	types "github.com/xavidop/dialogflow-cx-cli/internal/types/profileconversation"
	"github.com/xavidop/dialogflow-cx-cli/internal/types/profileconversation/configurations"
	"github.com/xavidop/dialogflow-cx-cli/internal/utils"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
	vertexaipkg "github.com/xavidop/dialogflow-cx-cli/pkg/vertexai"
)

func ExecuteSuite(suiteFile string) error {

	var errstrings []string

	suite, err := types.NewSuite(suiteFile)
	if err != nil {
		return err
	}

	agentClient, err := cxpkg.CreateAgentRESTClient(suite.LocationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, suite.AgentName, suite.ProjectID, suite.LocationID)
	if err != nil {
		return err
	}

	palmTextClient, err := vertexaipkg.CreatePalmClient(suite.ProjectID)
	if err != nil {
		return err
	}

	sessionsClient, err := cxpkg.CreateSessionRESTClient(suite.LocationID)
	if err != nil {
		return err
	}
	defer sessionsClient.Close()

	global.Log.Infof("Suite Information: %s", suite.AgentName)
	sessionId := uuid.NewString()

	for _, testInfo := range suite.Tests {
		testInfo.File = utils.GetRelativeFilePathFromParentFile(suiteFile, testInfo.File)
		global.Log.Infof("Test ID: %s", testInfo.ID)

		test, err := types.NewTest(testInfo.File)
		if err != nil {
			return err
		}

		for _, interaction := range test.Interactions {

			response, err := getResponse(sessionsClient, palmTextClient, agent, test, interaction, testInfo, sessionId)
			if err != nil {
				return err
			}

			responseText := ""

			for _, message := range response.GetQueryResult().GetResponseMessages() {
				if message.GetEndInteraction() != nil {
					return nil
				}

				for _, txtToShow := range message.GetText().GetText() {

					responseText += ". " + txtToShow
				}

			}

			validations := interaction.Agent.Validate

			for _, validation := range validations {
				global.Log.Printf("Validation %s: %s ", validation.Type, validation.Value)

				switch validation.Type {
				case "contains":
					if err := executeContains(validation, responseText); err != nil {
						errstrings = append(errstrings, err.Error())
						continue
					}
				case "regexp":
					if err := executeRegexp(validation, responseText); err != nil {
						errstrings = append(errstrings, err.Error())
						continue
					}
				case "equals":
					if err := executeEquals(validation, responseText); err != nil {
						errstrings = append(errstrings, err.Error())
						continue
					}
				case "similarity":
					if err := executeSimilarity(validation, responseText); err != nil {
						errstrings = append(errstrings, err.Error())
						continue
					}
				}

			}

		}

	}

	if len(errstrings) > 0 {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	} else {
		return nil
	}

}

func executeSimilarity(validation *types.Validate, agentResponse string) error {
	switch validation.Algorithm {
	case "hamming":
		if err := executeSimilarityHamming(validation, agentResponse); err != nil {
			return err
		}
	case "levenshtein":
		if err := executeSimilarityLevenshtein(validation, agentResponse); err != nil {
			return err
		}
	case "jaro":
		if err := executeSimilarityJaro(validation, agentResponse); err != nil {
			return err
		}
	case "jaro-winkler":
		if err := executeSimilarityJaroWinkler(validation, agentResponse); err != nil {
			return err
		}
	case "smith-waterman-gotoh":
		if err := executeSimilaritySmithWatermanGotoh(validation, agentResponse); err != nil {
			return err
		}
	case "sorensen-dice":
		if err := executeSimilaritySorensenDice(validation, agentResponse); err != nil {
			return err
		}
	case "jaccard":
		if err := executeSimilarityJaccard(validation, agentResponse); err != nil {
			return err
		}
	case "overlap-coefficient":
		if err := executeSimilarityOverlapCoefficient(validation, agentResponse); err != nil {
			return err
		}
	}

	return nil
}

func executeContains(validation *types.Validate, agentResponse string) error {
	if validation.ConfigurationContains == nil {
		validation.ConfigurationContains = &configurations.Contains{}
	}
	global.Log.Printf("Validation configuration: %v ", validation.ConfigurationContains)

	if validation.ConfigurationContains.CaseSensitive {
		if !strings.Contains(validation.Value, agentResponse) {
			return fmt.Errorf("text %s does not contains \"%s\"", validation.Value, agentResponse)
		}
	} else {
		if !strings.Contains(validation.Value, agentResponse) {
			return fmt.Errorf("text %s does not contains \"%s\"", validation.Value, agentResponse)
		}

	}
	return nil
}

func executeRegexp(validation *types.Validate, agentResponse string) error {
	if validation.ConfigurationRegexp == nil {
		validation.ConfigurationRegexp = &configurations.Regexp{}
	}
	global.Log.Printf("Validation configuration: %v ", validation.ConfigurationRegexp)

	return nil
}

func executeEquals(validation *types.Validate, agentResponse string) error {
	if validation.ConfigurationEquals == nil {
		validation.ConfigurationEquals = &configurations.Equals{}
	}
	global.Log.Printf("Validation configuration: %v ", validation.ConfigurationEquals)

	if validation.ConfigurationEquals.CaseSensitive {
		if validation.Value != agentResponse {
			return fmt.Errorf("text %s is not equal to \"%s\"", validation.Value, agentResponse)
		}
	} else {
		if validation.Value != agentResponse {
			return fmt.Errorf("text %s is not equal to \"%s\"", validation.Value, agentResponse)
		}

	}

	return nil
}

func executeSimilarityHamming(validation *types.Validate, agentResponse string) error {
	if validation.ConfigurationHamming == nil {
		validation.ConfigurationHamming = metrics.NewHamming()
	}
	global.Log.Printf("Validation configuration: %v", validation.ConfigurationHamming)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationHamming)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func executeSimilarityLevenshtein(validation *types.Validate, agentResponse string) error {
	if validation.ConfigurationLevenshtein == nil {
		validation.ConfigurationLevenshtein = metrics.NewLevenshtein()
	}
	global.Log.Printf("Validation configuration: %v", validation.ConfigurationLevenshtein)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationLevenshtein)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func executeSimilarityJaro(validation *types.Validate, agentResponse string) error {
	if validation.ConfigurationJaro == nil {
		validation.ConfigurationJaro = metrics.NewJaro()
	}
	global.Log.Printf("Validation configuration: %v", validation.ConfigurationJaro)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationJaro)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func executeSimilarityJaroWinkler(validation *types.Validate, agentResponse string) error {
	if validation.ConfigurationJaroWinkler == nil {
		validation.ConfigurationJaroWinkler = metrics.NewJaroWinkler()
	}
	global.Log.Printf("Validation configuration: %v", validation.ConfigurationJaroWinkler)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationJaroWinkler)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func executeSimilaritySmithWatermanGotoh(validation *types.Validate, agentResponse string) error {
	if validation.ConfigurationSmithWatermanGotoh == nil {
		validation.ConfigurationSmithWatermanGotoh = metrics.NewSmithWatermanGotoh()
	}
	global.Log.Printf("Validation configuration: %v", validation.ConfigurationSmithWatermanGotoh)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationSmithWatermanGotoh)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func executeSimilaritySorensenDice(validation *types.Validate, agentResponse string) error {
	if validation.ConfigurationSorensenDice == nil {
		validation.ConfigurationSorensenDice = metrics.NewSorensenDice()
	}
	global.Log.Printf("Validation configuration: %v", validation.ConfigurationSorensenDice)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationSorensenDice)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func executeSimilarityJaccard(validation *types.Validate, agentResponse string) error {
	if validation.ConfigurationJaccard == nil {
		validation.ConfigurationJaccard = metrics.NewJaccard()
	}
	global.Log.Printf("Validation configuration: %v", validation.ConfigurationJaccard)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationJaccard)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func executeSimilarityOverlapCoefficient(validation *types.Validate, agentResponse string) error {
	if validation.ConfigurationOverlapCoefficient == nil {
		validation.ConfigurationOverlapCoefficient = metrics.NewOverlapCoefficient()
	}
	global.Log.Printf("Validation configuration: %v", validation.ConfigurationOverlapCoefficient)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationOverlapCoefficient)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func getResponse(sessionsClient *cx.SessionsClient, palmTextClient *vertexai.LLM, agent *cxpb.Agent, test *types.Test, interaction *types.Interaction, testInfo *types.Tests, sessionId string) (*cxpb.DetectIntentResponse, error) {
	switch interaction.User.Type {
	case "prompt":
		completion, err := vertexaipkg.GenerateText(palmTextClient, interaction.User.Prompt)
		if err != nil {
			return nil, err
		}

		global.Log.Infof("User: type: %s, value: %s \n", interaction.User.Type, interaction.User.Prompt)
		global.Log.Infof("Information get: %s \n", completion)

		global.Log.Infof("prompt generated: %s \n", interaction.User.Text)

		return cxpkg.DetectIntentFromText(sessionsClient, agent, test.LocaleID, interaction.User.Text, sessionId)
	case "text":
		global.Log.Infof("User: type: %s, value: %s \n", interaction.User.Type, interaction.User.Text)

		return cxpkg.DetectIntentFromText(sessionsClient, agent, test.LocaleID, interaction.User.Text, sessionId)
	case "audio":
		global.Log.Infof("User: type: %s, value: %s \n", interaction.User.Type, interaction.User.Audio)

		audioFile := utils.GetRelativeFilePathFromParentFile(testInfo.File, interaction.User.Audio)
		return cxpkg.DetectIntentFromAudio(sessionsClient, agent, test.LocaleID, audioFile, sessionId)
	}

	return nil, fmt.Errorf("option not %s accepted", interaction.User.Type)

}
