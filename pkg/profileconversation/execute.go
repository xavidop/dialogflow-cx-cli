package profileconversation

import (
	"fmt"
	"regexp"
	"strings"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	"cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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

	global.Log.Infof("Running suite: %s", suite.Name)
	sessionId := uuid.NewString()

	for _, testInfo := range suite.Tests {
		testInfo.File = utils.GetRelativeFilePathFromParentFile(suiteFile, testInfo.File)

		test, err := types.NewTest(testInfo.File)
		if err != nil {
			return err
		}

		testFileLog := global.Log.WithField("test-file", testInfo.ID)

		for _, interaction := range test.Interactions {
			interactionLog := testFileLog.WithField("interaction", interaction.ID)

			interactionTypeLog := interactionLog.WithField("input", interaction.User.Type)

			response, err := getResponse(sessionsClient, palmTextClient, agent, test, interaction, testInfo, sessionId, interactionTypeLog)
			if err != nil {
				return err
			}

			responseText := ""

			for _, message := range response.GetQueryResult().GetResponseMessages() {
				if message.GetEndInteraction() != nil {
					return nil
				}

				responseText += strings.Join(message.GetText().GetText(), ". ")

			}

			interactionTypeLog.Infof("Agent> %s", responseText)

			validations := interaction.Agent.Validate

			for _, validation := range validations {
				validationLog := interactionLog.WithField("validation", validation.Type)

				validationLog.Printf("Validation with value \"%s\" ", validation.Value)

				switch validation.Type {
				case "contains":
					if err := executeContains(validation, responseText, validationLog); err != nil {
						validationLog.Errorf(err.Error())
						errstrings = append(errstrings, err.Error())
						continue
					}
				case "regexp":
					if err := executeRegexp(validation, responseText, validationLog); err != nil {
						validationLog.Errorf(err.Error())
						errstrings = append(errstrings, err.Error())
						continue
					}
				case "equals":
					if err := executeEquals(validation, responseText, validationLog); err != nil {
						validationLog.Errorf(err.Error())
						errstrings = append(errstrings, err.Error())
						continue
					}
				case "similarity":
					if err := executeSimilarity(validation, responseText, validationLog); err != nil {
						validationLog.Errorf(err.Error())
						errstrings = append(errstrings, err.Error())
						continue
					}
				}

			}

		}

	}

	if len(errstrings) > 0 {
		return fmt.Errorf("there are %d errors in total", len(errstrings))
	} else {
		return nil
	}

}

func executeSimilarity(validation *types.Validate, agentResponse string, log *logrus.Entry) error {
	switch validation.Algorithm {
	case "hamming":
		if err := executeSimilarityHamming(validation, agentResponse, log); err != nil {
			return err
		}
	case "levenshtein":
		if err := executeSimilarityLevenshtein(validation, agentResponse, log); err != nil {
			return err
		}
	case "jaro":
		if err := executeSimilarityJaro(validation, agentResponse, log); err != nil {
			return err
		}
	case "jaro-winkler":
		if err := executeSimilarityJaroWinkler(validation, agentResponse, log); err != nil {
			return err
		}
	case "smith-waterman-gotoh":
		if err := executeSimilaritySmithWatermanGotoh(validation, agentResponse, log); err != nil {
			return err
		}
	case "sorensen-dice":
		if err := executeSimilaritySorensenDice(validation, agentResponse, log); err != nil {
			return err
		}
	case "jaccard":
		if err := executeSimilarityJaccard(validation, agentResponse, log); err != nil {
			return err
		}
	case "overlap-coefficient":
		if err := executeSimilarityOverlapCoefficient(validation, agentResponse, log); err != nil {
			return err
		}
	}

	return nil
}

func executeContains(validation *types.Validate, agentResponse string, log *logrus.Entry) error {
	if validation.ConfigurationContains == nil {
		validation.ConfigurationContains = &configurations.Contains{}
	}
	log.Infof("Validation configuration: %+v", *validation.ConfigurationContains)

	if validation.ConfigurationContains.CaseSensitive {
		if !strings.Contains(validation.Value, agentResponse) {
			return fmt.Errorf("text %s does not contains \"%s\"", validation.Value, agentResponse)
		}
	} else {
		if !strings.Contains(strings.ToLower(validation.Value), strings.ToLower(agentResponse)) {
			return fmt.Errorf("text %s does not contains \"%s\"", validation.Value, agentResponse)
		}

	}
	return nil
}

func executeRegexp(validation *types.Validate, agentResponse string, log *logrus.Entry) error {
	if validation.ConfigurationRegexp == nil {
		validation.ConfigurationRegexp = &configurations.Regexp{}
	}

	log.Infof("Validation configuration: %+v", *validation.ConfigurationRegexp)

	r, err := regexp.Compile(validation.Value)
	if err != nil {
		return err
	}
	if validation.ConfigurationRegexp.FindInSubmatches {
		match := r.FindAllStringSubmatch(agentResponse, -1)
		if match == nil {
			return fmt.Errorf("regex %s does not match text \"%s\" in submatches", validation.Value, agentResponse)
		}
	} else {
		match := r.FindAllString(agentResponse, -1)
		if match == nil {
			return fmt.Errorf("regex %s does not match text \"%s\"", validation.Value, agentResponse)
		}
	}

	return nil
}

