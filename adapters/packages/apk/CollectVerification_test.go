//go:build alpinelinux

package apk

import "github.com/cookiengineer/systemintegrity/types"
import "testing"

func TestParseVerification(t *testing.T) {

	t.Run("ParseVerification(all)", func(t *testing.T) {

		buffer := "" +
			"- mode=755 uid=0 gid=0 hash=78fc93c415261ee408be691bbb5645488edbf18c\n" +
			"+ mode=755 uid=0 gid=0 hash=8c68c6290f44e2f6c75d00f13be8b11bf2300547\n" +
			"U usr/bin/ssh\n" +
			"- mode=755 uid=0 gid=0 hash=301cb080c1928d897dc5103b07c67936aff4b012\n" +
			"X usr/bin/ssh-copy-id\n" +
			"- mode=755 uid=0 gid=0 hash=0e4f1150f7afdbe1524530e169ce328c7c6ac648\n" +
			"+ mode=750 uid=0 gid=0 hash=0e4f1150f7afdbe1524530e169ce328c7c6ac648\n" +
			"M usr/bin/ssh-keygen\n" +
			"- mode=755 uid=0 gid=0 hash=edd3a1490821d7925210b0d60c7cbdcc00644095\n" +
			"+ mode=755 uid=12345 gid=0 hash=edd3a1490821d7925210b0d60c7cbdcc00644095\n" +
			"M usr/bin/ssh-add\n" +
			"- mode=755 uid=0 gid=0 hash=1c41b1c1a133e399411dfffdfcc081ff3747af16\n" +
			"+ mode=755 uid=0 gid=12345 hash=1c41b1c1a133e399411dfffdfcc081ff3747af16\n" +
			"M usr/bin/scp\n"

		files := ParseVerification(buffer)

		if len(files) != 5 {
			t.Fatalf("Expected 5 files but got %d", len(files))
		}

		if issues, ok := files["/usr/bin/ssh"]; ok == false || len(issues) != 1 || issues[0].String() != string(types.PackageVerificationIssueChecksumMismatch) {
			t.Errorf("Expected /usr/bin/ssh checksum_mismatch but got %v", files["/usr/bin/ssh"])
		}

		if issues, ok := files["/usr/bin/ssh-copy-id"]; ok == false || len(issues) != 1 || issues[0].String() != string(types.PackageVerificationIssueMissingFile) {
			t.Errorf("Expected /usr/bin/ssh-copy-id missing_file but got %v", files["/usr/bin/ssh-copy-id"])
		}

		if issues, ok := files["/usr/bin/ssh-keygen"]; ok == false || len(issues) != 1 || issues[0].String() != string(types.PackageVerificationIssueModeMismatch) {
			t.Errorf("Expected /usr/bin/ssh-keygen mode_mismatch but got %v", files["/usr/bin/ssh-keygen"])
		}

		if issues, ok := files["/usr/bin/ssh-add"]; ok == false || len(issues) != 1 || issues[0].String() != string(types.PackageVerificationIssueUserMismatch) {
			t.Errorf("Expected /usr/bin/ssh-add user_mismatch but got %v", files["/usr/bin/ssh-add"])
		}

		if issues, ok := files["/usr/bin/scp"]; ok == false || len(issues) != 1 || issues[0].String() != string(types.PackageVerificationIssueGroupMismatch) {
			t.Errorf("Expected /usr/bin/scp group_mismatch but got %v", files["/usr/bin/scp"])
		}

	})

	t.Run("ParseVerification(directories ignored)", func(t *testing.T) {

		buffer := "" +
			"- mode=755 uid=0 gid=0\n" +
			"+ mode=555 uid=65534 gid=65534\n" +
			"m proc/\n"

		files := ParseVerification(buffer)

		if len(files) != 0 {
			t.Errorf("Expected 0 files but got %v", files)
		}

	})

	t.Run("ParseVerification(empty)", func(t *testing.T) {

		files := ParseVerification("")

		if len(files) != 0 {
			t.Errorf("Expected 0 files but got %v", files)
		}

	})

}

func TestToRemediations(t *testing.T) {

	t.Run("toRemediations(checksum)", func(t *testing.T) {

		remediations := toRemediations("openssh-client-default", []types.PackageVerificationIssue{
			types.PackageVerificationIssueChecksumMismatch,
		})

		if len(remediations) != 1 {
			t.Errorf("Expected 1 remediation but got %v", remediations)
		} else if remediations[0].Command != "apk fix openssh-client-default" {
			t.Errorf("Expected command %q but got %q", "apk fix openssh-client-default", remediations[0].Command)
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
