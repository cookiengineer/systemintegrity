//go:build debian || linuxmint || trisquel || ubuntu

package actions

import "github.com/cookiengineer/systemintegrity/structs"
import "github.com/cookiengineer/systemintegrity/types"
import "testing"

func TestCollectVerifications(t *testing.T) {

	console := structs.NewConsole(nil, nil, 0)
	system := structs.NewSystem()

	CollectPackages(console, &system)

	if len(system.Packages) == 0 {
		t.Fatalf("Expected package index to contain packages")
	}

	CollectVerifications(console, &system)

	validateVerifications(t, &system)

	// dpkg --verify only reports digest (5), missing, and a heuristic mode (M)
	// check that fails when a path is no longer a regular file (file type
	// changed, e.g. regular file replaced by a symlink).
	assertIssues(t, &system, []issueExpectation{
		{"/usr/bin/ssh", types.PackageVerificationIssueChecksumMismatch},
		{"/usr/bin/ssh-copy-id", types.PackageVerificationIssueMissingFile},
		{"/usr/bin/scp", types.PackageVerificationIssueFileTypeMismatch},
	})

}
