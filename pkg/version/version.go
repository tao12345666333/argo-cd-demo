package version

import (
	"bytes"
	"fmt"
)

var (
	// The following fields are populated at build time using -ldflags -X.
	_buildVersion     = "unknown"
	_buildGitRevision = "unknown"
)

// Short produces a single-line version info with format:
// <version>-<git revision>
func Short() string {
	return fmt.Sprintf("%s-%s", _buildVersion, _buildGitRevision)
}

// Long produces a verbose version info with format:
// Version: xxx
// Git SHA: xxx
func Long() string {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintln(buf, "Version:", _buildVersion)
	fmt.Fprintln(buf, "Git SHA:", _buildGitRevision)
	return buf.String()
}
