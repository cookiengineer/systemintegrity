package matchers

import "github.com/cookiengineer/systemintegrity/types"
import "encoding/hex"
import "strings"

type Subnet struct {
	Name   string `json:"name"`
	Subnet string `json:"subnet"`
}

func NewSubnet() Subnet {

	var subnet Subnet

	subnet.Name = "any"
	subnet.Subnet = "any"

	return subnet

}

func ToSubnet(value string) Subnet {

	var subnet Subnet

	if strings.Contains(value, "/") {
		subnet.Name = "any"
		subnet.SetSubnet(value)
	} else {
		subnet.SetName(value)
		subnet.Subnet = "any"
	}

	return subnet

}

func (subnet *Subnet) IsIdentical(value Subnet) bool {

	var result bool

	if subnet.Name == value.Name && subnet.Subnet == value.Subnet {
		result = true
	}

	return result

}

func (subnet *Subnet) IsValid() bool {

	var result bool

	if subnet.Name != "any" || subnet.Subnet != "any" {
		result = true
	}

	return result

}

func (subnet *Subnet) Matches(name string, subnet_ string) bool {
	return subnet.MatchesName(name) && subnet.MatchesSubnet(subnet_)
}

func (subnet *Subnet) MatchesName(value string) bool {

	var result bool

	if subnet.Name == value {
		result = true
	} else if subnet.Name == "any" {
		result = true
	}

	return result

}

func (subnet *Subnet) MatchesSubnet(value string) bool {

	var result bool

	if subnet.Subnet != "any" && value != "any" {
		result = containsSubnet(value, subnet.Subnet)
	} else if subnet.Subnet == "any" {
		result = true
	}

	return result

}

func (subnet *Subnet) SetName(value string) {
	subnet.Name = strings.TrimSpace(value)
}

func (subnet *Subnet) SetSubnet(value string) {

	address, prefix := toSubnet(value)

	if value == "all" || value == "any" || value == "*" {
		subnet.Subnet = "any"
	} else if address != "" && prefix != 0 {
		subnet.Subnet = value
	}

}

func (subnet *Subnet) Hash() string {

	var hash string

	if subnet.Name != "any" {

		hash = subnet.Name

	} else if subnet.Subnet != "any" {

		address, prefix := toSubnet(subnet.Subnet)

		if types.IsIPv6(address) {

			ipv6 := types.ParseIPv6(address)

			if ipv6 != nil {
				bytes := ipv6.Bytes(prefix)
				hash = hex.EncodeToString(bytes)
			}

		} else if types.IsIPv4(address) {

			ipv4 := types.ParseIPv4(address)

			if ipv4 != nil {
				bytes := ipv4.Bytes(prefix)
				hash = hex.EncodeToString(bytes)
			}

		}

	}

	return hash

}
