package matchers

import "strings"

type Manager struct {
	Name string `json:"name"`
}

func NewManager() Manager {

	var manager Manager

	manager.Name = "any"

	return manager

}

func ToManager(value string) Manager {

	var manager Manager

	manager.Name = "any"

	manager.SetName(value)

	return manager

}

func (manager *Manager) IsIdentical(value Manager) bool {

	var result bool

	if manager.Name == value.Name {
		result = true
	}

	return result

}

func (manager *Manager) IsValid() bool {

	var result bool

	if manager.Name != "any" {
		result = true
	}

	return result

}

func (manager *Manager) Matches(name string) bool {
	return manager.MatchesName(name)
}

func (manager *Manager) MatchesName(value string) bool {

	var result bool

	if manager.Name == value {
		result = true
	} else if manager.Name == "any" {
		result = true
	}

	return result

}

func (manager *Manager) SetName(value string) {
	manager.Name = strings.TrimSpace(value)
}

func (manager *Manager) Hash() string {

	var hash string

	if manager.Name != "" {
		hash = manager.Name
	}

	return hash

}
