package profilenlu

import (
	"fmt"
	"strconv"
	"strings"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	"cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/tmc/langchaingo/llms/vertexai"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	types "github.com/xavidop/dialogflow-cx-cli/internal/types/profilenlu"
	"github.com/xavidop/dialogflow-cx-cli/internal/utils"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
	vertexaipkg "github.com/xavidop/dialogflow-cx-cli/pkg/vertexai"
	"google.golang.org/protobuf/types/known/structpb"
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

	palmTextClient, err := vertexaipkg.CreatePalmClient(suite.ProjectID)
	if err != nil {
		return err
	}

	agent, err := cxpkg.GetAgentIdByName(agentClient, suite.AgentName, suite.ProjectID, suite.LocationID)
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
		testFileLog := global.Log.WithField("test-file", testInfo.ID)

		test, err := types.NewTest(testInfo.File)
		if err != nil {
			return err
		}

		for _, check := range test.Checks {

			checkLog := testFileLog.WithField("check", check.ID)

			checkTypeLog := checkLog.WithField("input", check.Input.Type)

			response, err := getResponse(sessionsClient, palmTextClient, agent, test, check, testInfo, sessionId, checkTypeLog)
			if err != nil {
				return err
			}

			queryResult := response.GetQueryResult()
			intentDetected := queryResult.GetMatch().GetIntent().GetDisplayName()
			parametersDetected := queryResult.Parameters.GetFields()
			responseText := ""

			for _, message := range queryResult.GetResponseMessages() {
				if message.GetEndInteraction() != nil {
					return nil
				}

				responseText += strings.Join(message.GetText().GetText(), ". ")

			}

			checkTypeLog.Infof("Agent> %s", responseText)

			validationLog := checkLog.WithField("validation", check.Validate.Intent)

			validationLog.Infof("Intent Detected: %s", intentDetected)

			if check.Validate.Intent != intentDetected {
				intentError := fmt.Errorf("intent \"%s\" does not match with the intent detected \"%s\"", check.Validate.Intent, intentDetected)
				validationLog.Errorf("%s", intentError.Error())
				errstrings = append(errstrings, intentError.Error())
				continue
			}

			parameters := check.Validate.Parameters

			for paramName, p := range parametersDetected {
				extractedValue := extractDialogflowEntities(p)
				validationLog.Infof("Param %s: %s ", paramName, extractedValue)
				param, err := types.FindParameterByName(*check.Validate, paramName)
				if err != nil {
					errstrings = append(errstrings, err.Error())
					continue
				}
				if param.Value != extractedValue {
					parameterError := fmt.Errorf("parameter value \"%s\" does not match with the parameter detected \"%s\"", param.Value, extractedValue)
					validationLog.Errorf("%s", parameterError.Error())
					errstrings = append(errstrings, parameterError.Error())
				}
				parameters = types.RemoveParameterByName(parameters, param.Parameter)

			}

			if len(parameters) > 0 {
				parametersNotDetectedError := fmt.Errorf("parameters not detected: %v", parameters)
				validationLog.Errorf("%s", parametersNotDetectedError.Error())
				errstrings = append(errstrings, parametersNotDetectedError.Error())
			}
		}

	}

	if len(errstrings) > 0 {
		return fmt.Errorf("there are %d errors in total", len(errstrings))
	} else {
		return nil
	}

}

func getResponse(sessionsClient *cx.SessionsClient, palmTextClient *vertexai.LLM, agent *cxpb.Agent, test *types.Test, check *types.Check, testInfo *types.Tests, sessionId string, log *logrus.Entry) (*cxpb.DetectIntentResponse, error) {

	switch check.Input.Type {
	case "prompt":
		textGenerated, err := vertexaipkg.GenerateText(palmTextClient, check.Input.Prompt)
		if err != nil {
			return nil, err
		}
		textGenerated = strings.TrimSpace(textGenerated)

		log.Infof("User> %s (auto-generated from prompt: \"%s\")", textGenerated, check.Input.Prompt)

		return cxpkg.DetectIntentFromText(sessionsClient, agent, test.LocaleID, textGenerated, sessionId)
	case "text":
		log.Infof("User> %s", check.Input.Text)

		return cxpkg.DetectIntentFromText(sessionsClient, agent, test.LocaleID, check.Input.Text, sessionId)
	case "audio":
		log.Infof("User> %s", check.Input.Audio)

		audioFile := utils.GetRelativeFilePathFromParentFile(testInfo.File, check.Input.Audio)
		return cxpkg.DetectIntentFromAudio(sessionsClient, agent, test.LocaleID, audioFile, sessionId)
	}

	return nil, fmt.Errorf("option not \"%s\" accepted", check.Input.Type)

}

func extractDialogflowEntities(p *structpb.Value) (extractedEntity string) {
	kind := p.GetKind()
	switch kind.(type) {
	case *structpb.Value_StringValue:
		return p.GetStringValue()
	case *structpb.Value_NumberValue:
		return strconv.FormatFloat(p.GetNumberValue(), 'f', 0, 64)
	case *structpb.Value_BoolValue:
		return strconv.FormatBool(p.GetBoolValue())
	default:
		return ""
	}
}
