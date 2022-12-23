package tts

import (
	"os"

	"github.com/xavidop/dialogflow-cx-test-runner/internal/global"
)

func Synthesize(input string, locale string, output string) error {

	client, err := CreateTTSClient()
	if err != nil {
		return err
	}
	defer client.Close()

	resp, err := ExecuteSynthesize(client, input, locale)
	if err != nil {
		return err
	}

	err = os.WriteFile(output, resp.AudioContent, 0644)
	if err != nil {
		return err
	}
	global.Log.Infof("Audio content written to file: %v\n", output)

	return nil
}
