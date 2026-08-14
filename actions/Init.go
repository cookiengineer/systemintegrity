package actions

import "github.com/cookiengineer/systemintegrity/insights/countries"
import "github.com/cookiengineer/systemintegrity/matchers"
import "github.com/cookiengineer/systemintegrity/structs"
import "github.com/cookiengineer/systemintegrity/adapters/system/coreutils"
import "github.com/cookiengineer/systemintegrity/adapters/system/etc"
import "github.com/cookiengineer/systemintegrity/adapters/system/proc"
import "github.com/cookiengineer/systemintegrity/adapters/system/sys"

func Init(console *structs.Console) *structs.System {

	console.Group("actions/Init")

	var system structs.System = structs.NewSystem()

	if etc.SUPPORTED == true {
		system = etc.CollectSystem()
		console.Log("System Name: \"" + system.Name + "\"")
	}

	if coreutils.SUPPORTED == true {

		kernel := coreutils.CollectKernel()
		kernel_architecture := coreutils.CollectKernelArchitecture()
		kernel_modules := proc.CollectKernelModules()
		kernel_version := coreutils.CollectKernelVersion()

		if kernel != "" {
			system.Distribution.SetKernel(kernel)
		}

		if kernel_architecture != "" {
			system.Distribution.SetKernelArchitecture(kernel_architecture)
		}

		if len(kernel_modules) > 0 {
			system.Distribution.SetKernelModules(kernel_modules)
		}

		if kernel_version != "" {
			system.Distribution.SetKernelVersion(kernel_version)
		}

	}

	if sys.SUPPORTED == true {
		system.SetBIOS(sys.CollectBIOS())
		system.SetBoard(sys.CollectBoard())
	}

	if system.Fingerprint.Timezone != "" {

		candidates := countries.Countries.Query(matchers.Country{
			Name: "any",
			Continent: "any",
			Allegiance: "any",
			Subnet: "any",
			Timezone: system.Fingerprint.Timezone,
		})

		if len(candidates) > 0 {
			system.SetCountry(candidates[0].ISO)
		} else {
			system.SetCountry("??")
		}

	}

	console.GroupEnd("actions/Init")

	return &system

}
