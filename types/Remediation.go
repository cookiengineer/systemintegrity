package types

import "strings"

type Remediation struct {
	Manager     Manager                  `json:"manager"`
	Issue       PackageVerificationIssue `json:"issue"`
	Command     string                   `json:"command"`
	Description string                   `json:"description,omitempty"`
}

func NewRemediation(manager string, issue PackageVerificationIssue, command string) Remediation {

	var remediation Remediation

	remediation.SetManager(manager)
	remediation.Issue = issue
	remediation.Command = strings.TrimSpace(command)

	return remediation

}

func (remediation *Remediation) IsValid() bool {

	var result bool

	if remediation.Manager.IsValid() && remediation.Issue.IsValid() && remediation.Command != "" {
		result = true
	}

	return result

}

func (remediation *Remediation) SetManager(value string) {

	manager := ParseManager(value)

	if manager != nil {
		remediation.Manager = *manager
	}

}
