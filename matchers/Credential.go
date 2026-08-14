package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Credential struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

func NewCredential() Credential {

	var credential Credential

	credential.Name = "any"
	credential.Password = "any"
	credential.Type = "any"

	return credential

}

func ToCredential(value string) Credential {

	var credential Credential

	credential.Name = "any"
	credential.Password = "any"
	credential.Type = "any"

	credential.SetName(value)

	return credential

}

func (credential *Credential) IsIdentical(value Credential) bool {

	var result bool

	if credential.Name == value.Name &&
		credential.Password == value.Password &&
		credential.Type == value.Type {
		result = true
	}

	return result

}

func (credential *Credential) IsValid() bool {

	var result bool

	if credential.Name != "any" || credential.Password != "any" || credential.Type != "any" {
		result = true
	}

	return result

}

func (credential *Credential) Matches(name string, password string, typ string) bool {
	return credential.MatchesName(name) && credential.MatchesPassword(password) && credential.MatchesType(typ)
}

func (credential *Credential) MatchesName(value string) bool {

	var result bool

	if credential.Name == value {
		result = true
	} else if credential.Name == "any" {
		result = true
	}

	return result

}

func (credential *Credential) MatchesPassword(value string) bool {

	var result bool

	if credential.Password == value {
		result = true
	} else if credential.Password == "any" {
		result = true
	}

	return result

}

func (credential *Credential) MatchesType(value string) bool {

	var result bool

	if credential.Type == value {
		result = true
	} else if credential.Type == "any" {
		result = true
	}

	return result

}

func (credential *Credential) SetName(value string) {
	credential.Name = strings.TrimSpace(value)
}

func (credential *Credential) SetPassword(value string) {
	credential.Password = value
}

func (credential *Credential) SetType(value string) {

	if value == "all" || value == "any" || value == "*" {
		credential.Type = "any"
	} else if value != "" {
		credential.Type = value
	}

}

func (credential *Credential) Hash() string {

	var hash string

	if credential.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			credential.Name,
			credential.Password,
			credential.Type,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
