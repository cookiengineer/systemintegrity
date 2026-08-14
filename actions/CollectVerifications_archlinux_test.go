//go:build antergos || archlinux || manjaro

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

	assertIssues(t, &system, []issueExpectation{
		{"/usr/bin/ssh", types.PackageVerificationIssueChecksumMismatch},
		{"/usr/bin/ssh-copy-id", types.PackageVerificationIssueMissingFile},
		{"/usr/bin/ssh-keygen", types.PackageVerificationIssueModeMismatch},
		{"/usr/bin/ssh-add", types.PackageVerificationIssueUserMismatch},
		{"/usr/bin/scp", types.PackageVerificationIssueGroupMismatch},
		{"/usr/bin/sftp", types.PackageVerificationIssueSizeMismatch},
		{"/usr/bin/ssh-keyscan", types.PackageVerificationIssueModificationTimeMismatch},
	})

}
