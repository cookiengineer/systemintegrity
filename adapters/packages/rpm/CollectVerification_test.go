//go:build redhat || centos || oraclelinux || almalinux || rockylinux || fedora || amazonlinux || opensuse || suse_desktop || suse_server

package rpm

import "github.com/cookiengineer/systemintegrity/types"
import "testing"

func TestParseVerificationLine(t *testing.T) {

	t.Run("ParseVerificationLine(multi)", func(t *testing.T) {

		path, issues := ParseVerificationLine("S.5....T.  c /etc/foo.conf")

		if path != "/etc/foo.conf" {
			t.Errorf("Expected path %q but got %q", "/etc/foo.conf", path)
		}

		if len(issues) != 3 {
			t.Errorf("Expected 3 issues but got %v", issues)
		} else {
			if issues[0].String() != string(types.PackageVerificationIssueSizeMismatch) {
				t.Errorf("Expected issue %q but got %q", "size_mismatch", issues[0].String())
			}
			if issues[1].String() != string(types.PackageVerificationIssueChecksumMismatch) {
				t.Errorf("Expected issue %q but got %q", "checksum_mismatch", issues[1].String())
			}
			if issues[2].String() != string(types.PackageVerificationIssueModificationTimeMismatch) {
				t.Errorf("Expected issue %q but got %q", "modification_time_mismatch", issues[2].String())
			}
		}

	})

	t.Run("ParseVerificationLine(user/group)", func(t *testing.T) {

		path, issues := ParseVerificationLine("..?......    /usr/bin/foo")

		if path != "/usr/bin/foo" {
			t.Errorf("Expected path %q but got %q", "/usr/bin/foo", path)
		}

		if len(issues) != 1 || issues[0].String() != string(types.PackageVerificationIssueChecksumUnavailable) {
			t.Errorf("Expected issue %q but got %v", "checksum_unavailable", issues)
		}

	})

	t.Run("ParseVerificationLine(missing)", func(t *testing.T) {

		path, issues := ParseVerificationLine("missing     /usr/share/doc/foo/README")

		if path != "/usr/share/doc/foo/README" {
			t.Errorf("Expected path %q but got %q", "/usr/share/doc/foo/README", path)
		}

		if len(issues) != 1 || issues[0].String() != string(types.PackageVerificationIssueMissingFile) {
			t.Errorf("Expected issue %q but got %v", "missing_file", issues)
		}

	})

}

func TestToRemediations(t *testing.T) {

	t.Run("ToRemediations(group/rpm)", func(t *testing.T) {

		remediations := ToRemediations("rpm", "foo", []types.PackageVerificationIssue{
			types.PackageVerificationIssueGroupMismatch,
		})

		if len(remediations) != 1 {
			t.Errorf("Expected 1 remediation but got %v", remediations)
		} else if remediations[0].Command != "rpm --setugids foo" {
			t.Errorf("Expected command %q but got %q", "rpm --setugids foo", remediations[0].Command)
		}

	})

	t.Run("ToRemediations(mode/dnf)", func(t *testing.T) {

		remediations := ToRemediations("dnf", "foo", []types.PackageVerificationIssue{
			types.PackageVerificationIssueModeMismatch,
		})

		if len(remediations) != 1 {
			t.Errorf("Expected 1 remediation but got %v", remediations)
		} else if remediations[0].Command != "rpm --setperms foo" {
			t.Errorf("Expected command %q but got %q", "rpm --setperms foo", remediations[0].Command)
		}

	})

	t.Run("ToRemediations(size/zypper)", func(t *testing.T) {

		remediations := ToRemediations("zypper", "foo", []types.PackageVerificationIssue{
			types.PackageVerificationIssueSizeMismatch,
		})

		if len(remediations) != 1 {
			t.Errorf("Expected 1 remediation but got %v", remediations)
		} else if remediations[0].Command != "zypper install --force foo" {
			t.Errorf("Expected command %q but got %q", "zypper install --force foo", remediations[0].Command)
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
