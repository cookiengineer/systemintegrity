package sys

import "github.com/cookiengineer/systemintegrity/insights/devices"
import "github.com/cookiengineer/systemintegrity/matchers"

func toPCISystemName(id_vendor string, id_device string) string {

	var result string

	candidates := devices.Devices.Query(matchers.Device{
		Name: "any",
		Bus: "pci",
		Vendor: id_vendor,
		Device: id_device,
	})

	if len(candidates) > 0 {

		for _, device := range candidates {

			if device.System != nil {
				result = device.System.Name
				break
			}

		}

	}

	return result

}

