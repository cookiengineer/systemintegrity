package types

import "strconv"
import "strings"

type Connection struct {
	Source   Socket   `json:"source"`
	Target   Socket   `json:"target"`
	Protocol Protocol `json:"protocol"`
	Type     string   `json:"type"`
}

func IsConnection(value string) bool {

	var result bool

	if strings.Contains(value, " ") {

		tmp := strings.Split(strings.TrimSpace(value), " ")

		if len(tmp) == 4 {

			// [protocol] local:port -> remote:port
			// [protocol] local:port <- remote:port
			// [protocol] local:port <-> remote:port

			valid_source := false
			valid_target := false
			valid_protocol := false
			valid_type := false

			if strings.HasPrefix(tmp[0], "[") && strings.HasSuffix(tmp[0], "]") && IsProtocol(tmp[0][1:len(tmp[0])-1]) {
				valid_protocol = true
			}

			if strings.HasPrefix(tmp[1], "*:") {

				num, err := strconv.ParseUint(tmp[1][2:], 10, 16)

				if err == nil && num > 0 && num < 65535 {
					valid_source = true
				}

			} else if strings.HasSuffix(tmp[1], ":*") {

				if IsIPv4(tmp[1][0:len(tmp[1])-2]) {
					valid_source = true
				} else if IsIPv6(tmp[1][0:len(tmp[1])-2]) {
					valid_source = true
				}

			} else if IsIPv4AndPort(tmp[1]) {

				ipv4, port := ParseIPv4AndPort(tmp[1])

				if ipv4 != nil && port > 0 && port < 65535 {
					valid_source = true
				}

			} else if IsIPv6AndPort(tmp[1]) {

				ipv6, port := ParseIPv6AndPort(tmp[1])

				if ipv6 != nil && port > 0 && port < 65535 {
					valid_source = true
				}

			} else if IsDomainAndPort(tmp[1]) {
				valid_source = false
			}

			if tmp[2] == "->" {
				valid_type = true
			} else if tmp[2] == "<-" {
				valid_type = true
			} else if tmp[2] == "<->" {
				valid_type = true
			}

			if strings.HasPrefix(tmp[3], "*:") {

				num, err := strconv.ParseUint(tmp[3][2:], 10, 16)

				if err == nil && num > 0 && num < 65535 {
					valid_target = true
				}

			} else if strings.HasSuffix(tmp[3], ":*") {

				if IsIPv4(tmp[3][0:len(tmp[3])-2]) {
					valid_target = true
				} else if IsIPv6(tmp[3][0:len(tmp[3])-2]) {
					valid_target = true
				}

			} else if IsIPv4AndPort(tmp[3]) {

				ipv4, port := ParseIPv4AndPort(tmp[3])

				if ipv4 != nil && port > 0 && port < 65535 {
					valid_target = true
				}

			} else if IsIPv6AndPort(tmp[3]) {

				ipv6, port := ParseIPv6AndPort(tmp[3])

				if ipv6 != nil && port > 0 && port < 65535 {
					valid_target = true
				}

			} else if IsDomainAndPort(tmp[3]) {

				domain, port := ParseDomainAndPort(tmp[3])

				if domain != nil && port > 0 && port < 65535 {
					valid_target = true
				}

			}

			if valid_protocol && valid_source && valid_type && valid_target {
				result = true
			}

		}

	}

	return result

}

func NewConnection() Connection {

	var connection Connection

	connection.Source = NewSocket("0.0.0.0", 0)
	connection.Target = NewSocket("0.0.0.0", 0)
	connection.Protocol = ProtocolANY
	connection.Type = ""

	return connection

}

