package apt

import "github.com/cookiengineer/systemintegrity/structs"
import "os"
import "os/exec"
import "strings"

func listPackageNames() []string {

	names := make([]string, 0)

	if SUPPORTED == true {

		cmd := exec.Command("dpkg-query", "-W", "-f", "${binary:Package}\\n")
		buffer, err := cmd.Output()

		if err == nil && len(buffer) > 0 {

			lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

			for l := 0; l < len(lines); l++ {

				name := strings.TrimSpace(lines[l])

				if name != "" {
					names = append(names, name)
				}

			}

		}

	}

	return names

}

func CollectVerification() []structs.PackageVerification {

	var collected []structs.PackageVerification

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		verifications := make(map[string]*structs.PackageVerification)

		names := listPackageNames()

		for n := 0; n < len(names); n++ {

			name := names[n]

			cmd := exec.Command("dpkg", "--verify", name)
			buffer, err := cmd.Output()

			// dpkg --verify can exit non-zero for some packages; the
			// mismatch report is still in stdout.
			_ = err

			if len(buffer) > 0 {

				lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

				for l := 0; l < len(lines); l++ {

					path, issues := ParseVerificationLine(lines[l])

					if path != "" && len(issues) > 0 {

						verification, ok := verifications[name]

						if ok == false {
							tmp := structs.NewPackageVerification(name, "apt")
							verification = &tmp
							verifications[name] = verification
						}

						verification.AddFile(path, issues, toRemediations(name, issues))

					}

				}

			}

		}

		for _, verification := range verifications {

			if verification.IsValid() {
				collected = append(collected, *verification)
			}

		}

	}

	return collected

}
