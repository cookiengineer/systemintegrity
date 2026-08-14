package types

import "strconv"
import "strings"

type Socket struct {
	Host  string `json:"host"`
	Port  uint16 `json:"port"`
	Scope string `json:"scope"`
	Type  string `json:"type"`
}

func IsSocket(value string) bool {

	var result bool

	if strings.HasPrefix(value, "[") && strings.Contains(value, "]:") {

		_, err := strconv.ParseInt(strings.Split(value, "]:")[1], 10, 16)

		if err == nil {
			result = true
		}

	} else if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {

		result = true

	} else if strings.Contains(value, ".") && strings.Contains(value, ":") {

		_, err := strconv.ParseInt(strings.Split(value, ":")[1], 10, 16)

		if err == nil {
			result = true
		}

	} else if strings.Contains(value, ":") {

		_, err := strconv.ParseInt(strings.Split(value, ":")[1], 10, 16)

		if err == nil {
			result = true
		}

	} else if strings.Contains(value, ".") {

		result = true

	}

	return result

}

func NewSocket(host string, port uint16) Socket {

	var socket Socket

	socket.Host = ""
	socket.Port = 0
	socket.Scope = ""
	socket.Type = ""

	socket.SetHost(host)
	socket.SetPort(port)

	return socket

}

func ParseSocket(value string) *Socket {

	var result *Socket = nil

	if strings.HasPrefix(value, "[") && strings.Contains(value, "]:") {

		tmp1 := value[0 : strings.Index(value, "]:")+1]
		tmp2, err := strconv.ParseInt(strings.Split(value, "]:")[1], 10, 16)

		if err == nil {
			socket := NewSocket(tmp1, uint16(tmp2))
			result = &socket
		}

	} else if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {

		socket := NewSocket(value, uint16(0))
		result = &socket

	} else if strings.Contains(value, ".") && strings.Contains(value, ":") {

		tmp1 := strings.Split(value, ":")[0]
		tmp2, err := strconv.ParseInt(strings.Split(value, ":")[1], 10, 16)

		if err == nil {
			socket := NewSocket(tmp1, uint16(tmp2))
			result = &socket
		}

	} else if strings.Contains(value, ":") {

		tmp1 := strings.Split(value, ":")[0]
		tmp2, err := strconv.ParseInt(strings.Split(value, ":")[1], 10, 16)

		if err == nil {
			socket := NewSocket(tmp1, uint16(tmp2))
			result = &socket
		}

	} else if strings.Contains(value, ".") {

		socket := NewSocket(value, uint16(0))
		result = &socket

	}

	return result

}

func ToSocket(value string) Socket {

	var socket Socket

	if value != "" {

		tmp := ParseSocket(value)

		if tmp != nil {
			socket = *tmp
		}

	}

	return socket

}

func (socket *Socket) IsValid() bool {

	var result bool

	if socket.Type == "ipv4" {

		if socket.Host == "*" {

			if socket.Port > 0 && socket.Port < 65535 {
				result = true
			}

		} else if socket.Host == "0.0.0.0" {

			if socket.Port > 0 && socket.Port < 65535 {
				result = true
			}

		} else if socket.Host != "0.0.0.0" {

			if socket.Port > 0 && socket.Port < 65535 {
				result = true
			}

		}

	} else if socket.Type == "ipv6" {

		if socket.Host == "*" {

			if socket.Port > 0 && socket.Port < 65535 {
				result = true
			}

		} else if socket.Host == "[0000:0000:0000:0000:0000:0000:0000:0000]" {

			if socket.Port > 0 && socket.Port < 65535 {
				result = true
			}

		} else if socket.Host != "[0000:0000:0000:0000:0000:0000:0000:0000]" {

			if socket.Port > 0 && socket.Port < 65535 {
				result = true
			}

		}

	} else if socket.Type == "domain" {

		if socket.Host != "" {

			if socket.Port > 0 && socket.Port < 65535 {
				result = true
			}

		}

	}

	return result

}

func (socket *Socket) SetHost(value string) {

	if IsIPv6(value) {

		ipv6 := ParseIPv6(value)

		if ipv6 != nil {
			socket.Host = ipv6.String()
			socket.Scope = ipv6.Scope()
			socket.Type = "ipv6"
		}

	} else if IsIPv4(value) {

		ipv4 := ParseIPv4(value)

		if ipv4 != nil {
			socket.Host = ipv4.String()
			socket.Scope = ipv4.Scope()
			socket.Type = "ipv4"
		}

	} else if IsDomain(value) {

		domain := ParseDomain(value)

		if domain != nil {
			socket.Host = domain.String()
			socket.Scope = "public"
			socket.Type = "domain"
		}

	} else if value == "*" {

		socket.Host = "*"
		socket.Scope = "public"
		socket.Type = "ipv4"

	}

}

func (socket *Socket) SetPort(value uint16) {

	// port 0 is used for host blocking
	if value >= 0 && value < 65535 {
		socket.Port = value
	}

}

func (socket *Socket) SetScope(value string) {

	if value == "private" || value == "public" {
		socket.Scope = value
	}

}

func (socket *Socket) String() string {

	result := ""

	if socket.Host != "" {
		result = socket.Host + ":" + strconv.FormatUint(uint64(socket.Port), 10)
	} else {
		result = "*:" + strconv.FormatUint(uint64(socket.Port), 10)
	}

	return result

}

func (socket Socket) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(socket.String())), nil
}

func (socket *Socket) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))

	if err != nil {
		return err
	}

	tmp := ParseSocket(unquoted)

	if tmp != nil {
		*socket = *tmp
	}

	return nil

}
