package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Drive struct {
	Name       string `json:"name"`
	Mountpoint string `json:"mountpoint"`
	Type       string `json:"type"`
}

func NewDrive() Drive {

	var drive Drive

	drive.Name = "any"
	drive.Mountpoint = "any"
	drive.Type = "any"

	return drive

}

func ToDrive(value string) Drive {

	var drive Drive

	drive.Type = "any"

	if strings.HasPrefix(value, "/") {
		drive.Name = "any"
		drive.SetMountpoint(value)
	} else if value != "" {
		drive.SetName(value)
		drive.Mountpoint = "any"
	}

	return drive

}

func (drive *Drive) IsIdentical(value Drive) bool {

	var result bool

	if drive.Name == value.Name &&
		drive.Mountpoint == value.Mountpoint &&
		drive.Type == value.Type {
		result = true
	}

	return result

}

func (drive *Drive) IsValid() bool {

	var result bool

	if drive.Name != "any" || drive.Mountpoint != "any" || drive.Type != "any" {
		result = true
	}

	return result

}

func (drive *Drive) Matches(name string, mountpoint string, typ string) bool {
	return drive.MatchesName(name) && drive.MatchesMountpoint(mountpoint) && drive.MatchesType(typ)
}

func (drive *Drive) MatchesMountpoint(value string) bool {

	var result bool

	if drive.Mountpoint == value {
		result = true
	} else if drive.Mountpoint == "any" {
		result = true
	}

	return result

}

func (drive *Drive) MatchesName(value string) bool {

	var result bool

	if drive.Name == value {
		result = true
	} else if drive.Name == "any" {
		result = true
	}

	return result

}

func (drive *Drive) MatchesType(value string) bool {

	var result bool

	if drive.Type == value {
		result = true
	} else if drive.Type == "any" {
		result = true
	}

	return result

}

func (drive *Drive) SetName(value string) {

	if value == "all" || value == "any" || value == "*" {
		drive.Name = "any"
	} else if value != "" {
		drive.Name = strings.TrimSpace(value)
	}

}

func (drive *Drive) SetMountpoint(value string) {

	if value == "all" || value == "any" || value == "*" {
		drive.Mountpoint = "any"
	} else if value != "" {
		drive.Mountpoint = strings.TrimSpace(value)
	}

}

func (drive *Drive) SetType(value string) {

	if value == "all" || value == "any" || value == "*" {
		drive.Type = "any"
	} else if value != "" {
		drive.Type = strings.TrimSpace(value)
	}

}

func (drive *Drive) Hash() string {

	var hash string

	if drive.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			drive.Name,
			drive.Mountpoint,
			drive.Type,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
