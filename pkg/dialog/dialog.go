package dialog

import (
	"fmt"
	"io"
	"os"

	"github.com/xavidop/dialogflow-cx-cli/internal/global"
	terminal "golang.org/x/term"
)

func Dialog() error {
	global.Log.Infof("Please press Ctrl+C whenever you want to stop the interaction.\n")

	if err := chat(); err != nil {
		return err
	}

	return nil
}

func chat() error {
	if !terminal.IsTerminal(0) || !terminal.IsTerminal(1) {
		return fmt.Errorf("stdin/stdout should be terminal")
	}
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		return err
	}
	defer terminal.Restore(0, oldState)
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
		fmt.Fprintln(term, rePrefix, line)
	}
}
