package caches

import "github.com/cookiengineer/systemintegrity/structs"
import "github.com/cookiengineer/systemintegrity/types"
import "testing"

func mockupVerifications() *Verifications {

	verifications := NewVerifications()

	verifications.Add(structs.PackageVerification{
		Name:    "glibc",
		Version: types.ToVersion("0:2.41+r48"),
		Manager: types.ToManager("pacman"),
		Files: []structs.FileVerification{
			{
				Path: "/usr/lib/libc.so.6",
				Issues: []types.PackageVerificationIssue{
					types.PackageVerificationIssueModificationTimeMismatch,
				},
			},
		},
	})

	verifications.Add(structs.PackageVerification{
		Name:    "bind",
		Version: types.ToVersion("0:9.20+r1"),
		Manager: types.ToManager("pacman"),
		Files: []structs.FileVerification{
			{
				Path: "/var/named/127.0.0.zone",
				Issues: []types.PackageVerificationIssue{
					types.PackageVerificationIssuePermissionDenied,
				},
				Remediations: []types.Remediation{
					types.NewRemediation("pacman", types.PackageVerificationIssuePermissionDenied, "run as root"),
				},
			},
		},
	})

	return verifications

}

func TestVerifications(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		verifications := mockupVerifications()

		if verifications.Length() != 2 {
			t.Errorf("Expected %d verifications but got %d", 2, verifications.Length())
		}

	})

	t.Run("Get()", func(t *testing.T) {

		verifications := mockupVerifications()

		verification := verifications.Get("pacman:glibc:0:2.41.0r48")

		if verification == nil {
			t.Errorf("Expected nil to be %s version %s", "glibc", "0:2.41.0r48")
		} else if verification.Name != "glibc" || verification.Version.String() != "0:2.41.0r48" {
			t.Errorf("Expected %s version %s to be %s version %s", verification.Name, verification.Version.String(), "glibc", "0:2.41.0r48")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		verifications := mockupVerifications()

		verifications.Remove("pacman:glibc:0:2.41.0r48")

		verification := verifications.Get("pacman:glibc:0:2.41.0r48")

		if verification != nil {
			t.Errorf("Expected %s version %s to be nil", verification.Name, verification.Version.String())
		}

	})

}
