package pacman

import "github.com/cookiengineer/systemintegrity/structs"
import "bytes"
import "os"
import "os/exec"
import "strings"

func parseVerificationLine(line string) (string, string, string) {

	var name string
	var path string
	var reason string

	line = strings.TrimSpace(line)

	if strings.HasPrefix(line, "warning: ") {
		line = strings.TrimSpace(line[9:])
	} else if strings.HasPrefix(line, "error: ") {
		line = strings.TrimSpace(line[7:])
	}

	// "pkgname: /path/to/file (reason)"
	if strings.Contains(line, ": ") {

		name = strings.TrimSpace(line[0:strings.Index(line, ": ")])
		rest := strings.TrimSpace(line[strings.Index(line, ": ")+2:])

		if strings.Contains(rest, " (") && strings.HasSuffix(rest, ")") {

			last := strings.LastIndex(rest, " (")

			if last > 0 {

				path = strings.TrimSpace(rest[0:last])
				reason = strings.TrimSpace(rest[last+2 : len(rest)-1])

			}

		}

	}

	return name, path, reason

}

func CollectVerification() []structs.PackageVerification {

	var collected []structs.PackageVerification

	if SUPPORTED == true {

		os.Setenv("TZ", "Europe/Greenwich")
		os.Setenv("LC_TIME", "en_US")

		verifications := make(map[string]*structs.PackageVerification)

		var stdout bytes.Buffer
		var stderr bytes.Buffer

		cmd := exec.Command("pacman", "-Qkk", "--noconfirm")
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		// pacman exits non-zero when it detects modified files,
		// so the error is expected and the result is in stderr.
		_ = cmd.Run()

		buffer := stderr.String()

		if strings.TrimSpace(buffer) == "" {
			buffer = stdout.String()
		}

		lines := strings.Split(strings.TrimSpace(buffer), "\n")

		for l := 0; l < len(lines); l++ {

			name, path, reason := parseVerificationLine(lines[l])

			if name != "" && path != "" && reason != "" {

				verification, ok := verifications[name]

				if ok == false {
					tmp := structs.NewPackageVerification(name, "pacman")
					verification = &tmp
					verifications[name] = verification
				}

				verification.AddFile(path, reason)

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
