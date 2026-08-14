package actions

import "github.com/cookiengineer/systemintegrity/structs"
import "github.com/cookiengineer/systemintegrity/types"
import "testing"

type issueExpectation struct {
	Path  string
	Issue types.PackageVerificationIssue
}

func hasIssue(system *structs.System, path string, want types.PackageVerificationIssue) bool {

	for v := 0; v < len(system.Verifications); v++ {

		for f := 0; f < len(system.Verifications[v].Files); f++ {

			file := system.Verifications[v].Files[f]

			if file.Path == path {

				for i := 0; i < len(file.Issues); i++ {

					if file.Issues[i].String() == want.String() {
						return true
					}

				}

			}

		}

	}

	return false

}

func assertIssues(t *testing.T, system *structs.System, expectations []issueExpectation) {

	for e := 0; e < len(expectations); e++ {

		expectation := expectations[e]

		if hasIssue(system, expectation.Path, expectation.Issue) == false {
			t.Errorf("Expected %q to contain %q", expectation.Path, expectation.Issue.String())
		}

	}

}

func validateVerifications(t *testing.T, system *structs.System) {

	for v := 0; v < len(system.Verifications); v++ {

		verification := system.Verifications[v]

		if verification.Name == "" {
			t.Errorf("Expected verification to contain a Name")
		}

		if verification.Manager.IsValid() == false {
			t.Errorf("Expected verification %q to contain a valid Manager", verification.Name)
		}

		for f := 0; f < len(verification.Files); f++ {

			file := verification.Files[f]

			if file.Path == "" || len(file.Issues) == 0 {
				t.Errorf("Expected verification %q file to contain Path and Issues", verification.Name)
			}

			for i := 0; i < len(file.Issues); i++ {

				if file.Issues[i].IsValid() == false {
					t.Errorf("Expected issue %q to be valid", file.Issues[i].String())
				}

			}

			for r := 0; r < len(file.Remediations); r++ {

				if file.Remediations[r].IsValid() == false {
					t.Errorf("Expected remediation to be valid")
				}

			}

		}

	}

}
