package matchers

import "github.com/cookiengineer/systemintegrity/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Package struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Manager      string `json:"manager"`
	Vendor       string `json:"vendor"`
}

func NewPackage() Package {

	var pkg Package

	pkg.Name = "any"
	pkg.Version = "any"
	pkg.Architecture = "any"
	pkg.Manager = "any"
	pkg.Vendor = "any"

	return pkg

}

func ToPackage(value string) Package {

	var pkg Package

	pkg.Name = "any"
	pkg.Version = "any"
	pkg.Architecture = "any"
	pkg.Manager = "any"
	pkg.Vendor = "any"

	pkg.Parse(value)

	return pkg

}

func (pkg *Package) IsIdentical(value Package) bool {

	var result bool

	if pkg.Name == value.Name &&
		pkg.Version == value.Version &&
		pkg.Architecture == value.Architecture &&
		pkg.Manager == value.Manager &&
		pkg.Vendor == value.Vendor {
		result = true
	}

	return result

}

func (pkg *Package) IsValid() bool {

	var result bool

	if pkg.Name != "any" || pkg.Version != "any" || pkg.Architecture != "any" || pkg.Manager != "any" || pkg.Vendor != "any" {
		result = true
	}

	return result

}

func (pkg *Package) Matches(name string, version string, architecture string, manager string, vendor string) bool {
	return pkg.MatchesName(name) && pkg.MatchesVersion(version) && pkg.MatchesArchitecture(architecture) && pkg.MatchesManager(manager) && pkg.MatchesVendor(vendor)
}

func (pkg *Package) MatchesArchitecture(value string) bool {

	var result bool

	if pkg.Architecture == value {
		result = true
	} else if pkg.Architecture == "any" {
		result = true
	}

	return result

}

func (pkg *Package) MatchesManager(value string) bool {

	var result bool

	if pkg.Manager == value {
		result = true
	} else if pkg.Manager == "any" {
		result = true
	}

	return result

}

func (pkg *Package) MatchesName(value string) bool {

	var result bool

	if pkg.Name == value {
		result = true
	} else if pkg.Name == "any" {
		result = true
	}

	return result

}

func (pkg *Package) MatchesVendor(value string) bool {

	var result bool

	if pkg.Vendor == value {
		result = true
	} else if pkg.Vendor == "any" {
		result = true
	}

	return result

}

func (pkg *Package) MatchesVersion(value string) bool {

	var result bool

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(value, " ") {
		value = strings.TrimSpace(value[strings.Index(value, " ")+1:])
	}

	if pkg.Version == "any" {

		result = true

	} else if strings.HasPrefix(pkg.Version, "<= ") {

		pkg_version := types.ToVersion(pkg.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.IsSame(pkg_version) {
			result = true
		} else if other_version.IsBefore(pkg_version) {
			result = true
		}

	} else if strings.HasPrefix(pkg.Version, "< ") {

		pkg_version := types.ToVersion(pkg.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsBefore(pkg_version) {
			result = true
		}

	} else if strings.HasPrefix(pkg.Version, ">= ") {

		pkg_version := types.ToVersion(pkg.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.IsSame(pkg_version) {
			result = true
		} else if other_version.IsAfter(pkg_version) {
			result = true
		}

	} else if strings.HasPrefix(pkg.Version, "> ") {

		pkg_version := types.ToVersion(pkg.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsAfter(pkg_version) {
			result = true
		}

	} else if strings.HasPrefix(pkg.Version, "= ") {

		pkg_version := types.ToVersion(pkg.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsSame(pkg_version) {
			result = true
		}

	} else {

		pkg_version := types.ToVersion(pkg.Version)
		other_version := types.ToVersion(value)

		if other_version.IsSame(pkg_version) {
			result = true
		}

	}

	return result

}

func (pkg *Package) Parse(value string) {

	name, version, architecture := parseVersionCondition(value)

	pkg.Name = name
	pkg.Version = version

	if architecture != "" {
		pkg.Architecture = architecture
	}

}

func (pkg *Package) SetArchitecture(value string) {

	architecture := types.ParseArchitecture(value)

	if architecture != nil {
		pkg.Architecture = architecture.String()
	}

}

func (pkg *Package) SetManager(value string) {

	manager := types.ParseManager(value)

	if manager != nil {
		pkg.Manager = manager.String()
	}

}

func (pkg *Package) SetName(value string) {
	pkg.Name = strings.TrimSpace(value)
}

func (pkg *Package) SetVendor(value string) {

	if value == "all" || value == "any" || value == "*" {
		pkg.Vendor = "any"
	} else if value != "" {
		pkg.Vendor = value
	}

}

func (pkg *Package) SetVersion(value string) {

	if value == "all" || value == "any" || value == "*" {
		pkg.Version = "any"
	} else if value != "" {
		pkg.Version = value
	}

}

func (pkg *Package) Hash() string {

	var hash string

	if pkg.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			pkg.Name,
			pkg.Version,
			pkg.Architecture,
			pkg.Manager,
			pkg.Vendor,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
