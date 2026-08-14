package zypper

import "github.com/cookiengineer/systemintegrity/structs"
import "github.com/cookiengineer/systemintegrity/adapters/packages/rpm"

func CollectVerification() []structs.PackageVerification {
	return rpm.CollectVerificationFor("zypper")
}
