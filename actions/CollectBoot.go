package actions

import "github.com/cookiengineer/systemintegrity/structs"
import "github.com/cookiengineer/systemintegrity/adapters/boot/bootctl"
import "github.com/cookiengineer/systemintegrity/adapters/boot/mkinitcpio"

func CollectBoot(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectBoot")

	if bootctl.SUPPORTED == true {

		boot := bootctl.CollectBoot()

		boot.SetKernel(system.Distribution.Kernel)
		boot.SetKernelArchitecture(system.Distribution.KernelArchitecture)
		boot.SetKernelVersion(system.Distribution.KernelVersion)

		if mkinitcpio.SUPPORTED == true {
			boot.SetInitramfs(mkinitcpio.CollectInitramfs())
		}

		if boot.IsValid() {

			console.Info("bootctl.CollectBoot(): Found Boot Configuration")
			system.SetBoot(boot)
			result = true

		} else {

			console.Warn("bootctl.CollectBoot(): Found 0 Boot Configurations")

		}

	} else {

		// Fallback to kernel information gathered from the Distribution
		boot := structs.NewBoot()

		boot.SetKernel(system.Distribution.Kernel)
		boot.SetKernelArchitecture(system.Distribution.KernelArchitecture)
		boot.SetKernelVersion(system.Distribution.KernelVersion)

		if boot.IsValid() {

			console.Info("CollectBoot(): Found Kernel Configuration")
			system.SetBoot(boot)
			result = true

		} else {

			console.Warn("CollectBoot(): Unsupported")

		}

	}

	console.GroupEnd("actions/CollectBoot")

	return result

}