func executeEquals(validation *types.Validate, agentResponse string, log *logrus.Entry) error {
	if validation.ConfigurationEquals == nil {
		validation.ConfigurationEquals = &configurations.Equals{}
	}
	log.Infof("Validation configuration: %+v", *validation.ConfigurationEquals)

	if validation.ConfigurationEquals.CaseSensitive {
		if validation.Value != agentResponse {
			return fmt.Errorf("text %s is not equal to \"%s\"", validation.Value, agentResponse)
		}
	} else {
		if !strings.EqualFold(validation.Value, agentResponse) {
			return fmt.Errorf("text %s is not equal to \"%s\"", validation.Value, agentResponse)
		}

	}

	return nil
}

func executeSimilarityHamming(validation *types.Validate, agentResponse string, log *logrus.Entry) error {
	if validation.ConfigurationHamming == nil {
		validation.ConfigurationHamming = metrics.NewHamming()
	}
	log.Infof("Validation configuration: %+v", *validation.ConfigurationHamming)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationHamming)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func executeSimilarityLevenshtein(validation *types.Validate, agentResponse string, log *logrus.Entry) error {
	if validation.ConfigurationLevenshtein == nil {
		validation.ConfigurationLevenshtein = metrics.NewLevenshtein()
	}
	log.Infof("Validation configuration: %+v", *validation.ConfigurationLevenshtein)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationLevenshtein)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func executeSimilarityJaro(validation *types.Validate, agentResponse string, log *logrus.Entry) error {
	if validation.ConfigurationJaro == nil {
		validation.ConfigurationJaro = metrics.NewJaro()
	}
	log.Infof("Validation configuration: %+v", *validation.ConfigurationJaro)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationJaro)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func executeSimilarityJaroWinkler(validation *types.Validate, agentResponse string, log *logrus.Entry) error {
	if validation.ConfigurationJaroWinkler == nil {
		validation.ConfigurationJaroWinkler = metrics.NewJaroWinkler()
	}
	log.Infof("Validation configuration: %+v", *validation.ConfigurationJaroWinkler)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationJaroWinkler)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func executeSimilaritySmithWatermanGotoh(validation *types.Validate, agentResponse string, log *logrus.Entry) error {
	if validation.ConfigurationSmithWatermanGotoh == nil {
		validation.ConfigurationSmithWatermanGotoh = metrics.NewSmithWatermanGotoh()
	}
	log.Infof("Validation configuration: %+v", *validation.ConfigurationSmithWatermanGotoh)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationSmithWatermanGotoh)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func executeSimilaritySorensenDice(validation *types.Validate, agentResponse string, log *logrus.Entry) error {
	if validation.ConfigurationSorensenDice == nil {
		validation.ConfigurationSorensenDice = metrics.NewSorensenDice()
	}
	log.Infof("Validation configuration: %+v", *validation.ConfigurationSorensenDice)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationSorensenDice)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func executeSimilarityJaccard(validation *types.Validate, agentResponse string, log *logrus.Entry) error {
	if validation.ConfigurationJaccard == nil {
		validation.ConfigurationJaccard = metrics.NewJaccard()
	}
	log.Infof("Validation configuration: %+v", *validation.ConfigurationJaccard)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationJaccard)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func executeSimilarityOverlapCoefficient(validation *types.Validate, agentResponse string, log *logrus.Entry) error {
	if validation.ConfigurationOverlapCoefficient == nil {
		validation.ConfigurationOverlapCoefficient = metrics.NewOverlapCoefficient()
	}
	log.Infof("Validation configuration: %+v", *validation.ConfigurationOverlapCoefficient)

	similarity := strutil.Similarity(validation.Value, agentResponse, validation.ConfigurationOverlapCoefficient)

	if similarity < validation.Threshold {
		return fmt.Errorf("text not similar. similarity expected: %f, similarity obtained: %f", validation.Threshold, similarity)
	}

	return nil
}

func getResponse(sessionsClient *cx.SessionsClient, palmTextClient *vertexai.LLM, agent *cxpb.Agent, test *types.Test, interaction *types.Interaction, testInfo *types.Tests, sessionId string, log *logrus.Entry) (*cxpb.DetectIntentResponse, error) {

	switch interaction.User.Type {
	case "prompt":
		textGenerated, err := vertexaipkg.GenerateText(palmTextClient, interaction.User.Prompt)
		if err != nil {
			return nil, err
		}
		textGenerated = strings.TrimSpace(textGenerated)

		log.Infof("User> %s (auto-generated from prompt: \"%s\")", textGenerated, interaction.User.Prompt)

		return cxpkg.DetectIntentFromText(sessionsClient, agent, test.LocaleID, textGenerated, sessionId)
	case "text":
		log.Infof("User> %s", interaction.User.Text)

		return cxpkg.DetectIntentFromText(sessionsClient, agent, test.LocaleID, interaction.User.Text, sessionId)
	case "audio":
		log.Infof("User> %s", interaction.User.Audio)

		audioFile := utils.GetRelativeFilePathFromParentFile(testInfo.File, interaction.User.Audio)
		return cxpkg.DetectIntentFromAudio(sessionsClient, agent, test.LocaleID, audioFile, sessionId)
	}

	return nil, fmt.Errorf("option not \"%s\" accepted", interaction.User.Type)

}
