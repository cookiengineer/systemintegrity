package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Program struct {
	Name      string   `json:"name"`
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
}

func NewProgram() Program {

	var program Program

	program.Name = "any"
	program.Command = "any"
	program.Arguments = make([]string, 0)

	return program

}

func ToProgram(value string) Program {

	var program Program

	program.Name = "any"
	program.Command = "any"
	program.Arguments = make([]string, 0)

	if strings.HasPrefix(value, "/") {
		program.Name = "any"
		program.SetCommand(value)
	} else {
		program.SetName(value)
		program.Command = "any"
	}

	return program

}

func (program *Program) IsIdentical(value Program) bool {

	var result bool

	if program.Name == value.Name && program.Command == value.Command {

		if len(program.Arguments) == len(value.Arguments) {

			matches_args := true

			for a, arg := range program.Arguments {

				if arg == value.Arguments[a] {
					// Do Nothing
				} else {
					matches_args = false
					break
				}

			}

			if matches_args {
				result = true
			}

		}

	}

	return result

}

func (program *Program) IsValid() bool {

	var result bool

	if program.Name != "any" || program.Command != "any" {
		result = true
	}

	return result

}

func (program *Program) Matches(name string, command string, arguments []string) bool {
	return program.MatchesName(name) && program.MatchesCommand(command) && program.MatchesArguments(arguments)
}

func (program *Program) MatchesArguments(values []string) bool {

	var result bool

	if len(program.Arguments) == len(values) {

		matches_args := true

		for a, arg := range program.Arguments {

			if arg == values[a] {
				// Do Nothing
			} else if arg == "any" {
				// Do Nothing
			} else {
				matches_args = false
				break
			}

		}

		if matches_args {
			result = true
		}

	} else if len(program.Arguments) == 0 {
		result = true
	}

	return result

}

func (program *Program) MatchesCommand(value string) bool {

	var result bool

	if program.Command == value {
		result = true
	} else if program.Command == "any" {
		result = true
	}

	return result

}

func (program *Program) MatchesName(value string) bool {

	var result bool

	if program.Name == value {
		result = true
	} else if program.Name == "any" {
		result = true
	}

	return result

}

func (program *Program) SetArguments(values []string) {
	program.Arguments = values
}

func (program *Program) SetCommand(value string) {

	if value == "all" || value == "any" || value == "*" {
		program.Command = "any"
	} else if value != "" {
		program.Command = strings.TrimSpace(value)
	}

}

func (program *Program) SetName(value string) {

	if value == "all" || value == "any" || value == "*" {
		program.Name = "any"
	} else if strings.Contains(value, "/") {
		program.Name = strings.TrimSpace(value[0:strings.LastIndex(value, "/")])
	} else if value != "" {
		program.Name = strings.TrimSpace(value)
	}

}

func (program *Program) Hash() string {

	var hash string

	if program.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			program.Name,
			program.Command,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
