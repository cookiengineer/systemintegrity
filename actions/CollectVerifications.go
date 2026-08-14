package actions

import "github.com/cookiengineer/systemintegrity/structs"
import "github.com/cookiengineer/systemintegrity/adapters/packages/apt"
import "github.com/cookiengineer/systemintegrity/adapters/packages/dnf"
import "github.com/cookiengineer/systemintegrity/adapters/packages/pacman"
import "github.com/cookiengineer/systemintegrity/adapters/packages/rpm"
import "github.com/cookiengineer/systemintegrity/adapters/packages/zypper"
import "strconv"

func linkVerifications(collected []structs.PackageVerification, system *structs.System) []structs.PackageVerification {

	packages := make(map[string]*structs.Package)

	for p := 0; p < len(system.Packages); p++ {
		packages[system.Packages[p].Name] = &system.Packages[p]
	}

	result := make([]structs.PackageVerification, 0)

	for _, verification := range collected {

		pkg, ok := packages[verification.Name]

		if ok == true {
			verification.SetVersion(pkg.Version.String())
			verification.SetManager(pkg.Manager.String())
		}

		if verification.IsValid() {
			result = append(result, verification)
		}

	}

	return result

}

func CollectVerifications(console *structs.Console, system *structs.System) bool {

	var result bool

	console.Group("actions/CollectVerifications")

	collected := make([]structs.PackageVerification, 0)

	if pacman.SUPPORTED == true {

		// antergos
		// archlinux
		// manjaro

		tmp := pacman.CollectVerification()

		if len(tmp) > 0 {

			console.Info("pacman.CollectVerification(): Found " + strconv.Itoa(len(tmp)) + " affected Packages")

			collected = append(collected, linkVerifications(tmp, system)...)

		} else {

			console.Info("pacman.CollectVerification(): Found 0 affected Packages")

		}

		system.SetVerifications(collected)
		result = true

	} else if apt.SUPPORTED == true {

		// debian
		// ubuntu
		// linuxmint
		// trisquel

		tmp := apt.CollectVerification()

		if len(tmp) > 0 {

			console.Info("apt.CollectVerification(): Found " + strconv.Itoa(len(tmp)) + " affected Packages")

			collected = append(collected, linkVerifications(tmp, system)...)

		} else {

			console.Info("apt.CollectVerification(): Found 0 affected Packages")

		}

		system.SetVerifications(collected)
		result = true

	} else if rpm.SUPPORTED == true {

		// redhat
		// centos
		// oraclelinux
		// almalinux
		// rockylinux
		// fedora
		// amazonlinux
		// opensuse
		// suse-desktop
		// suse-server

		tmp := rpm.CollectVerification()

		if len(tmp) > 0 {

			console.Info("rpm.CollectVerification(): Found " + strconv.Itoa(len(tmp)) + " affected Packages")

			collected = append(collected, linkVerifications(tmp, system)...)

		} else {

			console.Info("rpm.CollectVerification(): Found 0 affected Packages")

		}

		system.SetVerifications(collected)
		result = true

	} else if dnf.SUPPORTED == true {

		tmp := dnf.CollectVerification()

		if len(tmp) > 0 {

			console.Info("dnf.CollectVerification(): Found " + strconv.Itoa(len(tmp)) + " affected Packages")

			collected = append(collected, linkVerifications(tmp, system)...)

		} else {

			console.Info("dnf.CollectVerification(): Found 0 affected Packages")

		}

		system.SetVerifications(collected)
		result = true

	} else if zypper.SUPPORTED == true {

		tmp := zypper.CollectVerification()

		if len(tmp) > 0 {

			console.Info("zypper.CollectVerification(): Found " + strconv.Itoa(len(tmp)) + " affected Packages")

			collected = append(collected, linkVerifications(tmp, system)...)

		} else {

			console.Info("zypper.CollectVerification(): Found 0 affected Packages")

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
