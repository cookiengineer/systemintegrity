package devices

import "encoding/json"
import _ "embed"
import "log"

//go:embed Devices.json
var embedded_Devices []byte

func init() {

	var err error = nil

	err1 := json.Unmarshal(embedded_Devices, &Devices)

	if err1 != nil {
		err = err1
	}

	if err != nil {
		log.Printf("Cannot decompress embedded Devices.json: %s", err.Error())
	}

}
