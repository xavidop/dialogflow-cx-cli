package cx

import (
	"context"
	"errors"
	"fmt"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"github.com/xavidop/dialogflow-cx-cli/internal/utils"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func CreateFlowRESTClient(locationId string) (*cx.FlowsClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewFlowsRESTClient(ctx, credentials, endpoint)
	} else {
		return cx.NewFlowsRESTClient(ctx, endpoint)
	}

}

func CreateFlowGRPCClient(locationId string) (*cx.FlowsClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewFlowsClient(ctx, credentials, endpoint)
	} else {
		return cx.NewFlowsClient(ctx, endpoint)
	}

}

func GetNluModelType(nluModelType string) cxpb.NluSettings_ModelType {
	switch nluModelType {
	case "advanced":
		return cxpb.NluSettings_MODEL_TYPE_ADVANCED
	case "standard":
		return cxpb.NluSettings_MODEL_TYPE_STANDARD
	default:
		return cxpb.NluSettings_MODEL_TYPE_UNSPECIFIED
	}
}

func GetNluModelTrainingMode(nluModelTrainingMode string) cxpb.NluSettings_ModelTrainingMode {
	switch nluModelTrainingMode {
	case "manual":
		return cxpb.NluSettings_MODEL_TRAINING_MODE_MANUAL
	case "automatic":
		return cxpb.NluSettings_MODEL_TRAINING_MODE_AUTOMATIC
	default:
		return cxpb.NluSettings_MODEL_TRAINING_MODE_UNSPECIFIED
	}
}

func CreateFlow(flowClient *cx.FlowsClient, agent *cxpb.Agent, flowName, description, localeId, nluClassificationThreshold, nluModelType, nluModelTrainingMode string) (*cxpb.Flow, error) {
	ctx := context.Background()
	localeToUse := agent.GetDefaultLanguageCode()
	if localeId != "" {
		localeToUse = localeId
	}

	nluClassificationThresholdFloat, err := utils.ParseFloat(nluClassificationThreshold)
	if err != nil {
		return nil, err
	}

	reqCreateFlow := &cxpb.CreateFlowRequest{
		Parent:       agent.GetName(),
		LanguageCode: localeToUse,
		Flow: &cxpb.Flow{
			DisplayName: flowName,
			Description: description,
			NluSettings: &cxpb.NluSettings{
				ModelType:               GetNluModelType(nluModelType),
				ModelTrainingMode:       GetNluModelTrainingMode(nluModelTrainingMode),
				ClassificationThreshold: nluClassificationThresholdFloat,
			},
		},
	}

	return flowClient.CreateFlow(ctx, reqCreateFlow)
}

func UpdateFlow(flowClient *cx.FlowsClient, agent *cxpb.Agent, flowName, description, localeId, nluClassificationThreshold, nluModelType, nluModelTrainingMode string) (*cxpb.Flow, error) {
	ctx := context.Background()
	paths := []string{}

	localeToUse := agent.GetDefaultLanguageCode()
	if localeId != "" {
		localeToUse = localeId
	}

	flow, err := GetFlowIdByName(flowClient, agent, flowName)
	if err != nil {
		return nil, err
	}

	if description != "" {
		flow.Description = description
		paths = append(paths, "description")
	}
	if nluModelType != "" {
		flow.NluSettings.ModelType = GetNluModelType(nluModelType)
		paths = append(paths, "nlu_settings")
	}

	if nluModelTrainingMode != "" {
		flow.NluSettings.ModelTrainingMode = GetNluModelTrainingMode(nluModelTrainingMode)
		paths = append(paths, "nlu_settings")
	}

	if nluClassificationThreshold != "" {
		nluClassificationThresholdFloat, err := utils.ParseFloat(nluClassificationThreshold)
		if err != nil {
			return nil, err
		}
		flow.NluSettings.ClassificationThreshold = nluClassificationThresholdFloat
		paths = append(paths, "nlu_settings")
	}

	reqUpdateFlow := &cxpb.UpdateFlowRequest{
		LanguageCode: localeToUse,
		Flow:         flow,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: paths,
		},
	}

	return flowClient.UpdateFlow(ctx, reqUpdateFlow)
}

func DeleteFlow(flowClient *cx.FlowsClient, agent *cxpb.Agent, flowName string) error {
	ctx := context.Background()

	flow, err := GetFlowIdByName(flowClient, agent, flowName)
	if err != nil {
		return err
	}

	reqDeleteFlow := &cxpb.DeleteFlowRequest{
		Name: flow.GetName(),
	}
	return flowClient.DeleteFlow(ctx, reqDeleteFlow)
}

func TrainAll(flowClient *cx.FlowsClient, agent *cxpb.Agent) error {
	ctx := context.Background()

	reqList := &cxpb.ListFlowsRequest{
		Parent: agent.GetName(),
	}

	flows := flowClient.ListFlows(ctx, reqList)

	for flow, err := flows.Next(); err == nil; {
		if err := Train(flowClient, agent, flow.GetDisplayName(), flow.GetName()); err != nil {
			return err
		}
		flow, err = flows.Next()
		if flows.PageInfo().Remaining() == 0 {
			return nil
		}
	}

	return nil

}

func Train(flowClient *cx.FlowsClient, agent *cxpb.Agent, flowName, flowId string) error {
	ctx := context.Background()
	if flowId == "" {
		flow, err := GetFlowIdByName(flowClient, agent, flowName)
		if err != nil {
			return err
		}
		flowId = flow.GetName()
	}

	reqTrainFlow := &cxpb.TrainFlowRequest{
		Name: flowId,
	}
	op, err := flowClient.TrainFlow(ctx, reqTrainFlow)
	if err != nil {
		return err
	}

	return op.Wait(ctx)

}

func GetFlowIdByName(flowClient *cx.FlowsClient, agent *cxpb.Agent, flowName string) (*cxpb.Flow, error) {
	ctx := context.Background()

	reqList := &cxpb.ListFlowsRequest{
		Parent: agent.GetName(),
	}

	flows := flowClient.ListFlows(ctx, reqList)

	for flow, err := flows.Next(); err == nil; {
		if flow.DisplayName == flowName {
			return flow, nil
		}
		flow, err = flows.Next()
		if err != nil {
			return nil, err
		}

	}

	return nil, errors.New("flow not found")
}
