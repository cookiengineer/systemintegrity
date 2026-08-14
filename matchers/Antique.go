package matchers

import "github.com/cookiengineer/systemintegrity/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Antique struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Manager      string `json:"manager"`
	Vendor       string `json:"vendor"`
	Service      string `json:"service"`
}

func NewAntique() Antique {

	var antique Antique

	antique.Name = "any"
	antique.Version = "any"
	antique.Architecture = "any"
	antique.Manager = "any"
	antique.Service = "any"
	antique.Vendor = "any"

	return antique

}

func ToAntique(value string) Antique {

	var antique Antique

	antique.Name = "any"
	antique.Version = "any"
	antique.Architecture = "any"
	antique.Manager = "any"
	antique.Vendor = "any"
	antique.Service = "any"

	antique.Parse(value)

	return antique

}

func (antique *Antique) IsIdentical(value Antique) bool {

	var result bool

	if antique.Name == value.Name &&
		antique.Version == value.Version &&
		antique.Architecture == value.Architecture &&
		antique.Manager == value.Manager &&
		antique.Vendor == value.Vendor &&
		antique.Service == value.Service {
		result = true
	}

	return result

}

func (antique *Antique) IsValid() bool {

	var result bool

	if antique.Name != "any" || antique.Version != "any" || antique.Architecture != "any" || antique.Manager != "any" || antique.Vendor != "any" || antique.Service != "any" {
		result = true
	}

	return result

}

func (antique *Antique) Matches(name string, version string, architecture string, manager string, vendor string, service string) bool {
	return antique.MatchesName(name) && antique.MatchesVersion(version) && antique.MatchesArchitecture(architecture) && antique.MatchesManager(manager) && antique.MatchesVendor(vendor) && antique.MatchesService(service)
}

func (antique *Antique) MatchesArchitecture(value string) bool {

	var result bool

	if antique.Architecture == value {
		result = true
	} else if antique.Architecture == "any" {
		result = true
	}

	return result

}

func (antique *Antique) MatchesManager(value string) bool {

	var result bool

	if antique.Manager == value {
		result = true
	} else if antique.Manager == "any" {
		result = true
	}

	return result

}

func (antique *Antique) MatchesName(value string) bool {

	var result bool

	if antique.Name == value {
		result = true
	} else if antique.Name == "any" {
		result = true
	}

	return result

}

func (antique *Antique) MatchesService(value string) bool {

	var result bool

	if antique.Service == value {
		result = true
	} else if antique.Service == "any" {
		result = true
	}

	return result

}

func (antique *Antique) MatchesVendor(value string) bool {

	var result bool

	if antique.Vendor == value {
		result = true
	} else if antique.Vendor == "any" {
		result = true
	}

	return result

}

func (antique *Antique) MatchesVersion(value string) bool {

	var result bool

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(value, " ") {
		value = strings.TrimSpace(value[strings.Index(value, " ")+1:])
	}

	if antique.Version == "any" {

		result = true

	} else if strings.HasPrefix(antique.Version, "<= ") {

		antique_version := types.ToVersion(antique.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.IsSame(antique_version) {
			result = true
		} else if other_version.IsBefore(antique_version) {
			result = true
		}

	} else if strings.HasPrefix(antique.Version, "< ") {

		antique_version := types.ToVersion(antique.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsBefore(antique_version) {
			result = true
		}

	} else if strings.HasPrefix(antique.Version, ">= ") {

		antique_version := types.ToVersion(antique.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.IsSame(antique_version) {
			result = true
		} else if other_version.IsAfter(antique_version) {
			result = true
		}

	} else if strings.HasPrefix(antique.Version, "> ") {

		antique_version := types.ToVersion(antique.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsAfter(antique_version) {
			result = true
		}

	} else if strings.HasPrefix(antique.Version, "= ") {

		antique_version := types.ToVersion(antique.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsSame(antique_version) {
			result = true
		}

	} else {

		antique_version := types.ToVersion(antique.Version)
		other_version := types.ToVersion(value)

		if other_version.IsSame(antique_version) {
			result = true
		}

	}

	return result

}

func (antique *Antique) Parse(value string) {

	name, version, architecture := parseVersionCondition(value)

	antique.Name = name
	antique.Version = version

	if architecture != "" {
		antique.Architecture = architecture
	}

}

func (antique *Antique) SetArchitecture(value string) {

	architecture := types.ParseArchitecture(value)

	if architecture != nil {
		antique.Architecture = architecture.String()
	}

}

func (antique *Antique) SetManager(value string) {

	manager := types.ParseManager(value)

	if manager != nil {
		antique.Manager = manager.String()
	}

}

func (antique *Antique) SetName(value string) {
	antique.Name = strings.TrimSpace(value)
}

func (antique *Antique) SetService(value string) {

	if value == "all" || value == "any" || value == "*" {
		antique.Service = "any"
	} else if value != "" {
		antique.Service = value
	}

}

func (antique *Antique) SetVersion(value string) {

	if value == "all" || value == "any" || value == "*" {
		antique.Version = "any"
	} else if value != "" {
		antique.Version = value
	}

}

func (antique *Antique) Hash() string {

	var hash string

	if antique.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			antique.Name,
			antique.Version,
			antique.Architecture,
			antique.Manager,
			antique.Service,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
