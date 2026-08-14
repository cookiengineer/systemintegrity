package rpm

import "github.com/cookiengineer/systemintegrity/structs"
import "os"
import "os/exec"
import "strings"

func listPackageNames() []string {

	names := make([]string, 0)

	if SUPPORTED == true {

		cmd := exec.Command("rpm", "-qa", "--qf", "%{name}\\n", "--noplugins")
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

func CollectVerificationFor(manager string) []structs.PackageVerification {

	var collected []structs.PackageVerification

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		verifications := make(map[string]*structs.PackageVerification)

		names := listPackageNames()

		for n := 0; n < len(names); n++ {

			name := names[n]

			cmd := exec.Command("rpm", "-V", name, "--noplugins")
			buffer, err := cmd.Output()

			// rpm -V exits non-zero when it detects modified files, so the
			// error is expected and the result is in stdout.
			_ = err

			if len(buffer) > 0 {

				lines := strings.Split(strings.TrimSpace(string(buffer)), "\n")

				for l := 0; l < len(lines); l++ {

					path, issues := ParseVerificationLine(lines[l])

					if path != "" && len(issues) > 0 {

						verification, ok := verifications[name]

						if ok == false {
							tmp := structs.NewPackageVerification(name, manager)
							verification = &tmp
							verifications[name] = verification
						}

						verification.AddFile(path, issues, ToRemediations(manager, name, issues))

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

func CollectVerification() []structs.PackageVerification {
	return CollectVerificationFor("rpm")
}
