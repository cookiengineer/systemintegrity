package apk

import "github.com/cookiengineer/systemintegrity/types"
import "strings"

type apkAuditDetail struct {
	mode string
	uid  string
	gid  string
}

func parseAuditDetail(raw string) apkAuditDetail {

	var detail apkAuditDetail

	fields := strings.Fields(raw)

	for f := 0; f < len(fields); f++ {

		field := fields[f]

		if strings.HasPrefix(field, "mode=") {
			detail.mode = strings.TrimPrefix(field, "mode=")
		} else if strings.HasPrefix(field, "uid=") {
			detail.uid = strings.TrimPrefix(field, "uid=")
		} else if strings.HasPrefix(field, "gid=") {
			detail.gid = strings.TrimPrefix(field, "gid=")
		}

	}

	return detail

}

// ParseVerification parses the output of "apk audit --system --check-permissions
// --details" and maps every changed path to its package-manager-agnostic issues.
//
// The output is line-oriented and stateful: each entry consists of a "-" record
// (expected metadata), an optional "+" record (actual metadata, only when the
// file changed) and a status line ("<status> <relative-path>") where status is:
//
//	U - file modified (checksum/hash mismatch)
//	X - file missing
//	M - file permissions (mode/uid/gid) differ
//	m - directory permissions differ (mount-point noise, ignored)
func ParseVerification(buffer string) map[string][]types.PackageVerificationIssue {

	result := make(map[string][]types.PackageVerificationIssue)

	lines := strings.Split(strings.TrimSpace(buffer), "\n")

	var expected apkAuditDetail
	var actual apkAuditDetail
	has_expected := false
	has_actual := false

	for l := 0; l < len(lines); l++ {

		line := strings.TrimSpace(lines[l])

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "- ") {

			expected = parseAuditDetail(strings.TrimSpace(line[2:]))
			has_expected = true
			has_actual = false

		} else if strings.HasPrefix(line, "+ ") {

			actual = parseAuditDetail(strings.TrimSpace(line[2:]))
			has_actual = true

		} else if len(line) >= 2 && line[1] == ' ' {

			status := line[0]
			path := strings.TrimSpace(line[2:])

			// Ignore directory entries (mount-point noise such as proc/, sys/).
			if path == "" || strings.HasSuffix(path, "/") {
				continue
			}

			issues := make([]types.PackageVerificationIssue, 0)

			if status == 'U' {
				issues = append(issues, types.PackageVerificationIssueChecksumMismatch)
			} else if status == 'X' {
				issues = append(issues, types.PackageVerificationIssueMissingFile)
			} else if status == 'M' || status == 'm' {

				if has_expected == true && has_actual == true {

					if expected.mode != actual.mode {
						issues = append(issues, types.PackageVerificationIssueModeMismatch)
					}

					if expected.uid != actual.uid {
						issues = append(issues, types.PackageVerificationIssueUserMismatch)
					}

					if expected.gid != actual.gid {
						issues = append(issues, types.PackageVerificationIssueGroupMismatch)
					}

				} else {
					issues = append(issues, types.PackageVerificationIssueModeMismatch)
				}

			}

			if len(issues) > 0 {
				result["/"+path] = issues
			}

		}

	}

	return result

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
			command = "apk fix " + name
		}

		remediation := types.NewRemediation("apk", issue, command)

		if remediation.IsValid() {
			remediations = append(remediations, remediation)
		}

	}

	return remediations

}
