package structs

import "github.com/cookiengineer/systemintegrity/types"
import "strings"

type FileVerification struct {
	Path         string                        `json:"path"`
	Issues       []types.PackageVerificationIssue `json:"issues"`
	Remediations []types.Remediation           `json:"remediations,omitempty"`
}

type PackageVerification struct {
	Name    string             `json:"name"`
	Version types.Version      `json:"version"`
	Manager types.Manager      `json:"manager"`
	Files   []FileVerification `json:"files"`
}

func NewPackageVerification(name string, manager string) PackageVerification {

	var verification PackageVerification

	verification.SetName(name)
	verification.SetManager(manager)
	verification.Files = make([]FileVerification, 0)

	return verification

}

func (verification *PackageVerification) IsIdentical(value PackageVerification) bool {

	var result bool

	if verification.Name == value.Name &&
		verification.Version.String() == value.Version.String() &&
		verification.Manager.String() == value.Manager.String() {
		result = true
	}

	return result

}

func (verification *PackageVerification) IsValid() bool {

	var result bool

	if verification.Name != "" && verification.Manager.IsValid() {

		result = true

		for f := 0; f < len(verification.Files); f++ {

			if verification.Files[f].Path == "" || len(verification.Files[f].Issues) == 0 {
				result = false
				break
			}

		}

	}

	return result

}

func (verification *PackageVerification) AddFile(path string, issues []types.PackageVerificationIssue, remediations []types.Remediation) {

	path = strings.TrimSpace(path)

	file := FileVerification{
		Path:         path,
		Issues:       make([]types.PackageVerificationIssue, 0),
		Remediations: make([]types.Remediation, 0),
	}

	for i := 0; i < len(issues); i++ {

		if issues[i].IsValid() {
			file.Issues = append(file.Issues, issues[i])
		}

	}

	for r := 0; r < len(remediations); r++ {

		if remediations[r].IsValid() {
			file.Remediations = append(file.Remediations, remediations[r])
		}

	}

	if path != "" && len(file.Issues) > 0 {

		found := false

		for f := 0; f < len(verification.Files); f++ {

			if verification.Files[f].IsIdentical(file) {
				found = true
				break
			}

		}

		if found == false {
			verification.Files = append(verification.Files, file)
		}

	}

}

func (verification *PackageVerification) SetFiles(files []FileVerification) {

	filtered := make([]FileVerification, 0)

	for f := 0; f < len(files); f++ {

		file := files[f]

		if file.Path != "" && len(file.Issues) > 0 {

			found := false

			for p := 0; p < len(filtered); p++ {

				if filtered[p].IsIdentical(file) {
					found = true
					break
				}

			}

			if found == false {
				filtered = append(filtered, file)
			}

		}

	}

	verification.Files = filtered

}

func (verification *PackageVerification) SetManager(value string) {

	manager := types.ParseManager(value)

	if manager != nil {
		verification.Manager = *manager
	}

}

func (verification *PackageVerification) SetName(value string) {
	verification.Name = strings.TrimSpace(value)
}

func (verification *PackageVerification) SetVersion(value string) {

	version := types.ToVersion(value)

	if version.IsValid() {
		verification.Version = version
	}

}

func (file *FileVerification) IsIdentical(value FileVerification) bool {

	var result bool

	if file.Path == value.Path && len(file.Issues) == len(value.Issues) {

		result = true

		for i := 0; i < len(file.Issues); i++ {

			if file.Issues[i].String() != value.Issues[i].String() {
				result = false
				break
			}

		}

	}

	return result

}
