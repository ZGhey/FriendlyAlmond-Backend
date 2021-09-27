package version

import "fmt"

var (
	// RELEASE returns the release version
	RELEASE = "UNKNOWN"
	// REPO returns the git repository URL
	REPO = "UNKNOWN"
	// COMMIT returns the short sha from git
	COMMIT = "UNKNOWN"

	SERVICENAME = "UNKNOWN"

	BUILDTIME = "UNKNOWN"

	API   = "v1"
	Short = fmt.Sprintf("ops %s", RELEASE)
	Long  = fmt.Sprintf("ops release: %s, repo: %s, commit: %s, service: %s build: %s",
		RELEASE, REPO, COMMIT, SERVICENAME, BUILDTIME)
)
