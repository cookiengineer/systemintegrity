package matchers

import "github.com/cookiengineer/systemintegrity/types"
import "strings"

type Distribution struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Manager string `json:"manager"`
	Vendor  string `json:"vendor"`
}

func NewDistribution() Distribution {

	var distribution Distribution

	distribution.Name = "any"
	distribution.Version = "any"
	distribution.Manager = "any"
	distribution.Vendor = "any"

	return distribution

}

func ToDistribution(value string) Distribution {

	var distribution Distribution

	distribution.Name = "any"
	distribution.Version = "any"
	distribution.Manager = "any"
	distribution.Vendor = "any"

	distribution.SetName(value)

	return distribution

}

func (distribution *Distribution) IsIdentical(value Distribution) bool {

	var result bool

	if distribution.Name == value.Name &&
		distribution.Version == value.Version &&
		distribution.Manager == value.Manager &&
		distribution.Vendor == value.Vendor {
		result = true
	}

	return result

}

func (distribution *Distribution) IsValid() bool {

	var result bool

	if distribution.Name != "any" || distribution.Version != "any" || distribution.Manager != "any" || distribution.Vendor != "any" {
		result = true
	}

	return result

}

func (distribution *Distribution) Matches(name string, version string, manager string, vendor string) bool {
	return distribution.MatchesName(name) && distribution.MatchesVersion(version) && distribution.MatchesManager(manager) && distribution.MatchesVendor(vendor)
}

func (distribution *Distribution) MatchesManager(value string) bool {

	var result bool

	if distribution.Manager == value {
		result = true
	} else if distribution.Manager == "any" {
		result = true
	}

	return result

}

func (distribution *Distribution) MatchesName(value string) bool {

	var result bool

	if distribution.Name == value {
		result = true
	} else if distribution.Name == "any" {
		result = true
	} else if strings.HasSuffix(distribution.Name, "-*") {

		prefix := distribution.Name[0 : len(distribution.Name)-2]

		if strings.HasPrefix(value, prefix) {
			result = true
		}

	}

	return result

}

func (distribution *Distribution) MatchesVendor(value string) bool {

	var result bool

	if distribution.Vendor == value {
		result = true
	} else if distribution.Vendor == "any" {
		result = true
	}

	return result

}

func (distribution *Distribution) MatchesVersion(value string) bool {

	var result bool

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(value, " ") {
		value = strings.TrimSpace(value[strings.Index(value, " ")+1:])
	}

	if distribution.Version == "any" {

		result = true

	} else if strings.HasPrefix(distribution.Version, "<= ") {

		distro_version := types.ToVersion(distribution.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.IsSame(distro_version) {
			result = true
		} else if other_version.IsBefore(distro_version) {
			result = true
		}

	} else if strings.HasPrefix(distribution.Version, "< ") {

		distro_version := types.ToVersion(distribution.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsBefore(distro_version) {
			result = true
		}

	} else if strings.HasPrefix(distribution.Version, ">= ") {

		distro_version := types.ToVersion(distribution.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.IsSame(distro_version) {
			result = true
		} else if other_version.IsAfter(distro_version) {
			result = true
		}

	} else if strings.HasPrefix(distribution.Version, "> ") {

		distro_version := types.ToVersion(distribution.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsAfter(distro_version) {
			result = true
		}

	} else {

		distro_version := types.ToVersion(distribution.Version)
		other_version := types.ToVersion(value)

		if other_version.IsSame(distro_version) {
			result = true
		}

	}

	return result

}

func (distribution *Distribution) SetName(value string) {

	if value == "all" || value == "any" || value == "*" {
		distribution.Name = "any"
	} else if value != "" {
		distribution.Name = strings.TrimSpace(value)
	}

}

func (distribution *Distribution) SetVendor(value string) {
	distribution.Vendor = strings.TrimSpace(value)
}

func (distribution *Distribution) SetVersion(value string) {

	if value == "all" || value == "any" || value == "*" {
		distribution.Version = "any"
	} else if value != "" {
		distribution.Version = value
	}

}

func (distribution *Distribution) Hash() string {

	var hash string

	if distribution.Name != "" {

		hash = distribution.Name

		if distribution.Version != "any" {
			hash += "-" + distribution.Version
		}

		if distribution.Vendor != "any" {
			hash += "-" + distribution.Vendor
		}

	}

	return hash

}
