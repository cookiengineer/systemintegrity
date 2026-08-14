package structs

import "github.com/cookiengineer/systemintegrity/types"
import "strings"

type FileVerification struct {
	Path   string `json:"path"`
	Reason string `json:"reason"`
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

			if verification.Files[f].Path == "" || verification.Files[f].Reason == "" {
				result = false
				break
			}

		}

	}

	return result

}

func (verification *PackageVerification) AddFile(path string, reason string) {

	path = strings.TrimSpace(path)
	reason = strings.TrimSpace(reason)

	if path != "" && reason != "" {

		found := false

		for f := 0; f < len(verification.Files); f++ {

			if verification.Files[f].Path == path && verification.Files[f].Reason == reason {
				found = true
				break
			}

		}

		if found == false {
			verification.Files = append(verification.Files, FileVerification{
				Path:   path,
				Reason: reason,
			})
		}

	}

}

func (verification *PackageVerification) SetFiles(files []FileVerification) {

	filtered := make([]FileVerification, 0)

	for f := 0; f < len(files); f++ {

		file := files[f]

		if file.Path != "" && file.Reason != "" {

			found := false

			for p := 0; p < len(filtered); p++ {

				if filtered[p].Path == file.Path && filtered[p].Reason == file.Reason {
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
