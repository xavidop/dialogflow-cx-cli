package profilenlu

import (
	"fmt"
	"strconv"
	"strings"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	"cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	"github.com/google/uuid"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/internal/types"
	"github.com/xavidop/dialogflow-cx-cli/internal/utils"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
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

	agent, err := cxpkg.GetAgentIdByName(agentClient, suite.AgentName, suite.ProjectID, suite.LocationID)
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

		for _, check := range test.Checks {

			response, err := getResponse(sessionsClient, agent, test, check, testInfo, sessionId)
			if err != nil {
				return err
			}

			queryResult := response.GetQueryResult()
			intentDetected := queryResult.GetMatch().GetIntent().GetDisplayName()
			parametersDetected := queryResult.Parameters.GetFields()
			global.Log.Infof("Intent Detected: %s\n", intentDetected)
			if check.Validate.Intent != intentDetected {
				intentError := fmt.Errorf("intent %s does not match with the intent detected %s", check.Validate.Intent, intentDetected)

				errstrings = append(errstrings, intentError.Error())
				continue
			}

			parameters := check.Validate.Parameters

			for paramName, p := range parametersDetected {
				extractedValue := extractDialogflowEntities(p)
				global.Log.Printf("Param %s: %s ", paramName, extractedValue)
				param, err := types.FindParameterByName(check.Validate, paramName)
				if err != nil {
					errstrings = append(errstrings, err.Error())
					continue
				}
				if param.Value != extractedValue {
					parameterError := fmt.Errorf("parameter value %s does not match with the parameter detected %s", param.Value, extractedValue)
					errstrings = append(errstrings, parameterError.Error())
				}
				parameters = types.RemoveParameterByName(parameters, param.Parameter)

			}

			if len(parameters) > 0 {
				parametersNotDetectedError := fmt.Errorf("parameters not detected: %v", parameters)
				errstrings = append(errstrings, parametersNotDetectedError.Error())
			}
		}

	}

	if len(errstrings) > 0 {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	} else {
		return nil
	}

}

func getResponse(sessionsClient *cx.SessionsClient, agent *cxpb.Agent, test *types.Test, check types.Check, testInfo types.Tests, sessionId string) (*cxpb.DetectIntentResponse, error) {
	if check.Input.Type == "text" {
		global.Log.Infof("Input: type: %s, value: %s \n", check.Input.Type, check.Input.Text)

		return cxpkg.DetectIntentFromText(sessionsClient, agent, test.LocaleID, check.Input.Text, sessionId)
	} else {
		global.Log.Infof("Input: type: %s, value: %s \n", check.Input.Type, check.Input.Audio)

		audioFile := utils.GetRelativeFilePathFromParentFile(testInfo.File, check.Input.Audio)
		return cxpkg.DetectIntentFromAudio(sessionsClient, agent, test.LocaleID, audioFile, sessionId)
	}

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
