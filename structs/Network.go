package structs

import "slices"
import "sort"
import "strings"

type Network struct {
	Name    string   `json:"name"`
	Subnets []Subnet `json:"subnets"`
}

func NewNetwork(name string) Network {

	var network Network

	network.SetName(name)
	network.Subnets = make([]Subnet, 0)

	return network

}

func (network *Network) IsValid() bool {

	if network.Name != "" {

		result := true

		for s := 0; s < len(network.Subnets); s++ {

			if network.Subnets[s].IsValid() == false {
				result = false
				break
			}

		}

		return result

	}

	return false

}

func (network *Network) SetName(value string) {
	network.Name = strings.TrimSpace(value)
}

func (network *Network) AddSubnet(value Subnet) {

	if value.IsValid() {

		found := false

		for s := 0; s < len(network.Subnets); s++ {

			if network.Subnets[s].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			network.Subnets = append(network.Subnets, value)
		}

	}

}

func (network *Network) RemoveSubnet(value Subnet) {

	index := -1

	for s := 0; s < len(network.Subnets); s++ {

		if network.Subnets[s].IsIdentical(value) {
			index = s
			break
		}

	}

	if index != -1 {
		network.Subnets = append(network.Subnets[:index], network.Subnets[index+1:]...)
	}

}

func (network *Network) SetSubnets(value []Subnet) {

	filtered := make([]Subnet, 0)

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	sort.Slice(filtered, func(a int, b int) bool {

		if filtered[a].Prefix == filtered[b].Prefix {

			hash_a := filtered[a].Hash()
			hash_b := filtered[b].Hash()

			return hash_a < hash_b

		} else {

			return filtered[a].Prefix < filtered[b].Prefix

		}

	})

	network.Subnets = filtered

}

func (network *Network) String() string {

	var result string

	if network.Name != "" {
		result = strings.TrimSpace(network.Name)
	} else {

		asns := make([]string, 0)
		subnets := make([]string, 0)

		for _, subnet := range network.Subnets {

			if strings.HasPrefix(subnet.Name, "AS") {

				if !slices.Contains(asns, subnet.Name) {
					asns = append(asns, subnet.Name)
				}

			} else {

				if !slices.Contains(subnets, subnet.String()) {
					subnets = append(subnets, subnet.String())
				}

			}

		}

		if len(asns) > 0 {
			sort.Strings(asns)
			result = strings.Join(asns, ",")
		} else if len(subnets) > 0 {
			sort.Strings(subnets)
			result = strings.Join(subnets, ",")
		}

	}

	return result

}
