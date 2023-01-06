package cx

import (
	"context"
	"errors"
	"fmt"
	"strings"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	"google.golang.org/api/option"
)

func CreateEntityTypesRESTClient(locationId string) (*cx.EntityTypesClient, error) {
	ctx := context.Background()
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com:443", locationId)
	endpoint := option.WithEndpoint(endpointString)

	if global.Credentials != "" {
		credentials := option.WithCredentialsFile(global.Credentials)
		return cx.NewEntityTypesRESTClient(ctx, credentials, endpoint)
	} else {
		return cx.NewEntityTypesRESTClient(ctx, endpoint)
	}

}

func CreateEntityType(entityTypesClient *cx.EntityTypesClient, agent *cxpb.Agent, entityTypeName, localeId string, entities []string) (*cxpb.EntityType, error) {
	ctx := context.Background()
	localeToUse := agent.GetDefaultLanguageCode()
	if localeId != "" {
		localeToUse = localeId
	}

	entityTypesEntities, err := CreateEntityTypesEntities(entities)
	if err != nil {
		return nil, err
	}

	reqCreateEntityType := &cxpb.CreateEntityTypeRequest{
		Parent:       agent.GetName(),
		LanguageCode: localeToUse,

		EntityType: &cxpb.EntityType{
			DisplayName: entityTypeName,
			Entities:    entityTypesEntities,
			Kind:        cxpb.EntityType_KIND_MAP,
		},
	}

	return entityTypesClient.CreateEntityType(ctx, reqCreateEntityType)
}

func CreateEntityTypesEntities(entities []string) ([]*cxpb.EntityType_Entity, error) {
	entityTypesEntities := []*cxpb.EntityType_Entity{}
	for _, entity := range entities {
		var entityTypesEntity *cxpb.EntityType_Entity
		if strings.Contains(entity, "@") {
			synonyms := strings.Split(strings.Split(entity, "@")[1], "|")
			value := strings.TrimSpace(strings.Split(entity, "@")[0])
			entityTypesEntity = &cxpb.EntityType_Entity{
				Value:    value,
				Synonyms: synonyms,
			}
		} else {
			return nil, errors.New("synonyms not provided")
		}
		entityTypesEntities = append(entityTypesEntities, entityTypesEntity)
	}
	return entityTypesEntities, nil
}

func GetEntityTypeIdByName(entityTypesClient *cx.EntityTypesClient, agent *cxpb.Agent, entityTypeName string) (*cxpb.EntityType, error) {
	ctx := context.Background()

	reqEntityTypeList := &cxpb.ListEntityTypesRequest{
		Parent: agent.GetName(),
	}

	entityTypes := entityTypesClient.ListEntityTypes(ctx, reqEntityTypeList)

	for entityType, err := entityTypes.Next(); err == nil; {
		if entityType.DisplayName == entityTypeName {
			return entityType, nil
		}
		entityType, err = entityTypes.Next()
		if err != nil {
			return nil, err
		}

	}

	return nil, errors.New("entity type not found")

}

func DeleteEntityType(entityTypesClient *cx.EntityTypesClient, agent *cxpb.Agent, entityTypeName string, force bool) error {
	ctx := context.Background()

	entityType, err := GetEntityTypeIdByName(entityTypesClient, agent, entityTypeName)
	if err != nil {
		return err
	}

	reqDeleteEntityType := &cxpb.DeleteEntityTypeRequest{
		Name:  entityType.GetName(),
		Force: force,
	}
	return entityTypesClient.DeleteEntityType(ctx, reqDeleteEntityType)
}
