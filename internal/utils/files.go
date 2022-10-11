package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/hajimehoshi/go-mp3"
)

func CheckIfFileExists(file string) error {

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return err
	}

	return nil
}

func WriteFile(b []byte, file string) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	err = ioutil.WriteFile(file, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetRelativeFilePathFromParentFile(parentFile string, file string) string {
	base := filepath.Dir(parentFile)

	if !filepath.IsAbs(file) {
		return path.Join(base, file)
	} else {
		return file
	}
}

func GetAudioSampleHertzRate(mp3file string) (int, error) {
	f, err := os.Open(mp3file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)

	if err != nil {
		return 0, err
	}

	return d.SampleRate(), nil
}
