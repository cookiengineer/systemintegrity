package bootctl

import "github.com/cookiengineer/systemintegrity/structs"
import "os"
import "os/exec"
import "strings"

func CollectBoot() structs.Boot {

	var boot structs.Boot = structs.NewBoot()

	// Firmware mode detection via /sys/firmware/efi
	if _, err := os.Stat("/sys/firmware/efi"); err == nil {
		boot.SetMode("uefi")
	} else {
		boot.SetMode("bios")
	}

	// Secure Boot state via efivars
	entries, err := os.ReadDir("/sys/firmware/efi/efivars")

	if err == nil {

		for e := 0; e < len(entries); e++ {

			name := entries[e].Name()

			if strings.HasPrefix(name, "SecureBoot-") {

				buffer, err := os.ReadFile("/sys/firmware/efi/efivars/" + name)

				// efivar layout: 4 bytes attributes followed by data
				if err == nil && len(buffer) >= 5 {

					if buffer[4] == 0x01 {
						boot.SetSecureBoot("enabled")
					} else {
						boot.SetSecureBoot("disabled")
					}

				}

				break

			}

		}

	}

	// Bootloader detection
	if SUPPORTED == true {

		cmd := exec.Command("bootctl", "status")
		buffer, err := cmd.Output()

		if err == nil {

			if strings.Contains(strings.ToLower(string(buffer)), "systemd-boot") {
				boot.SetBootloader("systemd-boot")
			}

		}

	}

	if boot.Bootloader == "" {

		if _, err := os.Stat("/boot/grub/grub.cfg"); err == nil {
			boot.SetBootloader("grub")
		} else if _, err := os.Stat("/boot/grub2/grub.cfg"); err == nil {
			boot.SetBootloader("grub")
		}

	}

	// EFI System Partition detection
	if boot.Mode == "uefi" {

		candidates := []string{"/efi/EFI", "/boot/efi/EFI", "/boot/EFI"}

		for c := 0; c < len(candidates); c++ {

			if _, err := os.Stat(candidates[c]); err == nil {

				if strings.HasSuffix(candidates[c], "/EFI") {
					boot.SetESP(strings.TrimSuffix(candidates[c], "/EFI"))
				}

				break

			}

		}

	}

	return boot

}
