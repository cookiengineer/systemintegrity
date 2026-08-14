package pacman

import "github.com/cookiengineer/systemintegrity/structs"
import "github.com/cookiengineer/systemintegrity/types"
import "bytes"
import "os"
import "os/exec"
import "strings"

func toIssues(reason string) []types.PackageVerificationIssue {

	issues := make([]types.PackageVerificationIssue, 0)

	reason = strings.TrimSpace(reason)

	if reason == "SHA256 checksum mismatch" {
		issues = append(issues, types.PackageVerificationIssueChecksumMismatch)
	} else if reason == "failed to calculate SHA256 checksum" {
		issues = append(issues, types.PackageVerificationIssueChecksumUnavailable)
	} else if reason == "GID mismatch" {
		issues = append(issues, types.PackageVerificationIssueGroupMismatch)
	} else if reason == "Permissions mismatch" {
		issues = append(issues, types.PackageVerificationIssueModeMismatch)
	} else if reason == "Modification time mismatch" {
		issues = append(issues, types.PackageVerificationIssueModificationTimeMismatch)
	} else if reason == "No such file or directory" {
		issues = append(issues, types.PackageVerificationIssueMissingFile)
	} else if reason == "Permission denied" {
		issues = append(issues, types.PackageVerificationIssuePermissionDenied)
	} else if reason == "Symlink path mismatch" {
		issues = append(issues, types.PackageVerificationIssueReadlinkMismatch)
	} else if reason == "Size mismatch" {
		issues = append(issues, types.PackageVerificationIssueSizeMismatch)
	} else if reason == "UID mismatch" {
		issues = append(issues, types.PackageVerificationIssueUserMismatch)
	}

	return issues

}

func toRemediations(name string, issues []types.PackageVerificationIssue) []types.Remediation {

	remediations := make([]types.Remediation, 0)

	for i := 0; i < len(issues); i++ {

		issue := issues[i]

		if issue.IsIssue() == false {
			continue
		}

		var command string

		if issue.String() == string(types.PackageVerificationIssuePermissionDenied) {
			command = "run as root"
		} else {
			command = "pacman -S --overwrite '*' --noconfirm " + name
		}

		remediation := types.NewRemediation("pacman", issue, command)

		if remediation.IsValid() {
			remediations = append(remediations, remediation)
		}

	}

	return remediations

}

func parseVerificationLine(line string) (string, string, []types.PackageVerificationIssue) {

	var name string
	var path string
	var issues []types.PackageVerificationIssue

	line = strings.TrimSpace(line)

	if strings.HasPrefix(line, "warning: ") {
		line = strings.TrimSpace(line[9:])
	} else if strings.HasPrefix(line, "error: ") {
		line = strings.TrimSpace(line[7:])
	}

	// "pkgname: /path/to/file (reason)"
	if strings.Contains(line, ": ") {

		name = strings.TrimSpace(line[0:strings.Index(line, ": ")])
		rest := strings.TrimSpace(line[strings.Index(line, ": ")+2:])

		if strings.Contains(rest, " (") && strings.HasSuffix(rest, ")") {

			last := strings.LastIndex(rest, " (")

			if last > 0 {

				path = strings.TrimSpace(rest[0:last])
				reason := strings.TrimSpace(rest[last+2 : len(rest)-1])

				issues = toIssues(reason)

			}

		}

	}

	return name, path, issues

}

func CollectVerification() []structs.PackageVerification {

	var collected []structs.PackageVerification

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		verifications := make(map[string]*structs.PackageVerification)

		var stdout bytes.Buffer
		var stderr bytes.Buffer

		cmd := exec.Command("pacman", "-Qkk", "--noconfirm")
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		// pacman exits non-zero when it detects modified files,
		// so the error is expected and the result is in stderr.
		_ = cmd.Run()

		buffer := stderr.String()

		if strings.TrimSpace(buffer) == "" {
			buffer = stdout.String()
		}

		lines := strings.Split(strings.TrimSpace(buffer), "\n")

		for l := 0; l < len(lines); l++ {

			name, path, issues := parseVerificationLine(lines[l])

			if name != "" && path != "" && len(issues) > 0 {

				verification, ok := verifications[name]

				if ok == false {
					tmp := structs.NewPackageVerification(name, "pacman")
					verification = &tmp
					verifications[name] = verification
				}

				verification.AddFile(path, issues, toRemediations(name, issues))

			}

		}

		for _, verification := range verifications {

			if verification.IsValid() {
				collected = append(collected, *verification)
			}

		}

	}

	return collected

}
