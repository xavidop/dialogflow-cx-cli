package dialog

import (
	"fmt"
	"io"
	"os"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	"github.com/google/uuid"
	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-cli/pkg/cx"
	terminal "golang.org/x/term"
)

func Dialog(locationID, projectID, agentName, localeId string) error {
	global.Log.Infof("Please press Ctrl+C whenever you want to stop the interaction.\n")

	agentClient, err := cxpkg.CreateAgentRESTClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	localeToUse := agent.GetDefaultLanguageCode()
	if localeId != "" {
		localeToUse = localeId
	}

	sessionsClient, err := cxpkg.CreateSessionRESTClient(locationID)
	if err != nil {
		return err
	}
	defer sessionsClient.Close()

	sessionId := uuid.NewString()

	if err := chat(sessionsClient, agent, localeToUse, sessionId); err != nil {
		return err
	}

	return nil
}

func chat(sessionsClient *cx.SessionsClient, agent *cxpb.Agent, locale, sessionId string) error {

	if !terminal.IsTerminal(0) || !terminal.IsTerminal(1) {
		return fmt.Errorf("stdin/stdout should be terminal")
	}
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		return err
	}
	defer restore(oldState)
	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	term := terminal.NewTerminal(screen, "")
	term.SetPrompt(string(term.Escape.Red) + "User> " + string(term.Escape.Reset))

	rePrefix := string(term.Escape.Cyan) + "Agent>" + string(term.Escape.Reset)

	for {
		line, err := term.ReadLine()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if line == "" {
			continue
		}
		response, _ := cxpkg.DetectIntentFromText(sessionsClient, agent, locale, line, sessionId)

		for _, message := range response.GetQueryResult().GetResponseMessages() {
			if message.GetEndInteraction() != nil {
				return nil
			}

			for _, txtToShow := range message.GetText().GetText() {

				fmt.Fprintln(term, rePrefix, txtToShow)
			}

		}

	}
}

func restore(oldState *terminal.State) {
	if err := terminal.Restore(0, oldState); err != nil {
		global.Log.Errorf("failed to restore terminal: %v", err)
	}
}
