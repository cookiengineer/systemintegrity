package matchers

import "github.com/cookiengineer/systemintegrity/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Update struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Manager      string `json:"manager"`
	Vendor       string `json:"vendor"`
}

func NewUpdate() Update {

	var update Update

	update.Name = "any"
	update.Version = "any"
	update.Architecture = "any"
	update.Manager = "any"
	update.Vendor = "any"

	return update

}

func ToUpdate(value string) Update {

	var update Update

	update.Name = "any"
	update.Version = "any"
	update.Architecture = "any"
	update.Manager = "any"
	update.Vendor = "any"

	update.Parse(value)

	return update

}

func (update *Update) IsIdentical(value Update) bool {

	var result bool

	if update.Name == value.Name &&
		update.Version == value.Version &&
		update.Architecture == value.Architecture &&
		update.Manager == value.Manager &&
		update.Vendor == value.Vendor {
		result = true
	}

	return result

}

func (update *Update) IsValid() bool {

	var result bool

	if update.Name != "any" || update.Version != "any" || update.Architecture != "any" || update.Manager != "any" || update.Vendor != "any" {
		result = true
	}

	return result

}

func (update *Update) Matches(name string, version string, architecture string, manager string, vendor string) bool {
	return update.MatchesName(name) && update.MatchesVersion(version) && update.MatchesArchitecture(architecture) && update.MatchesManager(manager) && update.MatchesVendor(vendor)
}

func (update *Update) MatchesArchitecture(value string) bool {

	var result bool

	if update.Architecture == value {
		result = true
	} else if update.Architecture == "any" {
		result = true
	}

	return result

}

func (update *Update) MatchesManager(value string) bool {

	var result bool

	if update.Manager == value {
		result = true
	} else if update.Manager == "any" {
		result = true
	}

	return result

}

func (update *Update) MatchesName(value string) bool {

	var result bool

	if update.Name == value {
		result = true
	} else if update.Name == "any" {
		result = true
	}

	return result

}

func (update *Update) MatchesVendor(value string) bool {

	var result bool

	if update.Vendor == value {
		result = true
	} else if update.Vendor == "any" {
		result = true
	}

	return result

}

func (update *Update) MatchesVersion(value string) bool {

	var result bool

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(value, " ") {
		value = strings.TrimSpace(value[strings.Index(value, " ")+1:])
	}

	if update.Version == "any" {

		result = true

	} else if strings.HasPrefix(update.Version, "<= ") {

		update_version := types.ToVersion(update.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.IsSame(update_version) {
			result = true
		} else if other_version.IsBefore(update_version) {
			result = true
		}

	} else if strings.HasPrefix(update.Version, "< ") {

		update_version := types.ToVersion(update.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsBefore(update_version) {
			result = true
		}

	} else if strings.HasPrefix(update.Version, ">= ") {

		update_version := types.ToVersion(update.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.IsSame(update_version) {
			result = true
		} else if other_version.IsAfter(update_version) {
			result = true
		}

	} else if strings.HasPrefix(update.Version, "> ") {

		update_version := types.ToVersion(update.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsAfter(update_version) {
			result = true
		}

	} else if strings.HasPrefix(update.Version, "= ") {

		update_version := types.ToVersion(update.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsSame(update_version) {
			result = true
		}

	} else {

		update_version := types.ToVersion(update.Version)
		other_version := types.ToVersion(value)

		if other_version.IsSame(update_version) {
			result = true
		}

	}

	return result

}

func (update *Update) Parse(value string) {

	name, version, architecture := parseVersionCondition(value)

	update.Name = name
	update.Version = version

	if architecture != "" {
		update.Architecture = architecture
	}

}

func (update *Update) SetArchitecture(value string) {

	architecture := types.ParseArchitecture(value)

	if architecture != nil {
		update.Architecture = architecture.String()
	}

}

func (update *Update) SetManager(value string) {

	manager := types.ParseManager(value)

	if manager != nil {
		update.Manager = manager.String()
	}

}

func (update *Update) SetName(value string) {

	if value == "all" || value == "any" || value == "*" {
		update.Name = "any"
	} else if value != "" {
		update.Name = strings.TrimSpace(value)
	}

}

func (update *Update) SetVersion(value string) {

	if value == "all" || value == "any" || value == "*" {
		update.Version = "any"
	} else if value != "" {
		update.Version = value
	}

}

func (update *Update) Hash() string {

	var hash string

	if update.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			update.Name,
			update.Version,
			update.Architecture,
			update.Manager,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
