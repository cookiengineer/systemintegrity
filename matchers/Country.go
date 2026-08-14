package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Country struct {
	Name       string `json:"name"`
	Continent  string `json:"continent"`
	Allegiance string `json:"allegiance"`
	Subnet     string `json:"subnet"`
	Timezone   string `json:"timezone"`
}

func NewCountry() Country {

	var country Country

	country.Name = "any"
	country.Continent = "any"
	country.Allegiance = "any"
	country.Subnet = "any"
	country.Timezone = "any"

	return country

}

func ToCountry(value string) Country {

	var country Country

	country.Name = "any"
	country.Continent = "any"
	country.Allegiance = "any"
	country.Subnet = "any"
	country.Timezone = "any"

	country.SetName(value)

	return country

}

func (country *Country) IsIdentical(value Country) bool {

	var result bool

	if country.Name == value.Name &&
		country.Continent == value.Continent &&
		country.Allegiance == value.Allegiance &&
		country.Subnet == value.Subnet &&
		country.Timezone == value.Timezone {
		result = true
	}

	return result

}

func (country *Country) IsValid() bool {

	var result bool

	if country.Name != "any" || country.Continent != "any" || country.Allegiance != "any" || country.Subnet != "any" || country.Timezone != "any" {
		result = true
	}

	return result

}

func (country *Country) Matches(name string, continent string, allegiance string, subnet string, timezone string) bool {
	return country.MatchesName(name) && country.MatchesContinent(continent) && country.MatchesAllegiance(allegiance) && country.MatchesSubnet(subnet) && country.MatchesTimezone(timezone)
}

func (country *Country) MatchesAllegiance(value string) bool {

	var result bool

	if country.Allegiance == value {
		result = true
	} else if country.Allegiance == "any" {
		result = true
	}

	return result

}

func (country *Country) MatchesContinent(value string) bool {

	var result bool

	if country.Continent == value {
		result = true
	} else if country.Continent == "any" {
		result = true
	}

	return result

}

func (country *Country) MatchesName(value string) bool {

	var result bool

	if country.Name == value {
		result = true
	} else if country.Name == "any" {
		result = true
	}

	return result

}

func (country *Country) MatchesSubnet(value string) bool {

	var result bool

	if country.Subnet != "any" && value != "any" {
		result = containsSubnet(value, country.Subnet)
	} else if country.Subnet == "any" {
		result = true
	}

	return result

}

func (country *Country) MatchesTimezone(value string) bool {

	var result bool

	if country.Timezone == value {
		result = true
	} else if country.Timezone == "any" {
		result = true
	}

	return result

}

func (country *Country) SetAllegiance(value string) {

	if value == "all" || value == "any" || value == "*" {
		country.Allegiance = "any"
	} else if value != "" {
		country.Allegiance = value
	}

}

func (country *Country) SetName(value string) {
	country.Name = strings.TrimSpace(value)
}

func (country *Country) SetContinent(value string) {

	if value == "all" || value == "any" || value == "*" {
		country.Continent = "any"
	} else if value != "" {
		country.Continent = value
	}

}

func (country *Country) SetSubnet(value string) {

	address, prefix := toSubnet(value)

	if value == "all" || value == "any" || value == "*" {
		country.Subnet = "any"
	} else if address != "" && prefix != 0 {
		country.Subnet = value
	}

}

func (country *Country) SetTimezone(value string) {

	if value == "all" || value == "any" || value == "*" {
		country.Timezone = "any"
	} else if value != "" {
		country.Timezone = value
	}

}

func (country *Country) Hash() string {

	var hash string

	if country.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			country.Name,
			country.Continent,
			country.Allegiance,
			country.Subnet,
			country.Timezone,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
