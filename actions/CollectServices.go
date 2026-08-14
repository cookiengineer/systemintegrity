package actions

import "github.com/cookiengineer/systemintegrity/structs"
import "github.com/cookiengineer/systemintegrity/adapters/programs/proc"
import "strconv"

func CollectServices(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectServices")

	collected := make([]structs.Program, 0)

	if proc.SUPPORTED == true {

		tmp := proc.CollectServices()

		if len(tmp) > 0 {

			console.Info("proc.CollectServices(): Found " + strconv.Itoa(len(tmp)) + " Services")

			for _, program := range tmp {

				if program.IsService() {
					collected = append(collected, program)
				}

			}

		} else {

			console.Warn("proc.CollectServices(): Found 0 Services")

		}

		system.SetServices(collected)
		result = true

	} else {

		console.Warn("proc.CollectServices(): Unsupported")

	}

	console.Log("Collected " + strconv.Itoa(len(system.Services)) + "/" + strconv.Itoa(len(collected)) + " Services")
	console.GroupEnd("actions/CollectServices")

	return result

}
