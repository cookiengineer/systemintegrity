package countries

import "encoding/json"
import _ "embed"
import "log"

//go:embed Countries.json
var embedded_Countries []byte

func init() {

	var err error = nil

	err1 := json.Unmarshal(embedded_Countries, &Countries)

	if err1 != nil {
		err = err1
	}

	if err != nil {
		log.Printf("Cannot decompress embedded Countries.json: %s", err.Error())
	}

}
