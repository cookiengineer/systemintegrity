package proc

import "github.com/cookiengineer/systemintegrity/structs"

func CollectServices() []structs.Program {

	var collected []structs.Program

	if SUPPORTED == true {

		refresh()

		for _, program := range Programs.Programs {

			if program.IsService() {
				collected = append(collected, *program)
			}

		}

	}

	return collected

}
