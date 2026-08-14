package types

import "strconv"

type PackageVerificationIssue string

const (
	PackageVerificationIssueChecksumMismatch         PackageVerificationIssue = "checksum_mismatch"
	PackageVerificationIssueChecksumUnavailable      PackageVerificationIssue = "checksum_unavailable"
	PackageVerificationIssueDeviceMismatch           PackageVerificationIssue = "device_mismatch"
	PackageVerificationIssueFileTypeMismatch         PackageVerificationIssue = "file_type_mismatch"
	PackageVerificationIssueGroupMismatch            PackageVerificationIssue = "group_mismatch"
	PackageVerificationIssueModeMismatch             PackageVerificationIssue = "mode_mismatch"
	PackageVerificationIssueModificationTimeMismatch PackageVerificationIssue = "modification_time_mismatch"
	PackageVerificationIssueMissingFile              PackageVerificationIssue = "missing_file"
	PackageVerificationIssuePermissionDenied         PackageVerificationIssue = "permission_denied"
	PackageVerificationIssueReadlinkMismatch         PackageVerificationIssue = "readlink_mismatch"
	PackageVerificationIssueSizeMismatch             PackageVerificationIssue = "size_mismatch"
	PackageVerificationIssueUserMismatch             PackageVerificationIssue = "user_mismatch"
	PackageVerificationIssueCapabilitiesMismatch     PackageVerificationIssue = "capabilities_mismatch"
)

func IsPackageVerificationIssue(value string) bool {

	var result bool

	if value == string(PackageVerificationIssueChecksumMismatch) {
		result = true
	} else if value == string(PackageVerificationIssueChecksumUnavailable) {
		result = true
	} else if value == string(PackageVerificationIssueDeviceMismatch) {
		result = true
	} else if value == string(PackageVerificationIssueFileTypeMismatch) {
		result = true
	} else if value == string(PackageVerificationIssueGroupMismatch) {
		result = true
	} else if value == string(PackageVerificationIssueModeMismatch) {
		result = true
	} else if value == string(PackageVerificationIssueModificationTimeMismatch) {
		result = true
	} else if value == string(PackageVerificationIssueMissingFile) {
		result = true
	} else if value == string(PackageVerificationIssuePermissionDenied) {
		result = true
	} else if value == string(PackageVerificationIssueReadlinkMismatch) {
		result = true
	} else if value == string(PackageVerificationIssueSizeMismatch) {
		result = true
	} else if value == string(PackageVerificationIssueUserMismatch) {
		result = true
	} else if value == string(PackageVerificationIssueCapabilitiesMismatch) {
		result = true
	}

	return result

}

func ParsePackageVerificationIssue(value string) *PackageVerificationIssue {

	var result *PackageVerificationIssue = nil

	if IsPackageVerificationIssue(value) {
		issue := PackageVerificationIssue(value)
		result = &issue
	}

	return result

}

func ToPackageVerificationIssue(value string) PackageVerificationIssue {

	var issue PackageVerificationIssue

	tmp := ParsePackageVerificationIssue(value)

	if tmp != nil {
		issue = *tmp
	}

	return issue

}

func (issue PackageVerificationIssue) String() string {
	return string(issue)
}

func (issue PackageVerificationIssue) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(string(issue))), nil
}

func (issue *PackageVerificationIssue) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))

	if err != nil {
		return err
	}

	tmp := ParsePackageVerificationIssue(unquoted)

	if tmp != nil {
		*issue = *tmp
	}

	return nil

}

func (issue *PackageVerificationIssue) IsValid() bool {

	var result bool

	if IsPackageVerificationIssue(issue.String()) {
		result = true
	}

	return result

}

// IsIssue reports whether the mismatch is a genuine integrity problem.
// Modification time mismatches are expected noise (e.g. touched files) and
// are therefore not treated as issues.
func (issue *PackageVerificationIssue) IsIssue() bool {

	var result bool = true

	if issue.String() == string(PackageVerificationIssueModificationTimeMismatch) {
		result = false
	}

	return result

}

func (issue *PackageVerificationIssue) Description() string {

	var description string

	if issue.String() == string(PackageVerificationIssueChecksumMismatch) {
		description = "Checksum mismatch"
	} else if issue.String() == string(PackageVerificationIssueChecksumUnavailable) {
		description = "Checksum unavailable"
	} else if issue.String() == string(PackageVerificationIssueDeviceMismatch) {
		description = "Device mismatch"
	} else if issue.String() == string(PackageVerificationIssueFileTypeMismatch) {
		description = "File type mismatch"
	} else if issue.String() == string(PackageVerificationIssueGroupMismatch) {
		description = "Group ownership mismatch"
	} else if issue.String() == string(PackageVerificationIssueModeMismatch) {
		description = "Permissions mismatch"
	} else if issue.String() == string(PackageVerificationIssueModificationTimeMismatch) {
		description = "Modification time mismatch"
	} else if issue.String() == string(PackageVerificationIssueMissingFile) {
		description = "Missing file"
	} else if issue.String() == string(PackageVerificationIssuePermissionDenied) {
		description = "Permission denied"
	} else if issue.String() == string(PackageVerificationIssueReadlinkMismatch) {
		description = "Symlink path mismatch"
	} else if issue.String() == string(PackageVerificationIssueSizeMismatch) {
		description = "Size mismatch"
	} else if issue.String() == string(PackageVerificationIssueUserMismatch) {
		description = "User ownership mismatch"
	} else if issue.String() == string(PackageVerificationIssueCapabilitiesMismatch) {
		description = "Capabilities mismatch"
	}

	return description

}
