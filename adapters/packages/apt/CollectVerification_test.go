//go:build debian || linuxmint || trisquel || ubuntu

package apt

import "github.com/cookiengineer/systemintegrity/types"
import "testing"

func TestParseVerificationLine(t *testing.T) {

	t.Run("ParseVerificationLine(digest)", func(t *testing.T) {

		path, issues := ParseVerificationLine("??5?????? c /etc/hosts")

		if path != "/etc/hosts" {
			t.Errorf("Expected path %q but got %q", "/etc/hosts", path)
		}

		if len(issues) != 1 || issues[0].String() != string(types.PackageVerificationIssueChecksumMismatch) {
			t.Errorf("Expected issue %q but got %v", "checksum_mismatch", issues)
		}

	})

	t.Run("ParseVerificationLine(mode)", func(t *testing.T) {

		path, issues := ParseVerificationLine("?M??????? /usr/bin/foo")

		if path != "/usr/bin/foo" {
			t.Errorf("Expected path %q but got %q", "/usr/bin/foo", path)
		}

		if len(issues) != 1 || issues[0].String() != string(types.PackageVerificationIssueFileTypeMismatch) {
			t.Errorf("Expected issue %q but got %v", "file_type_mismatch", issues)
		}

	})

	t.Run("ParseVerificationLine(missing)", func(t *testing.T) {

		path, issues := ParseVerificationLine("missing c /etc/some.conf (cannot open)")

		if path != "/etc/some.conf" {
			t.Errorf("Expected path %q but got %q", "/etc/some.conf", path)
		}

		if len(issues) != 1 || issues[0].String() != string(types.PackageVerificationIssueMissingFile) {
			t.Errorf("Expected issue %q but got %v", "missing_file", issues)
		}

	})

	t.Run("ParseVerificationLine(clean)", func(t *testing.T) {

		path, issues := ParseVerificationLine("????????? /etc/unchanged")

		if path != "/etc/unchanged" {
			t.Errorf("Expected path %q but got %q", "/etc/unchanged", path)
		}

		if len(issues) != 0 {
			t.Errorf("Expected no issues but got %v", issues)
		}

	})

}

func TestToRemediations(t *testing.T) {

	t.Run("toRemediations(checksum)", func(t *testing.T) {

		remediations := toRemediations("coreutils", []types.PackageVerificationIssue{
			types.PackageVerificationIssueChecksumMismatch,
		})

		if len(remediations) != 1 {
			t.Errorf("Expected 1 remediation but got %v", remediations)
		} else if remediations[0].Command == "" {
			t.Errorf("Expected remediation to contain a Command")
		} else if remediations[0].Issue.String() != string(types.PackageVerificationIssueChecksumMismatch) {
			t.Errorf("Expected issue %q but got %q", "checksum_mismatch", remediations[0].Issue.String())
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

			}

		}

	})

}
