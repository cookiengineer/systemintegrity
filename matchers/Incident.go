package matchers

import "github.com/cookiengineer/systemintegrity/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Incident struct {
	System   string `json:"system"`
	Type     string `json:"type"`
	Datetime string `json:"datetime"`
}

func NewIncident() Incident {

	var incident Incident

	incident.System = "any"
	incident.Type = "any"
	incident.Datetime = "any"

	return incident

}

func ToIncident(value string) Incident {

	var incident Incident

	incident.System = "any"
	incident.Type = "any"
	incident.Datetime = "any"

	if strings.Contains(value, ":") {

		tmp := strings.Split(value, ":")

		if len(tmp) == 2 {

			incident.SetSystem(tmp[0])
			incident.SetType(tmp[1])

		}

	} else {
		incident.SetSystem(value)
	}

	return incident

}

func (incident *Incident) IsIdentical(value Incident) bool {

	var result bool

	if incident.System == value.System &&
		incident.Type == value.Type &&
		incident.Datetime == value.Datetime {
		result = true
	}

	return result

}

func (incident *Incident) IsValid() bool {

	var result bool

	if incident.System != "any" || incident.Type != "any" || incident.Datetime != "any" {
		result = true
	}

	return result

}

func (incident *Incident) Matches(system string, typ string, datetime string) bool {
	return incident.MatchesSystem(system) && incident.MatchesType(typ) && incident.MatchesDatetime(datetime)
}

func (incident *Incident) MatchesDatetime(value string) bool {

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

	if incident.Datetime == "any" {

		matches_from = true
		matches_until = true

	} else if strings.Contains(incident.Datetime, " - ") {

		tmp1 := strings.TrimSpace(incident.Datetime[0:strings.Index(incident.Datetime, " - ")])
		tmp2 := strings.TrimSpace(incident.Datetime[strings.Index(incident.Datetime, " - ")+3:])

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

	} else if strings.HasPrefix(incident.Datetime, "<= ") {

		current := types.ToDatetime(value)
		until := types.ToDatetime(strings.TrimSpace(incident.Datetime[3:]))
		matches_from = true

		if until.IsValid() {

			if until.IsSame(current) {
				matches_until = true
			} else if until.IsAfter(current) {
				matches_until = true
			}

		}

	} else if strings.HasPrefix(incident.Datetime, "< ") {

		current := types.ToDatetime(value)
		until := types.ToDatetime(strings.TrimSpace(incident.Datetime[2:]))
		matches_from = true

		if current.IsValid() && until.IsValid() {

			if until.IsAfter(current) {
				matches_until = true
			}

		}

	} else if strings.HasPrefix(incident.Datetime, ">= ") {

		current := types.ToDatetime(value)
		from := types.ToDatetime(strings.TrimSpace(incident.Datetime[3:]))
		matches_until = true

		if current.IsValid() && from.IsValid() {

			if from.IsSame(current) {
				matches_from = true
			} else if from.IsBefore(current) {
				matches_from = true
			}

		}

	} else if strings.HasPrefix(incident.Datetime, "> ") {

		current := types.ToDatetime(value)
		from := types.ToDatetime(strings.TrimSpace(incident.Datetime[2:]))
		matches_until = true

		if current.IsValid() && from.IsValid() {

			if from.IsBefore(current) {
				matches_from = true
			}

		}

	} else if strings.HasPrefix(incident.Datetime, "= ") {

		current := types.ToDatetime(value)
		datetime := types.ToDatetime(strings.TrimSpace(incident.Datetime[2:]))

		if current.IsValid() && datetime.IsValid() {

			if datetime.IsSame(current) {
				matches_from = true
				matches_until = true
			}

		}

	}

	return matches_from && matches_until

}

func (incident *Incident) MatchesSystem(value string) bool {

	var result bool

	if incident.System == value {
		result = true
	} else if incident.System == "any" {
		result = true
	}

	return result

}

func (incident *Incident) MatchesType(value string) bool {

	var result bool

	if incident.Type == value {
		result = true
	} else if incident.Type == "any" {
		result = true
	}

	return result

}

func (incident *Incident) SetDatetime(value string) {

	if value == "all" || value == "any" || value == "*" {
		incident.Datetime = "any"
	} else if strings.Contains(value, " - ") {

		tmp1 := strings.TrimSpace(value[0:strings.Index(value, " - ")])
		tmp2 := strings.TrimSpace(value[strings.Index(value, " - ")+3:])

		from := types.ToDatetime(tmp1)
		until := types.ToDatetime(tmp2)

		if from.IsBefore(until) {
			incident.Datetime = from.String() + " - " + until.String()
		}

	} else if strings.HasPrefix(value, "<= ") {

		until := types.ToDatetime(strings.TrimSpace(value[3:]))

		if until.IsValid() {
			incident.Datetime = "<= " + until.String()
		}

	} else if strings.HasPrefix(value, "< ") {

		until := types.ToDatetime(strings.TrimSpace(value[2:]))

		if until.IsValid() {
			incident.Datetime = "< " + until.String()
		}

	} else if strings.HasPrefix(value, ">= ") {

		from := types.ToDatetime(strings.TrimSpace(value[3:]))

		if from.IsValid() {
			incident.Datetime = ">= " + from.String()
		}

	} else if strings.HasPrefix(value, "> ") {

		from := types.ToDatetime(strings.TrimSpace(value[2:]))

		if from.IsValid() {
			incident.Datetime = "> " + from.String()
		}

	} else if strings.HasPrefix(value, "= ") {

		datetime := types.ToDatetime(strings.TrimSpace(value[2:]))

		if datetime.IsValid() {
			incident.Datetime = "= " + datetime.String()
		}

	} else {

		datetime := types.ToDatetime(strings.TrimSpace(value))

		if datetime.IsValid() {
			incident.Datetime = "= " + datetime.String()
		}

	}

}

func (incident *Incident) SetSystem(value string) {
	incident.System = strings.TrimSpace(value)
}

func (incident *Incident) SetType(value string) {

	if value == "all" || value == "any" || value == "*" {
		incident.Type = "any"
	} else if value != "" {
		incident.Type = strings.TrimSpace(value)
	}

}

func (incident *Incident) Hash() string {

	var hash string

	if incident.Type != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			incident.System,
			incident.Type,
			incident.Datetime,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
