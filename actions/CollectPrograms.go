package actions

import "github.com/cookiengineer/systemintegrity/structs"
import "github.com/cookiengineer/systemintegrity/adapters/programs/ldd"
import "github.com/cookiengineer/systemintegrity/adapters/programs/npm"
import "github.com/cookiengineer/systemintegrity/adapters/programs/proc"
import "strconv"

func CollectPrograms(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectPrograms")

	collected := make([]structs.Program, 0)

	if proc.SUPPORTED == true {

		tmp := proc.CollectPrograms()

		if len(tmp) > 0 {

			console.Info("proc.CollectPrograms(): Found " + strconv.Itoa(len(tmp)) + " Programs")

			for _, program := range tmp {

				if program.IsProgram() {

					if ldd.SUPPORTED == true && ldd.IsProgram(program) {
						ldd.AssembleProgramFilesystem(&program)
					}

					if npm.SUPPORTED == true && npm.IsProgram(program) {
						npm.AssembleProgramPackages(&program)
						npm.AssembleProgramDependencies(&program)
					}

					collected = append(collected, program)

				}

			}

		} else {

			console.Warn("proc.CollectPrograms(): Found 0 Programs")

		}

		system.SetPrograms(collected)
		result = true

	} else {

		console.Warn("proc.CollectPrograms(): Unsupported")

	}

	console.Log("Collected " + strconv.Itoa(len(system.Programs)) + "/" + strconv.Itoa(len(collected)) + " Programs")
	console.GroupEnd("actions/CollectPrograms")

	return result

}
