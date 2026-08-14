package matchers

import "github.com/cookiengineer/systemintegrity/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strconv"
import "strings"

type Connection struct {
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	Protocol string `json:"protocol"`
	Type     string `json:"type"`
}

func NewConnection() Connection {

	var connection Connection

	connection.Host = "any"
	connection.Port = 0
	connection.Protocol = "any"
	connection.Type = "peer"

	return connection

}

func ToConnection(value string) Connection {

	var connection Connection

	connection.Host = "any"
	connection.Port = 0
	connection.Protocol = "any"
	connection.Type = "peer"

	connection.Parse(value)

	return connection

}

func (connection *Connection) IsIdentical(value Connection) bool {

	var result bool

	if connection.Host == value.Host &&
		connection.Port == value.Port &&
		connection.Protocol == value.Protocol &&
		connection.Type == value.Type {
		result = true
	}

	return result

}

func (connection *Connection) IsValid() bool {

	var result bool

	if connection.Host != "any" || (uint(connection.Port) > 0 && uint(connection.Port) < 65535) || connection.Protocol != "any" || connection.Type != "any" {
		result = true
	}

	return result

}

func (connection *Connection) Matches(host string, port uint16, protocol string, typ string) bool {
	return connection.MatchesHost(host) && connection.MatchesPort(port) && connection.MatchesProtocol(protocol) && connection.MatchesType(typ)
}

func (connection *Connection) MatchesHost(value string) bool {

	var result bool

	if connection.Host == value {
		result = true
	} else if connection.Host != "any" && value != "any" {
		result = containsSubnet(value, connection.Host)
	} else if connection.Host == "any" {
		result = true
	}

	return result

}

func (connection *Connection) MatchesPort(value uint16) bool {

	var result bool

	if connection.Port == value {
		result = true
	} else if connection.Port == 0 {
		result = true
	}

	return result

}

func (connection *Connection) MatchesProtocol(value string) bool {

	var result bool

	if connection.Protocol == value {
		result = true
	} else if connection.Protocol == "any" {
		result = true
	}

	return result

}

func (connection *Connection) MatchesType(value string) bool {

	var result bool

	if connection.Type == value {
		result = true
	} else if connection.Type == "any" {
		result = true
	}

	return result

}

func (connection *Connection) Parse(value string) {

	if strings.Contains(value, ":") {

		host := strings.TrimSpace(value[0:strings.LastIndex(value, ":")])
		port := strings.TrimSpace(value[strings.LastIndex(value, ":")+1:])

		if strings.HasSuffix(port, "TC") {

			connection.SetProtocol("tcp")
			connection.SetType("client")
			port = port[0 : len(port)-2]

		} else if strings.HasSuffix(port, "TP") {

			connection.SetProtocol("tcp")
			connection.SetType("peer")
			port = port[0 : len(port)-2]

		} else if strings.HasSuffix(port, "TS") {

			connection.SetProtocol("tcp")
			connection.SetType("server")
			port = port[0 : len(port)-2]

		} else if strings.HasSuffix(port, "UC") {

			connection.SetProtocol("udp")
			connection.SetType("client")
			port = port[0 : len(port)-2]

		} else if strings.HasSuffix(port, "UP") {

			connection.SetProtocol("udp")
			connection.SetType("peer")
			port = port[0 : len(port)-2]

		} else if strings.HasSuffix(port, "US") {

			connection.SetProtocol("udp")
			connection.SetType("server")
			port = port[0 : len(port)-2]

		} else if strings.HasSuffix(port, "AC") {

			connection.SetProtocol("any")
			connection.SetType("client")
			port = port[0 : len(port)-2]

		} else if strings.HasSuffix(port, "AP") {

			connection.SetProtocol("any")
			connection.SetType("peer")
			port = port[0 : len(port)-2]

		} else if strings.HasSuffix(port, "AS") {

			connection.SetProtocol("any")
			connection.SetType("server")
			port = port[0 : len(port)-2]

		}

		num, err := strconv.ParseUint(port, 10, 16)

		if err == nil {

			connection.SetHost(host)
			connection.SetPort(uint16(num))

		}

	} else {

		connection.SetHost(value)

	}

}

func (connection *Connection) SetHost(value string) {

	if value == "all" || value == "any" || value == "*" {

		connection.Host = "any"

	} else if types.IsIPv6(value) {

		ipv6 := types.ParseIPv6(value)

		if ipv6 != nil {
			connection.Host = ipv6.String()
		}

	} else if types.IsIPv4(value) {

		ipv4 := types.ParseIPv4(value)

		if ipv4 != nil {
			connection.Host = ipv4.String()
		}

	} else if types.IsDomain(value) {

		domain := types.ParseDomain(value)

		if domain != nil {
			connection.Host = domain.String()
		}

	} else if strings.Contains(value, "/") {

		address, prefix := toSubnet(value)

		if address != "" && prefix != 0 {
			connection.Host = value
		}

	}

}

func (connection *Connection) SetPort(value uint16) {

	if value > 0 && value < 65535 {
		connection.Port = value
	}

}

func (connection *Connection) SetProtocol(value string) {

	if value == "tcp" {
		connection.Protocol = "tcp"
	} else if value == "udp" {
		connection.Protocol = "udp"
	} else if value == "any" {
		connection.Protocol = "any"
	}

}

func (connection *Connection) SetType(value string) {

	if value == "client" {
		connection.Type = "client"
	} else if value == "peer" {
		connection.Type = "peer"
	} else if value == "server" {
		connection.Type = "server"
	} else if value == "any" {
		connection.Type = "peer"
	}

}

func (connection *Connection) Hash() string {

	var hash string

	if connection.Host != "" && connection.Port != 0 {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			connection.Host,
			strconv.FormatUint(uint64(connection.Port), 10),
			connection.Protocol,
			connection.Type,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)
	}

	return hash

}
