package actions

import "github.com/cookiengineer/systemintegrity/structs"
import "github.com/cookiengineer/systemintegrity/adapters/packages/apk"
import "github.com/cookiengineer/systemintegrity/adapters/packages/apt"
import "github.com/cookiengineer/systemintegrity/adapters/packages/dnf"
import "github.com/cookiengineer/systemintegrity/adapters/packages/pacman"
import "github.com/cookiengineer/systemintegrity/adapters/packages/tdnf"
import "github.com/cookiengineer/systemintegrity/adapters/packages/zypper"
import "strconv"

func CollectUpdates(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectUpdates")

	collected := make([]structs.Update, 0)

	if pacman.SUPPORTED == true {

		// antergos
		// archlinux
		// manjaro

		tmp := pacman.CollectUpdates()

		if len(tmp) > 0 {

			console.Info("pacman.CollectUpdate(): Found " + strconv.Itoa(len(tmp)) + " Updates")

			for _, update := range tmp {

				if update.IsValid() {
					collected = append(collected, update)
				}

			}

		} else {

			console.Warn("pacman.CollectUpdates(): Found 0 Updates")

		}

		system.SetUpdates(collected)
		result = true

	} else if apt.SUPPORTED == true {

		// debian
		// ubuntu
		// linuxmint

		tmp := apt.CollectUpdates()

		if len(tmp) > 0 {

			console.Info("apt.CollectUpdate(): Found " + strconv.Itoa(len(tmp)) + " Updates")

			for _, update := range tmp {

				if update.IsValid() {
					collected = append(collected, update)
				}

			}

		} else {

			console.Warn("apt.CollectUpdates(): Found 0 Updates")

		}

		system.SetUpdates(collected)
		result = true

	} else if apk.SUPPORTED == true {

		// alpinelinux

		tmp := apk.CollectUpdates()

		if len(tmp) > 0 {

			console.Info("apk.CollectUpdate(): Found " + strconv.Itoa(len(tmp)) + " Updates")

			for _, update := range tmp {

				if update.IsValid() {
					collected = append(collected, update)
				}

			}

		} else {

			console.Warn("apk.CollectUpdates(): Found 0 Updates")

		}

		system.SetUpdates(collected)
		result = true

	} else if dnf.SUPPORTED == true {

		// redhat
		// centos
		// oraclelinux
		// almalinux
		// rockylinux
		// fedora
		// amazonlinux

		tmp := dnf.CollectUpdates()

		if len(tmp) > 0 {

			console.Info("dnf.CollectUpdate(): Found " + strconv.Itoa(len(tmp)) + " Updates")

			for _, update := range tmp {

				if update.IsValid() {
					collected = append(collected, update)
				}

			}

		} else {

			console.Warn("dnf.CollectUpdates(): Found 0 Updates")

		}

		system.SetUpdates(collected)
		result = true

	} else if tdnf.SUPPORTED == true {

		// cblmariner
		// photonos

		// TODO: tdnf.CollectUpdates()

	} else if zypper.SUPPORTED == true {

		// opensuse
		// suse-desktop
		// suse-server

		tmp := zypper.CollectUpdates()

		if len(tmp) > 0 {

			console.Info("zypper.CollectUpdate(): Found " + strconv.Itoa(len(tmp)) + " Updates")

			for _, update := range tmp {

				if update.IsValid() {
					collected = append(collected, update)
				}

			}

		} else {

			console.Warn("zypper.CollectUpdates(): Found 0 Updates")

		}

		system.SetUpdates(collected)
		result = true

	} else {

		// TODO: pkg.CollectUpdates()

		console.Warn("CollectUpdates(): Unsupported")

	}

	console.Log("Collected " + strconv.Itoa(len(system.Updates)) + "/" + strconv.Itoa(len(collected)) + " Updates")
	console.GroupEnd("actions/CollectUpdates")

	return result

}
