package actions

import "github.com/cookiengineer/systemintegrity/structs"
import "github.com/cookiengineer/systemintegrity/adapters/packages/pacman"
import "strconv"

func CollectVerifications(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectVerifications")

	collected := make([]structs.PackageVerification, 0)

	if pacman.SUPPORTED == true {

		tmp := pacman.CollectVerification()

		if len(tmp) > 0 {

			console.Info("pacman.CollectVerification(): Found " + strconv.Itoa(len(tmp)) + " affected Packages")

			packages := make(map[string]*structs.Package)

			for p := 0; p < len(system.Packages); p++ {
				packages[system.Packages[p].Name] = &system.Packages[p]
			}

			for _, verification := range tmp {

				pkg, ok := packages[verification.Name]

				if ok == true {
					verification.SetVersion(pkg.Version.String())
					verification.SetManager(pkg.Manager.String())
				}

				if verification.IsValid() {
					collected = append(collected, verification)
				}

			}

		} else {

			console.Info("pacman.CollectVerification(): Found 0 affected Packages")

		}

		system.SetVerifications(collected)
		result = true

	} else {

		console.Warn("CollectVerifications(): Unsupported")

	}

	console.Log("Collected " + strconv.Itoa(len(system.Verifications)) + "/" + strconv.Itoa(len(collected)) + " affected Packages")
	console.GroupEnd("actions/CollectVerifications")

	return result

}
