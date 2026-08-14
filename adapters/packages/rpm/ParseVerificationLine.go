package rpm

import "github.com/cookiengineer/systemintegrity/types"
import "strings"

func ParseVerificationLine(line string) (string, []types.PackageVerificationIssue) {

	line = strings.TrimSpace(line)

	if line == "" {
		return "", nil
	}

	issues := make([]types.PackageVerificationIssue, 0)

	// "missing [attr] pathname"
	if strings.HasPrefix(line, "missing") {

		rest := strings.TrimSpace(line[len("missing"):])

		if strings.HasSuffix(rest, ")") {

			if idx := strings.LastIndex(rest, " ("); idx > 0 {
				rest = strings.TrimSpace(rest[0:idx])
			}

		}

		rest = strings.TrimSpace(strings.TrimPrefix(rest, "c "))

		issues = append(issues, types.PackageVerificationIssueMissingFile)

		return rest, issues

	}

	// check-string form: "<flags> [attr] <pathname>"
	idx := strings.Index(line, " /")

	if idx == -1 {
		return "", nil
	}

	flags := strings.TrimSpace(line[0:idx])
	path := strings.TrimSpace(line[idx+1:])

	for i := 0; i < len(flags); i++ {

		character := flags[i]

		var issue types.PackageVerificationIssue

		if character == 'S' {
			issue = types.PackageVerificationIssueSizeMismatch
		} else if character == 'M' {
			issue = types.PackageVerificationIssueModeMismatch
		} else if character == '5' {
			issue = types.PackageVerificationIssueChecksumMismatch
		} else if character == 'D' {
			issue = types.PackageVerificationIssueDeviceMismatch
		} else if character == 'L' {
			issue = types.PackageVerificationIssueReadlinkMismatch
		} else if character == 'U' {
			issue = types.PackageVerificationIssueUserMismatch
		} else if character == 'G' {
			issue = types.PackageVerificationIssueGroupMismatch
		} else if character == 'T' {
			issue = types.PackageVerificationIssueModificationTimeMismatch
		} else if character == 'P' {
			issue = types.PackageVerificationIssueCapabilitiesMismatch
		} else if character == '?' {
			issue = types.PackageVerificationIssueChecksumUnavailable
		}

		if issue.IsValid() {

			found := false

			for e := 0; e < len(issues); e++ {

				if issues[e].String() == issue.String() {
					found = true
					break
				}

			}

			if found == false {
				issues = append(issues, issue)
			}

		}

	}

	return path, issues

}

func ToRemediations(manager string, name string, issues []types.PackageVerificationIssue) []types.Remediation {

	remediations := make([]types.Remediation, 0)

	for i := 0; i < len(issues); i++ {

		issue := issues[i]

		if issue.IsIssue() == false {
			continue
		}

		var command string

		if issue.String() == string(types.PackageVerificationIssueUserMismatch) ||
			issue.String() == string(types.PackageVerificationIssueGroupMismatch) {
			command = "rpm --setugids " + name
		} else if issue.String() == string(types.PackageVerificationIssueModeMismatch) {
			command = "rpm --setperms " + name
		} else if issue.String() == string(types.PackageVerificationIssuePermissionDenied) {
			command = "run as root"
		} else if manager == "dnf" {
			command = "dnf reinstall " + name
		} else if manager == "zypper" {
			command = "zypper install --force " + name
		} else {
			command = "rpm -U --replacepkgs --replacefiles " + name
		}

		remediation := types.NewRemediation(manager, issue, command)

		if remediation.IsValid() {
			remediations = append(remediations, remediation)
		}

	}

	return remediations

}