func ParseConnection(value string) *Connection {

	var result *Connection = nil

	if strings.Contains(value, " ") {

		tmp := strings.Split(strings.TrimSpace(value), " ")

		if len(tmp) == 4 {

			// [protocol] local:port -> remote:port
			// [protocol] local:port <- remote:port
			// [protocol] local:port <-> remote:port

			protocol := ""
			source_host := "0.0.0.0"
			source_port := uint16(0)
			target_host := "0.0.0.0"
			target_port := uint16(0)
			typ := ""

			if strings.HasPrefix(tmp[0], "[") && strings.HasSuffix(tmp[0], "]") && IsProtocol(tmp[0][1:len(tmp[0])-1]) {
				protocol = strings.TrimSpace(tmp[0][1:len(tmp[0])-1])
			}

			if strings.HasPrefix(tmp[1], "*:") {

				num, err := strconv.ParseUint(tmp[1][2:], 10, 16)

				if err == nil && num > 0 && num < 65535 {
					source_port = uint16(num)
				}

			} else if strings.HasSuffix(tmp[1], ":*") {

				if IsIPv4(tmp[1][0:len(tmp[1])-2]) {

					ipv4 := ParseIPv4(tmp[1][0:len(tmp[1])-2])

					if ipv4 != nil {
						source_host = ipv4.String()
						source_port = uint16(0)
					}

				} else if IsIPv6(tmp[1][0:len(tmp[1])-2]) {

					ipv6 := ParseIPv6(tmp[1][0:len(tmp[1])-2])

					if ipv6 != nil {
						source_host = ipv6.String()
						source_port = uint16(0)
					}

				}

			} else if IsIPv4AndPort(tmp[1]) {

				ipv4, port := ParseIPv4AndPort(tmp[1])

				if ipv4 != nil && port > 0 && port < 65535 {
					source_host = ipv4.String()
					source_port = port
				}

			} else if IsIPv6AndPort(tmp[1]) {

				ipv6, port := ParseIPv6AndPort(tmp[1])

				if ipv6 != nil && port > 0 && port < 65535 {
					source_host = ipv6.String()
					source_port = port
				}

			} else if IsDomainAndPort(tmp[1]) {

				source_host = ""
				source_port = uint16(0)

			}

			if tmp[2] == "->" {
				typ = "client"
			} else if tmp[2] == "<-" {
				typ = "server"
			} else if tmp[2] == "<->" {
				typ = "peer"
			}

			if strings.HasPrefix(tmp[3], "*:") {

				num, err := strconv.ParseUint(tmp[3][2:], 10, 16)

				if err == nil {
					target_port = uint16(num)
				}

			} else if strings.HasSuffix(tmp[3], ":*") {

				if IsIPv4(tmp[3][0:len(tmp[3])-2]) {

					ipv4 := ParseIPv4(tmp[3][0:len(tmp[3])-2])

					if ipv4 != nil {
						target_host = ipv4.String()
						target_port = uint16(0)
					}

				} else if IsIPv6(tmp[3][0:len(tmp[3])-2]) {

					ipv6 := ParseIPv6(tmp[3][0:len(tmp[3])-2])

					if ipv6 != nil {
						target_host = ipv6.String()
						target_port = uint16(0)
					}

				}

			} else if IsIPv4AndPort(tmp[3]) {

				ipv4, port := ParseIPv4AndPort(tmp[3])

				if ipv4 != nil && port > 0 && port < 65535 {
					target_host = ipv4.String()
					target_port = port
				}

			} else if IsIPv6AndPort(tmp[3]) {

				ipv6, port := ParseIPv6AndPort(tmp[3])

				if ipv6 != nil && port > 0 && port < 65535 {
					target_host = ipv6.String()
					target_port = port
				}

			} else if IsDomainAndPort(tmp[3]) {

				domain, port := ParseDomainAndPort(tmp[3])

				if domain != nil && port > 0 && port < 65535 {
					target_host = domain.String()
					target_port = port
				}

			}

			if protocol != "" && typ != "" {

				if source_host != "" && source_port > 0 && source_port < 65535 {

					if target_host != "" && target_port > 0 && target_port < 65535 {

						result = &Connection{
							Source:   NewSocket(source_host, source_port),
							Target:   NewSocket(target_host, target_port),
							Protocol: Protocol(protocol),
							Type:     typ,
						}

					}

				}

			}

		}

	}

	return result

}

func ToConnection(value string) Connection {

	var connection Connection

	if value != "" {

		tmp := ParseConnection(value)

		if tmp != nil {
			connection = *tmp
		}

	}

	return connection

}

func (connection *Connection) IsIdentical(value Connection) bool {

	var result bool

	if connection.Source.Host == value.Source.Host &&
		connection.Source.Port == value.Source.Port &&
		connection.Target.Host == value.Target.Host &&
		connection.Target.Port == value.Target.Port &&
		connection.Protocol == value.Protocol &&
		connection.Type == value.Type {
		result = true
	}

	return result

}

