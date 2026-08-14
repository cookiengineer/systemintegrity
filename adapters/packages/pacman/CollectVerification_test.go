//go:build antergos || archlinux || manjaro

package pacman

import "testing"

func TestParseVerificationLine(t *testing.T) {

	t.Run("parseVerificationLine(modification)", func(t *testing.T) {

		name, path, reason := parseVerificationLine("warning: glibc: /usr/lib/libc.so.6 (Modification time mismatch)")

		if name != "glibc" {
			t.Errorf("Expected name %q but got %q", "glibc", name)
		}

		if path != "/usr/lib/libc.so.6" {
			t.Errorf("Expected path %q but got %q", "/usr/lib/libc.so.6", path)
		}

		if reason != "Modification time mismatch" {
			t.Errorf("Expected reason %q but got %q", "Modification time mismatch", reason)
		}

	})

	t.Run("parseVerificationLine(missing)", func(t *testing.T) {

		name, path, reason := parseVerificationLine("glibc: /etc/ld.so.conf (No such file or directory)")

		if name != "glibc" {
			t.Errorf("Expected name %q but got %q", "glibc", name)
		}

		if path != "/etc/ld.so.conf" {
			t.Errorf("Expected path %q but got %q", "/etc/ld.so.conf", path)
		}

		if reason != "No such file or directory" {
			t.Errorf("Expected reason %q but got %q", "No such file or directory", reason)
		}

	})

	t.Run("parseVerificationLine(invalid)", func(t *testing.T) {

		name, path, reason := parseVerificationLine("garbage line without separator")

		if name != "" || path != "" || reason != "" {
			t.Errorf("Expected empty result but got %q, %q, %q", name, path, reason)
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

				if file.Path == "" || file.Reason == "" {
					t.Errorf("Expected verification %q file to contain Path and Reason", verification.Name)
				}

			}

		}

	})

}
