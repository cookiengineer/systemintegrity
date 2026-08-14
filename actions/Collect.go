package actions

import "github.com/cookiengineer/systemintegrity/structs"

func Collect(console *structs.Console, system *structs.System) bool {

	console.Group("actions/Collect")

	CollectBoot(console, system)
	CollectDrives(console, system)
	CollectDevices(console, system)
	CollectBoot(console, system)
	CollectNetworks(console, system)
	CollectPrograms(console, system)
	CollectServices(console, system)
	CollectPackages(console, system)
	CollectVerifications(console, system)
	CollectUpdates(console, system)
	CollectUsers(console, system)

	for p := 0; p < len(system.Programs); p++ {
		system.Programs[p].ResolveDependencies(system.Packages)
	}

	CollectAntiques(console, system)

	console.GroupEnd("actions/Collect")

	return true

}
