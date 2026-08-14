package matchers

import "github.com/cookiengineer/systemintegrity/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type System struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Datetime string `json:"datetime"`
	Country  string `json:"country"`
	Timezone string `json:"timezone"`
}

func NewSystem() System {

	var system System

	system.Name = "any"
	system.Hostname = "any"
	system.Datetime = "any"
	system.Country = "any"
	system.Timezone = "any"

	return system

}

func ToSystem(value string) System {

	var system System

	system.Name = "any"
	system.Hostname = "any"
	system.Datetime = "any"
	system.Country = "any"
	system.Timezone = "any"

	system.SetName(value)

	return system

}

func (system *System) IsIdentical(value System) bool {

	var result bool

	if system.Name == value.Name &&
		system.Hostname == value.Hostname &&
		system.Datetime == value.Datetime &&
		system.Country == value.Country &&
		system.Timezone == value.Timezone {
		result = true
	}

	return result

}

func (system *System) IsValid() bool {

	var result bool

	if system.Name != "any" || system.Hostname != "any" || system.Datetime != "any" || system.Country != "any" || system.Timezone != "any" {
		result = true
	}

	return result

}

func (system *System) Matches(name string, hostname string, datetime string, country string, timezone string) bool {
	return system.MatchesName(name) && system.MatchesHostname(hostname) && system.MatchesDatetime(datetime) && system.MatchesCountry(country) && system.MatchesTimezone(timezone)
}

func (system *System) MatchesCountry(value string) bool {

	var result bool

	if system.Country == value {
		result = true
	} else if system.Country == "any" {
		result = true
	}

	return result

}


func (system *System) MatchesDatetime(value string) bool {

	var matches_from bool
	var matches_until bool

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(value, " ") {

		tmp := strings.Split(value, " ")

		// >= 2001-02-03 04:05:06
		if len(tmp) == 3 {
			value = strings.TrimSpace(tmp[1] + " " + tmp[2])
		}

	}

	if system.Datetime == "any" {

		matches_from = true
		matches_until = true

	} else if strings.Contains(system.Datetime, " - ") {

		tmp1 := strings.TrimSpace(system.Datetime[0:strings.Index(system.Datetime, " - ")])
		tmp2 := strings.TrimSpace(system.Datetime[strings.Index(system.Datetime, " - ")+3:])

		from := types.ToDatetime(tmp1)
		until := types.ToDatetime(tmp2)
		current := types.ToDatetime(value)

		if current.IsValid() && from.IsBefore(until) {

			if from.IsBefore(current) {
				matches_from = true
			}

			if until.IsAfter(current) {
				matches_until = true
			}

		}

	} else if strings.HasPrefix(system.Datetime, "<= ") {

		current := types.ToDatetime(value)
		until := types.ToDatetime(strings.TrimSpace(system.Datetime[3:]))
		matches_from = true

		if until.IsValid() {

			if until.IsSame(current) {
				matches_until = true
			} else if until.IsAfter(current) {
				matches_until = true
			}

		}

	} else if strings.HasPrefix(system.Datetime, "< ") {

		current := types.ToDatetime(value)
		until := types.ToDatetime(strings.TrimSpace(system.Datetime[2:]))
		matches_from = true

		if current.IsValid() && until.IsValid() {

			if until.IsAfter(current) {
				matches_until = true
			}

		}

	} else if strings.HasPrefix(system.Datetime, ">= ") {

		current := types.ToDatetime(value)
		from := types.ToDatetime(strings.TrimSpace(system.Datetime[3:]))
		matches_until = true

		if current.IsValid() && from.IsValid() {

			if from.IsSame(current) {
				matches_from = true
			} else if from.IsBefore(current) {
				matches_from = true
			}

		}

	} else if strings.HasPrefix(system.Datetime, "> ") {

		current := types.ToDatetime(value)
		from := types.ToDatetime(strings.TrimSpace(system.Datetime[2:]))
		matches_until = true

		if current.IsValid() && from.IsValid() {

			if from.IsBefore(current) {
				matches_from = true
			}

		}

	} else if strings.HasPrefix(system.Datetime, "= ") {

		current := types.ToDatetime(value)
		datetime := types.ToDatetime(strings.TrimSpace(system.Datetime[2:]))

		if current.IsValid() && datetime.IsValid() {

			if datetime.IsSame(current) {
				matches_from = true
				matches_until = true
			}

		}

	}

	return matches_from && matches_until

}

