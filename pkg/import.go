package pkg

import (
	"fmt"
	"io"

	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/xavidop/dialogflow-cx-test-runner/internal/global"
	"github.com/xavidop/dialogflow-cx-test-runner/internal/types"
	"github.com/xavidop/dialogflow-cx-test-runner/internal/utils"

	"github.com/google/uuid"

	"google.golang.org/api/option"
)

func ImportType(typeName string, file string, header bool) error {

	projectId := "test-cx-346408"
	locationId := "us-central1"
	agentName := "test-agent"
	sessionId := uuid.NewString()
	environmentName := "production"
	localeId := "en"

	credentials := option.WithCredentialsFile("credentials.json")
	endpointString := fmt.Sprintf("%s-dialogflow.googleapis.com", locationId)
	endpoint := option.WithEndpoint(endpointString)
	fmt.Sprintf("projects/%s/locations/%s", projectId, locationId, credentials, localeId, sessionId, agentName, environmentName, endpoint)

	destinationFile := typeName + ".yaml"
	if err := utils.ValidateFileType(file); err != nil {
		return err
	}

	reader, f, err := utils.ReadCsvFile(file)
	defer f.Close()

	if err != nil {
		return err
	}

	googleActionType := types.GoogleActionType{}
	googleActionType.Synonym = &types.Synonym{}
	googleActionType.Synonym.Entities = map[string]*types.Synonyms{}
	googleActionType.Synonym.MatchType = "EXACT_MATCH"

	skipFirst := header
	for {
		record, err := reader.Read()
		if skipFirst {
			skipFirst = false
			continue
		}
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if global.Verbose {
			log.Info(record)
			log.Info(len(record))
		}
		synonyms := types.Synonyms{}
		synonyms.Synonyms = utils.DeleteEmpty(record)
		// the ID es the second column
		googleActionType.Synonym.Entities[record[1]] = &synonyms

	}

	b, err := yaml.Marshal(googleActionType)
	if err != nil {
		return err
	}

	if err := utils.CheckIfFileExists(destinationFile); err == nil {
		prompt := promptui.Select{
			Label: "The file " + destinationFile + " already exists. Do you want to overwrite it?",
			Items: []string{"Yes", "No"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			return err
		}
		if result == "Yes" {
			if err = utils.WriteFile(b, destinationFile); err != nil {
				return err
			}

		} else {
			log.Infof("Import aborted.")
			return nil
		}
	} else {
		if err = utils.WriteFile(b, destinationFile); err != nil {
			return err
		}
	}

	log.Infof("Import succesful. You can find the type created at %s", destinationFile)
	return nil
}