func (connection *Connection) IsValid() bool {

	var result bool

	if connection.Type == "client" {

		if connection.Source.Type == "ipv4" && connection.Target.Type == "ipv4" {

			// ipv4:port -> host and port
			// ipv4:0    -> host
			// *:port    -> port

			if connection.Target.Host == "*" && connection.Target.Port != 0 {
				result = true
			} else if connection.Target.Host != "0.0.0.0" && connection.Target.Port != 0 {
				result = true
			} else if connection.Target.Host != "0.0.0.0" && connection.Target.Port == 0 {
				result = true
			}

		} else if connection.Source.Type == "ipv6" && connection.Target.Type == "ipv6" {

			// ipv6:port -> host and port
			// ipv6:0    -> host
			// *:port    -> port

			if connection.Target.Host == "*" && connection.Target.Port != 0 {
				result = true
			} else if connection.Target.Host != "[0000:0000:0000:0000:0000:0000:0000:0000]" && connection.Target.Port != 0 {
				result = true
			} else if connection.Target.Host != "[0000:0000:0000:0000:0000:0000:0000:0000]" && connection.Target.Port == 0 {
				result = true
			}

		} else if connection.Target.Type == "domain" {

			// domain:port -> domain and port
			// domain:0    -> domain

			if connection.Target.Host != "*" {
				result = true
			}

		}

	} else if connection.Type == "server" {

		if connection.Target.Type == "ipv4" {

			// servers can be bound to 0.0.0.0 host

			if connection.Source.Host == "*" && connection.Source.Port != 0 {
				result = true
			} else if connection.Source.Host != "0.0.0.0" && connection.Source.Port != 0 {
				result = true
			} else if connection.Source.Host != "0.0.0.0" && connection.Source.Port == 0 {
				result = true
			} else if connection.Target.Host != "0.0.0.0" && connection.Target.Port != 0 {
				result = true
			} else if connection.Target.Host == "0.0.0.0" && connection.Target.Port != 0 {
				result = true
			}

		} else if connection.Target.Type == "ipv6" {

			// servers can be bound to ::0 host

			if connection.Source.Host == "*" && connection.Source.Port != 0 {
				result = true
			} else if connection.Source.Host != "[0000:0000:0000:0000:0000:0000:0000:0000]" && connection.Source.Port != 0 {
				result = true
			} else if connection.Source.Host != "[0000:0000:0000:0000:0000:0000:0000:0000]" && connection.Source.Port == 0 {
				result = true
			}

		} else if connection.Target.Type == "domain" {

			// Reverse DNS Blocking (?) Not Supported

		}

	} else if connection.Type == "peer" {

		if connection.Source.Type == "ipv4" && connection.Target.Type == "ipv4" {

			if connection.Source.Host != "*" && connection.Target.Host != "*" {

				if connection.Source.Host != "0.0.0.0" && connection.Source.Port != 0 {

					if connection.Target.Host != "0.0.0.0" && connection.Target.Port != 0 {
						result = true
					}

				}

			}

		} else if connection.Source.Type == "ipv6" && connection.Target.Type == "ipv6" {

			if connection.Source.Host != "*" && connection.Target.Host != "*" {

				if connection.Source.Host != "[0000:0000:0000:0000:0000:0000:0000:0000]" && connection.Source.Port != 0 {

					if connection.Target.Host != "[0000:0000:0000:0000:0000:0000:0000:0000]" && connection.Target.Port != 0 {
						result = true
					}

				}

			}

		}

	}

	return result

}

func (connection *Connection) SetSource(value Socket) {

	if value.IsValid() && value.Port > 0 && value.Port < 65535 {

		if value.Host == "*" {

			connection.Source = Socket{
				Host:  "0.0.0.0",
				Port:  value.Port,
				Scope: "private",
				Type:  "ipv4",
			}

		} else {
			connection.Source = value
		}

	}

}

func (connection *Connection) SetTarget(value Socket) {

	if value.IsValid() && value.Port > 0 && value.Port < 65535 {

		if value.Host == "*" {

			connection.Target = Socket{
				Host:  "0.0.0.0",
				Port:  value.Port,
				Scope: "private",
				Type:  "ipv4",
			}

		} else {
			connection.Target = value
		}

	}

}

func (connection *Connection) SetProtocol(value Protocol) {

	if value == ProtocolTCP {
		connection.Protocol = value
	} else if value == ProtocolUDP {
		connection.Protocol = value
	} else if value == ProtocolANY {
		connection.Protocol = value
	}

}

func (connection *Connection) SetType(value string) {

	if value == "client" {

		// established, syn_sent, syn_recv, fin_wait1, fin_wait2, new_syn_recv

		connection.Type = "client"

	} else if value == "server" {

		// listen

		connection.Type = "server"

	} else if value == "peer" {

		// both directions

		connection.Type = "peer"

	}

}

func (connection *Connection) String() string {

	result := ""

	if connection.Type != "" {

		typ := "-"

		if connection.Type == "client" {
			typ = "->"
		} else if connection.Type == "server" {
			typ = "<-"
		} else if connection.Type == "peer" {
			typ = "<->"
		}

		result = strings.Join([]string{
			"[" + connection.Protocol.String() + "]",
			connection.Source.String(),
			typ,
			connection.Target.String(),
		}, " ")

	}

	return result

}

func (connection Connection) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(connection.String())), nil
}

func (connection *Connection) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))

	if err != nil {
		return err
	}

	tmp := ParseConnection(unquoted)

	if tmp != nil {
		*connection = *tmp
	}

	return nil

}
