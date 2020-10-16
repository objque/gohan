package version

import "fmt"

//nolint:gochecknoglobals
var (
	// BuildTime is a time label of the moment when the binary was built
	BuildTime = "unset"
	// Commit is a last commit hash at the moment when the binary was built
	Commit = "unset"
	// Release is a semantic version of current build
	Release = "unset"

	FullInfo = fmt.Sprintf(
		"build info: commit: %s, build time: %s, release: %s",
		Commit, BuildTime, Release)
)
