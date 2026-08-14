package types

import "testing"

func TestPackageVerificationIssue(t *testing.T) {

	issues := []struct {
		value       PackageVerificationIssue
		is_issue    bool
		description string
	}{
		{PackageVerificationIssueChecksumMismatch, true, "Checksum mismatch"},
		{PackageVerificationIssueChecksumUnavailable, true, "Checksum unavailable"},
		{PackageVerificationIssueDeviceMismatch, true, "Device mismatch"},
		{PackageVerificationIssueFileTypeMismatch, true, "File type mismatch"},
		{PackageVerificationIssueGroupMismatch, true, "Group ownership mismatch"},
		{PackageVerificationIssueModeMismatch, true, "Permissions mismatch"},
		{PackageVerificationIssueModificationTimeMismatch, false, "Modification time mismatch"},
		{PackageVerificationIssueMissingFile, true, "Missing file"},
		{PackageVerificationIssuePermissionDenied, true, "Permission denied"},
		{PackageVerificationIssueReadlinkMismatch, true, "Symlink path mismatch"},
		{PackageVerificationIssueSizeMismatch, true, "Size mismatch"},
		{PackageVerificationIssueUserMismatch, true, "User ownership mismatch"},
		{PackageVerificationIssueCapabilitiesMismatch, true, "Capabilities mismatch"},
	}

	for _, check := range issues {

		t.Run("Issue("+check.value.String()+")", func(t *testing.T) {

			if check.value.IsValid() == false {
				t.Errorf("Expected issue %q to be valid", check.value.String())
			}

			if check.value.IsIssue() != check.is_issue {
				t.Errorf("Expected issue %q IsIssue() to be %v but got %v", check.value.String(), check.is_issue, check.value.IsIssue())
			}

			if check.value.Description() != check.description {
				t.Errorf("Expected description %q but got %q", check.description, check.value.Description())
			}

		})

	}

}

func TestParsePackageVerificationIssue(t *testing.T) {

	t.Run("Parse(valid)", func(t *testing.T) {

		issue := ParsePackageVerificationIssue("group_mismatch")

		if issue == nil || issue.String() != "group_mismatch" {
			t.Errorf("Expected issue %q but got %v", "group_mismatch", issue)
		}

	})

	t.Run("Parse(invalid)", func(t *testing.T) {

		issue := ParsePackageVerificationIssue("not_an_issue")

		if issue != nil {
			t.Errorf("Expected nil but got %v", issue)
		}

	})

}

func TestRemediation(t *testing.T) {

	t.Run("Remediation(valid)", func(t *testing.T) {

		remediation := NewRemediation("pacman", PackageVerificationIssueGroupMismatch, "pacman -S glibc")

		if remediation.IsValid() == false {
			t.Errorf("Expected remediation to be valid")
		}

	})

	t.Run("Remediation(invalid manager)", func(t *testing.T) {

		remediation := NewRemediation("not-a-manager", PackageVerificationIssueGroupMismatch, "pacman -S glibc")

		if remediation.IsValid() == true {
			t.Errorf("Expected remediation to be invalid")
		}

	})

	t.Run("Remediation(empty command)", func(t *testing.T) {

		remediation := NewRemediation("pacman", PackageVerificationIssueGroupMismatch, "")

		if remediation.IsValid() == true {
			t.Errorf("Expected remediation to be invalid")
		}

	})

}
