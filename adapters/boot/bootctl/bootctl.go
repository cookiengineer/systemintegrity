package bootctl

import "os"

var SUPPORTED bool = false

func init() {

	_, err1 := os.Stat("/usr/bin/bootctl")
	_, err2 := os.Stat("/bin/bootctl")

	if err1 == nil || err2 == nil {
		SUPPORTED = true
	}

}
