package utils

import (
	"context"
	"strings"

	"github.com/xavidop/dialogflow-cx-cli/internal/global"

	"github.com/google/go-github/v48/github"
)

func CheckAvailableUpdate(currVersion string, output bool) (string, error) {

	// Don't check version while in development and local compiling
	if global.VersionString == "development" {
		return "", nil
	}

	remoteVersion := ""

	client := github.NewClient(nil)

	// list all repositories for the authenticated user
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), global.RepoOwner, global.RepoName)

	if err != nil {
		return "", err
	}

	remoteVersion = strings.ReplaceAll(*release.Name, "v", "")
	// Compare the latest semver with the current semver
	if isRemoteVersionNewer(currVersion, remoteVersion) {
		// Returns the latest semver if it is higher than current semver
		return remoteVersion, nil
	}

	return "", nil
}

// isRemoteVersionNewer compares versions
func isRemoteVersionNewer(local string, remote string) bool {

	return local != remote
}
