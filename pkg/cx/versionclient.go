package cx

import (
	"context"
	"errors"
	"fmt"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func CreateVersionRESTClient(locationId string) (*cx.VersionsClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewVersionsRESTClient(ctx, credentials, endpoint)
	} else {
		return cx.NewVersionsRESTClient(ctx, endpoint)
	}

}

func CreateVersionGRPCClient(locationId string) (*cx.VersionsClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewVersionsClient(ctx, credentials, endpoint)
	} else {
		return cx.NewVersionsClient(ctx, endpoint)
	}

}

func CreateVersion(versionClient *cx.VersionsClient, flow *cxpb.Flow, description, name string) (*cxpb.Version, error) {
	ctx := context.Background()

	reqCreateVersion := &cxpb.CreateVersionRequest{
		Parent: flow.GetName(),
		Version: &cxpb.Version{
			DisplayName: name,
			Description: description,
		},
	}

	op, err := versionClient.CreateVersion(ctx, reqCreateVersion)
	if err != nil {
		return nil, err
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func UpdateVersion(versionClient *cx.VersionsClient, flow *cxpb.Flow, description, name string) (*cxpb.Version, error) {
	ctx := context.Background()

	version, err := GetVersionIdByName(versionClient, flow, name)
	if err != nil {
		return nil, err
	}

	if description != "" {
		version.Description = description
	}

	reqUpdateVersion := &cxpb.UpdateVersionRequest{
		Version: version,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"description"},
		},
	}

	return versionClient.UpdateVersion(ctx, reqUpdateVersion)
}

func GetVersionIdByName(versionClient *cx.VersionsClient, flow *cxpb.Flow, name string) (*cxpb.Version, error) {
	ctx := context.Background()

	reqVersionList := &cxpb.ListVersionsRequest{
		Parent: flow.GetName(),
	}

	versions := versionClient.ListVersions(ctx, reqVersionList)

	for version, err := versions.Next(); err == nil; {
		if version.DisplayName == name {
			return version, nil
		}
		version, err = versions.Next()
		if err != nil {
			return nil, err
		}

	}

	return nil, errors.New("version not found")

}

func DeleteVersion(versionClient *cx.VersionsClient, flow *cxpb.Flow, name string) error {
	ctx := context.Background()

	version, err := GetVersionIdByName(versionClient, flow, name)
	if err != nil {
		return err
	}

	reqDeleteVersion := &cxpb.DeleteVersionRequest{
		Name: version.GetName(),
	}

	return versionClient.DeleteVersion(ctx, reqDeleteVersion)
}
