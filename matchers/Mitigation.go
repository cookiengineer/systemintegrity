package matchers

import "github.com/cookiengineer/systemintegrity/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Mitigation struct {
	System       string `json:"system"`
	IncidentType string `json:"incident_type"`
	ResponseType string `json:"response_type"`
	Datetime     string `json:"datetime"`
}

func NewMitigation() Mitigation {

	var mitigation Mitigation

	mitigation.System = "any"
	mitigation.IncidentType = "any"
	mitigation.ResponseType = "any"
	mitigation.Datetime = "any"

	return mitigation

}

func ToMitigation(value string) Mitigation {

	var mitigation Mitigation

	mitigation.System = "any"
	mitigation.IncidentType = "any"
	mitigation.ResponseType = "any"
	mitigation.Datetime = "any"

	mitigation.SetSystem(value)

	return mitigation

}

func (mitigation *Mitigation) IsIdentical(value Mitigation) bool {

	var result bool

	if mitigation.System == value.System &&
		mitigation.IncidentType == value.IncidentType &&
		mitigation.ResponseType == value.ResponseType &&
		mitigation.Datetime == value.Datetime {
		result = true
	}

	return result

}

func (mitigation *Mitigation) IsValid() bool {

	var result bool

	if mitigation.System != "any" || mitigation.IncidentType != "any" || mitigation.ResponseType != "any" || mitigation.Datetime != "any" {
		result = true
	}

	return result

}

func (mitigation *Mitigation) SetDatetime(value string) {

	if value == "all" || value == "any" || value == "*" {
		mitigation.Datetime = "any"
	} else if strings.Contains(value, " - ") {

		tmp1 := strings.TrimSpace(value[0:strings.Index(value, " - ")])
		tmp2 := strings.TrimSpace(value[strings.Index(value, " - ")+3:])

		from := types.ToDatetime(tmp1)
		until := types.ToDatetime(tmp2)

		if from.IsBefore(until) {
			mitigation.Datetime = from.String() + " - " + until.String()
		}

	} else if strings.HasPrefix(value, "<= ") {

		until := types.ToDatetime(strings.TrimSpace(value[3:]))

		if until.IsValid() {
			mitigation.Datetime = "<= " + until.String()
		}

	} else if strings.HasPrefix(value, "< ") {

		until := types.ToDatetime(strings.TrimSpace(value[2:]))

		if until.IsValid() {
			mitigation.Datetime = "< " + until.String()
		}

	} else if strings.HasPrefix(value, ">= ") {

		from := types.ToDatetime(strings.TrimSpace(value[3:]))

		if from.IsValid() {
			mitigation.Datetime = ">= " + from.String()
		}

	} else if strings.HasPrefix(value, "> ") {

		from := types.ToDatetime(strings.TrimSpace(value[2:]))

		if from.IsValid() {
			mitigation.Datetime = "> " + from.String()
		}

	} else if strings.HasPrefix(value, "= ") {

		datetime := types.ToDatetime(strings.TrimSpace(value[2:]))

		if datetime.IsValid() {
			mitigation.Datetime = "= " + datetime.String()
		}

	} else {

		datetime := types.ToDatetime(strings.TrimSpace(value))

		if datetime.IsValid() {
			mitigation.Datetime = "= " + datetime.String()
		}

	}

}

func (mitigation *Mitigation) SetSystem(value string) {
	mitigation.System = strings.TrimSpace(value)
}

func (mitigation *Mitigation) SetIncidentType(value string) {

	if value == "all" || value == "any" || value == "*" {
		mitigation.IncidentType = "any"
	} else if value != "" {
		mitigation.IncidentType = strings.TrimSpace(value)
	}

}

func (mitigation *Mitigation) SetResponseType(value string) {

	if value == "all" || value == "any" || value == "*" {
		mitigation.ResponseType = "any"
	} else if value != "" {
		mitigation.ResponseType = strings.TrimSpace(value)
	}

}

func (mitigation *Mitigation) Hash() string {

	var hash string

	if mitigation.IncidentType != "" && mitigation.ResponseType != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			mitigation.System,
			mitigation.IncidentType,
			mitigation.ResponseType,
			mitigation.Datetime,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
