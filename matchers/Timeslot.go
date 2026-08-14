package matchers

import "github.com/cookiengineer/systemintegrity/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Timeslot struct {
	Offset   string `json:"offset"`
	Country  string `json:"country"`
	Timezone string `json:"timezone"`
}

func NewTimeslot() Timeslot {

	var timeslot Timeslot

	timeslot.Offset = "any"
	timeslot.Country = "any"
	timeslot.Timezone = "any"

	return timeslot

}

func ToTimeslot(value string) Timeslot {

	var timeslot Timeslot

	timeslot.Offset = "any"
	timeslot.Country = "any"
	timeslot.Timezone = "any"

	timeslot.Parse(value)

	return timeslot

}

func (timeslot *Timeslot) IsIdentical(value Timeslot) bool {

	var result bool

	if timeslot.Offset == value.Offset &&
		timeslot.Country == value.Country &&
		timeslot.Timezone == value.Timezone {
		result = true
	}

	return result

}

func (timeslot *Timeslot) IsValid() bool {

	var result bool

	if timeslot.Offset != "any" || timeslot.Country != "any" || timeslot.Timezone != "any" {
		result = true
	}

	return result

}

func (timeslot *Timeslot) Matches(offset string, country string, timezone string) bool {
	return timeslot.MatchesOffset(offset) && timeslot.MatchesCountry(country) && timeslot.MatchesTimezone(timezone)
}

func (timeslot *Timeslot) MatchesCountry(value string) bool {

	var result bool

	if timeslot.Country == value {
		result = true
	} else if timeslot.Country == "any" {
		result = true
	}

	return result

}

func (timeslot *Timeslot) MatchesOffset(value string) bool {

	var result bool

	// Compatibility with "<operator> <offset>" syntax
	if strings.Contains(value, " ") {
		value = strings.TrimSpace(value[strings.Index(value, " ")+1:])
	}

	if timeslot.Offset == "any" {

		result = true

	} else if strings.HasPrefix(timeslot.Offset, "<= ") {

		timeslot_offset := types.NewTime()
		timeslot_offset.Offset(timeslot.Offset[3:])

		other_offset := types.NewTime()
		other_offset.Offset(value)

		if other_offset.IsSame(timeslot_offset) {
			result = true
		} else if other_offset.IsBefore(timeslot_offset) {
			result = true
		}

	} else if strings.HasPrefix(timeslot.Offset, "< ") {

		timeslot_offset := types.NewTime()
		timeslot_offset.Offset(timeslot.Offset[2:])

		other_offset := types.NewTime()
		other_offset.Offset(value)

		if other_offset.IsBefore(timeslot_offset) {
			result = true
		}

	} else if strings.HasPrefix(timeslot.Offset, ">= ") {

		timeslot_offset := types.NewTime()
		timeslot_offset.Offset(timeslot.Offset[3:])

		other_offset := types.NewTime()
		other_offset.Offset(value)

		if other_offset.IsSame(timeslot_offset) {
			result = true
		} else if other_offset.IsAfter(timeslot_offset) {
			result = true
		}

	} else if strings.HasPrefix(timeslot.Offset, "> ") {

		timeslot_offset := types.NewTime()
		timeslot_offset.Offset(timeslot.Offset[2:])

		other_offset := types.NewTime()
		other_offset.Offset(value)

		if other_offset.IsAfter(timeslot_offset) {
			result = true
		}

	} else if strings.HasPrefix(timeslot.Offset, "= ") {

		timeslot_offset := types.NewTime()
		timeslot_offset.Offset(timeslot.Offset[2:])

		other_offset := types.NewTime()
		other_offset.Offset(value)

		if other_offset.IsSame(timeslot_offset) {
			result = true
		}

	} else {

		timeslot_offset := types.NewTime()
		timeslot_offset.Offset(timeslot.Offset)

		other_offset := types.NewTime()
		other_offset.Offset(value)

		if other_offset.IsSame(timeslot_offset) {
			result = true
		}

	}

	return result

}

func (timeslot *Timeslot) MatchesTimezone(value string) bool {

	var result bool

	if timeslot.Timezone == value {
		result = true
	} else if timeslot.Timezone == "any" {
		result = true
	} else if strings.HasPrefix(timeslot.Timezone, "*/") {

		if strings.HasSuffix(value, timeslot.Timezone[strings.Index(timeslot.Timezone, "*/")+1:]) {
			result = true
		}

	} else if strings.HasSuffix(timeslot.Timezone, "/*") {

		if strings.HasPrefix(value, timeslot.Timezone[0:strings.Index(timeslot.Timezone, "/*")+1]) {
			result = true
		}

	}

	return result

}

func (timeslot *Timeslot) Parse(value string) {

	timezone, offset := parseOffsetCondition(value)

	timeslot.Offset = offset
	timeslot.Timezone = timezone

}

func (timeslot *Timeslot) SetCountry(value string) {

	if value == "all" || value == "any" || value == "*" {
		timeslot.Country = "any"
	} else if value != "" {
		timeslot.Country = value
	}

}

func (timeslot *Timeslot) SetOffset(value string) {

	if value == "all" || value == "any" || value == "*" {
		timeslot.Offset = "any"
	} else if strings.HasPrefix(value, "+") || strings.HasPrefix(value, "-") {

		if len(value) == 6 && string(value[3]) == ":" {
			timeslot.Offset = value
		}

	}

}

func (timeslot *Timeslot) SetTimezone(value string) {

	if value == "all" || value == "any" || value == "*" {
		timeslot.Timezone = "any"
	} else if value != "" {
		timeslot.Timezone = strings.TrimSpace(value)
	}

}

func (timeslot *Timeslot) Hash() string {

	var hash string

	if timeslot.Offset != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			timeslot.Offset,
			timeslot.Country,
			timeslot.Timezone,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
