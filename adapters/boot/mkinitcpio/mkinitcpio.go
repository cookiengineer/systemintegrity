package mkinitcpio

import "os"

var SUPPORTED bool = false

func init() {

	_, err := os.Stat("/etc/mkinitcpio.conf")

	if err == nil {
		SUPPORTED = true
	}

}
