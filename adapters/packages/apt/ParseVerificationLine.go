package apt

import "github.com/cookiengineer/systemintegrity/types"
import "strings"

func ParseVerificationLine(line string) (string, []types.PackageVerificationIssue) {

	line = strings.TrimSpace(line)

	if line == "" {
		return "", nil
	}

	issues := make([]types.PackageVerificationIssue, 0)

	// "missing [c] pathname [(error-message)]"
	if strings.HasPrefix(line, "missing") {

		rest := strings.TrimSpace(line[len("missing"):])

		if strings.HasSuffix(rest, ")") {

			if idx := strings.LastIndex(rest, " ("); idx > 0 {
				rest = strings.TrimSpace(rest[0:idx])
			}

		}

		// optional conffile attribute "c"
		rest = strings.TrimSpace(strings.TrimPrefix(rest, "c "))

		issues = append(issues, types.PackageVerificationIssueMissingFile)

		return rest, issues

	}

	// check-string form: "<flags> [c] <pathname>"
	idx := strings.Index(line, " /")

	if idx == -1 {
		return "", nil
	}

	flags := strings.TrimSpace(line[0:idx])
	path := strings.TrimSpace(line[idx+1:])

	// Only positions 2 (mode, 'M') and 3 (digest, '5') report failures.
	// Positions 1 and 4-9 are always '?' (unsupported) and must be ignored.
	//
	// dpkg does not track pathname metadata, so position 2 ('M') is only a
	// heuristic: it fails when a path with a known digest is no longer a
	// regular file (e.g. replaced by a symlink). It is NOT a permission-bits
	// check like pacman's "Permissions mismatch" or rpm's 'M'.
	for i := 0; i < len(flags); i++ {

		character := flags[i]

		if i == 1 && character == 'M' {
			issues = append(issues, types.PackageVerificationIssueFileTypeMismatch)
		} else if i == 2 && character == '5' {
			issues = append(issues, types.PackageVerificationIssueChecksumMismatch)
		}

	}

	return path, issues

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
			command = "apt-get install --reinstall " + name
		}

		remediation := types.NewRemediation("apt", issue, command)

		if remediation.IsValid() {
			remediations = append(remediations, remediation)
		}

	}

	return remediations

}