func (system *System) MatchesHostname(value string) bool {

	var result bool

	if system.Hostname == value {
		result = true
	} else if system.Hostname == "any" {
		result = true
	}

	return result

}

func (system *System) MatchesName(value string) bool {

	var result bool

	if system.Name == value {
		result = true
	} else if system.Name == "any" {
		result = true
	}

	return result

}

func (system *System) MatchesTimezone(value string) bool {

	var result bool

	if system.Timezone == value {
		result = true
	} else if system.Timezone == "any" {
		result = true
	} else if strings.HasPrefix(system.Timezone, "*/") {

		if strings.HasSuffix(value, system.Timezone[strings.Index(system.Timezone, "*/")+1:]) {
			result = true
		}

	} else if strings.HasSuffix(system.Timezone, "/*") {

		if strings.HasPrefix(value, system.Timezone[0:strings.Index(system.Timezone, "/*")+1]) {
			result = true
		}

	}

	return result

}

func (system *System) SetCountry(value string) {

	if value == "all" || value == "any" || value == "*" {
		system.Country = "any"
	} else if value != "" {
		system.Country = value
	}

}

func (system *System) SetHostname(value string) {

	if value == "all" || value == "any" || value == "*" {
		system.Hostname = "any"
	} else if value != "" {
		system.Hostname = strings.TrimSpace(value)
	}

}

func (system *System) SetName(value string) {
	system.Name = strings.TrimSpace(value)
}

func (system *System) SetDatetime(value string) {

	if value == "all" || value == "any" || value == "*" {
		system.Datetime = "any"
	} else if strings.Contains(value, " - ") {

		tmp1 := strings.TrimSpace(value[0:strings.Index(value, " - ")])
		tmp2 := strings.TrimSpace(value[strings.Index(value, " - ")+3:])

		from := types.ToDatetime(tmp1)
		until := types.ToDatetime(tmp2)

		if from.IsBefore(until) {
			system.Datetime = from.String() + " - " + until.String()
		}

	} else if strings.HasPrefix(value, "<= ") {

		until := types.ToDatetime(strings.TrimSpace(value[3:]))

		if until.IsValid() {
			system.Datetime = "<= " + until.String()
		}

	} else if strings.HasPrefix(value, "< ") {

		until := types.ToDatetime(strings.TrimSpace(value[2:]))

		if until.IsValid() {
			system.Datetime = "< " + until.String()
		}

	} else if strings.HasPrefix(value, ">= ") {

		from := types.ToDatetime(strings.TrimSpace(value[3:]))

		if from.IsValid() {
			system.Datetime = ">= " + from.String()
		}

	} else if strings.HasPrefix(value, "> ") {

		from := types.ToDatetime(strings.TrimSpace(value[2:]))

		if from.IsValid() {
			system.Datetime = "> " + from.String()
		}

	} else if strings.HasPrefix(value, "= ") {

		datetime := types.ToDatetime(strings.TrimSpace(value[2:]))

		if datetime.IsValid() {
			system.Datetime = "= " + datetime.String()
		}

	} else {

		datetime := types.ToDatetime(strings.TrimSpace(value))

		if datetime.IsValid() {
			system.Datetime = "= " + datetime.String()
		}

	}

}

func (system *System) SetTimezone(value string) {

	if value == "all" || value == "any" || value == "*" {
		system.Timezone = "any"
	} else if value != "" {
		system.Timezone = strings.TrimSpace(value)
	}

}

func (system *System) Hash() string {

	var hash string

	if system.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			system.Name,
			system.Hostname,
			system.Datetime,
			system.Country,
			system.Timezone,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
