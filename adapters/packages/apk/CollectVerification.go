package apk

import "github.com/cookiengineer/systemintegrity/structs"
import "os"
import "os/exec"
import "sort"
import "strings"

func packageOwner(path string) string {

	var name string

	cmd := exec.Command("apk", "info", "--who-owns", path)
	buffer, err := cmd.Output()

	if err == nil && len(buffer) > 0 {

		line := strings.TrimSpace(string(buffer))

		// "/usr/bin/ssh is owned by openssh-client-default-10.3_p1-r0"
		if strings.Contains(line, " is owned by ") {

			owner := strings.TrimSpace(line[strings.Index(line, " is owned by ")+len(" is owned by "):])

			name, _ = toNameAndVersion(owner)

		}

	}

	return name

}

func CollectVerification() []structs.PackageVerification {

	var collected []structs.PackageVerification

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		verifications := make(map[string]*structs.PackageVerification)

		cmd := exec.Command("apk", "audit", "--system", "--check-permissions", "--details")
		buffer, err := cmd.Output()

		// apk audit reports changed files on stdout; the exit code is not a
		// reliable indicator of whether changes were found.
		_ = err

		if len(buffer) > 0 {

			files := ParseVerification(string(buffer))

			paths := make([]string, 0)

			for path, _ := range files {
				paths = append(paths, path)
			}

			sort.Strings(paths)

			for p := 0; p < len(paths); p++ {

				path := paths[p]
				issues := files[path]

				name := packageOwner(path)

				if name == "" {
					continue
				}

				verification, ok := verifications[name]

				if ok == false {
					tmp := structs.NewPackageVerification(name, "apk")
					verification = &tmp
					verifications[name] = verification
				}

				verification.AddFile(path, issues, toRemediations(name, issues))

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
