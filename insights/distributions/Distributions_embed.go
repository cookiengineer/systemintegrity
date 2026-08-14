package distributions

import "encoding/json"
import _ "embed"
import "log"

//go:embed Distributions.json
var embedded_Distributions []byte

func init() {

	var err error = nil

	err1 := json.Unmarshal(embedded_Distributions, &Distributions)

	if err1 != nil {
		err = err1
	}

	if err != nil {
		log.Printf("Cannot decompress embedded Distributions.json: %s", err.Error())
	}

}
