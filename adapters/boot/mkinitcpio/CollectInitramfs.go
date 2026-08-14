package mkinitcpio

import "os"
import "path/filepath"
import "strings"

func CollectInitramfs() string {

	var result string

	if SUPPORTED == true {

		entries, err := os.ReadDir("/etc/mkinitcpio.d")

		if err == nil {

			for e := 0; e < len(entries); e++ {

				name := entries[e].Name()

				if strings.HasSuffix(name, ".preset") {
					result = strings.TrimSuffix(name, filepath.Ext(name))
					break
				}

			}

		}

	}

	return result

}
