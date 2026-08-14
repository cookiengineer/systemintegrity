//go:build antergos || archlinux || manjaro

package pacman

import "github.com/cookiengineer/systemintegrity/types"
import "testing"

func TestParseVerificationLine(t *testing.T) {

	t.Run("parseVerificationLine(modification)", func(t *testing.T) {

		name, path, issues := parseVerificationLine("warning: glibc: /usr/lib/libc.so.6 (Modification time mismatch)")

		if name != "glibc" {
			t.Errorf("Expected name %q but got %q", "glibc", name)
		}

		if path != "/usr/lib/libc.so.6" {
			t.Errorf("Expected path %q but got %q", "/usr/lib/libc.so.6", path)
		}

		if len(issues) != 1 || issues[0].String() != string(types.PackageVerificationIssueModificationTimeMismatch) {
			t.Errorf("Expected issue %q but got %v", "modification_time_mismatch", issues)
		}

	})

	t.Run("parseVerificationLine(missing)", func(t *testing.T) {

		name, path, issues := parseVerificationLine("glibc: /etc/ld.so.conf (No such file or directory)")

		if name != "glibc" {
			t.Errorf("Expected name %q but got %q", "glibc", name)
		}

		if path != "/etc/ld.so.conf" {
			t.Errorf("Expected path %q but got %q", "/etc/ld.so.conf", path)
		}

		if len(issues) != 1 || issues[0].String() != string(types.PackageVerificationIssueMissingFile) {
			t.Errorf("Expected issue %q but got %v", "missing_file", issues)
		}

	})

	t.Run("parseVerificationLine(invalid)", func(t *testing.T) {

		name, path, issues := parseVerificationLine("garbage line without separator")

		if name != "" || path != "" || len(issues) != 0 {
			t.Errorf("Expected empty result but got %q, %q, %v", name, path, issues)
		}

	})

}

func TestToIssues(t *testing.T) {

	checks := []struct {
		reason string
		issue  types.PackageVerificationIssue
	}{
		{"SHA256 checksum mismatch", types.PackageVerificationIssueChecksumMismatch},
		{"failed to calculate SHA256 checksum", types.PackageVerificationIssueChecksumUnavailable},
		{"GID mismatch", types.PackageVerificationIssueGroupMismatch},
		{"Permissions mismatch", types.PackageVerificationIssueModeMismatch},
		{"Modification time mismatch", types.PackageVerificationIssueModificationTimeMismatch},
		{"No such file or directory", types.PackageVerificationIssueMissingFile},
		{"Permission denied", types.PackageVerificationIssuePermissionDenied},
		{"Symlink path mismatch", types.PackageVerificationIssueReadlinkMismatch},
		{"Size mismatch", types.PackageVerificationIssueSizeMismatch},
		{"UID mismatch", types.PackageVerificationIssueUserMismatch},
	}

	for _, check := range checks {

		t.Run("toIssues("+check.reason+")", func(t *testing.T) {

			issues := toIssues(check.reason)

			if len(issues) != 1 || issues[0].String() != check.issue.String() {
				t.Errorf("Expected issue %q but got %v", check.issue.String(), issues)
			}

		})

	}

}

func TestToRemediations(t *testing.T) {

	t.Run("toRemediations(non-issue)", func(t *testing.T) {

		remediations := toRemediations("glibc", []types.PackageVerificationIssue{
			types.PackageVerificationIssueModificationTimeMismatch,
		})

		if len(remediations) != 0 {
			t.Errorf("Expected 0 remediations for non-issue but got %v", remediations)
		}

	})

	t.Run("toRemediations(group)", func(t *testing.T) {

		remediations := toRemediations("glibc", []types.PackageVerificationIssue{
			types.PackageVerificationIssueGroupMismatch,
		})

		if len(remediations) != 1 {
			t.Errorf("Expected 1 remediation but got %v", remediations)
		} else if remediations[0].Issue.String() != string(types.PackageVerificationIssueGroupMismatch) {
			t.Errorf("Expected issue %q but got %q", "group_mismatch", remediations[0].Issue.String())
		} else if remediations[0].Command == "" {
			t.Errorf("Expected remediation to contain a Command")
		}

	})

}

func TestCollectVerification(t *testing.T) {

	t.Run("CollectVerification()", func(t *testing.T) {

		verifications := CollectVerification()

		for _, verification := range verifications {

			if verification.Name == "" {
				t.Errorf("Expected verification to contain a Name")
			}

			if len(verification.Files) == 0 {
				t.Errorf("Expected verification %q to contain Files", verification.Name)
			}

			for _, file := range verification.Files {

				if file.Path == "" || len(file.Issues) == 0 {
					t.Errorf("Expected verification %q file to contain Path and Issues", verification.Name)
				}

				for _, issue := range file.Issues {

					if issue.IsValid() == false {
						t.Errorf("Expected issue %q to be valid", issue.String())
					}

				}

			}

		}

	})

}
