package stt

import (
	"time"

	"github.com/xavidop/dialogflow-cx-cli/internal/global"
)

func Recognize(input string, locale string) error {

	client, err := CreateSTTClient()
	if err != nil {
		return err
	}
	defer client.Close()

	resp, err := ExecuteRecognition(client, input, locale)
	if err != nil {
		return err
	}

	if global.Output == "json" {
		global.Log.Infof("%v\n", resp.GetResults())
	} else {

		for _, resp := range resp.GetResults() {
			if global.Verbose {
				global.Log.Infof("Duration time: %d miliseconds\n", resp.GetResultEndTime().GetNanos()/int32(time.Millisecond))
				global.Log.Infof("Detections: %d \n", len(resp.GetAlternatives()))

			}
			i := 1
			for _, alternative := range resp.GetAlternatives() {
				global.Log.Infof("%d. Text detected: %s\n", i, alternative.GetTranscript())
				if global.Verbose {
					global.Log.Infof("%d. Confidence: %f%% \n", i, alternative.GetConfidence()*100)
				}
				i++
			}
		}
	}

	return nil
}
