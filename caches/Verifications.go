package caches

import "github.com/cookiengineer/systemintegrity/structs"
import "encoding/json"
import "sort"
import "sync"

func toVerificationIdentifier(value structs.PackageVerification) string {
	return value.Manager.String() + ":" + value.Name + ":" + value.Version.String()
}

type Verifications struct {
	Verifications map[string]*structs.PackageVerification `json:"verifications"`
	mutex         *sync.RWMutex                           `json:"-"`
}

func NewVerifications() *Verifications {

	var cache Verifications

	cache.Verifications = make(map[string]*structs.PackageVerification)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Verifications) Add(value structs.PackageVerification) bool {

	cache.mutex.Lock()

	var result bool

	if value.Manager.IsValid() && value.Name != "" && len(value.Files) > 0 {

		identifier := toVerificationIdentifier(value)

		cache.Verifications[identifier] = &value
		result = true

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Verifications) Get(identifier string) *structs.PackageVerification {

	cache.mutex.RLock()

	var result *structs.PackageVerification = nil

	if identifier != "" {

		verification, ok := cache.Verifications[identifier]

		if ok == true {
			result = verification
		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Verifications) Length() int {
	return len(cache.Verifications)
}

func (cache *Verifications) Remove(identifier string) bool {

	cache.mutex.Lock()

	var result bool

	if identifier != "" {

		_, ok := cache.Verifications[identifier]

		if ok == true {

			delete(cache.Verifications, identifier)
			result = true

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Verifications) MarshalJSON() ([]byte, error) {

	cache.mutex.RLock()

	tmp := make(map[string]*structs.PackageVerification)
	identifiers := make([]string, 0)

	for identifier, _ := range cache.Verifications {
		identifiers = append(identifiers, identifier)
	}

	sort.Strings(identifiers)

	for _, identifier := range identifiers {
		tmp[identifier] = cache.Verifications[identifier]
	}

	cache.mutex.RUnlock()

	return json.MarshalIndent(&struct {
		Verifications map[string]*structs.PackageVerification `json:"verifications"`
	}{
		Verifications: tmp,
	}, "", "\t")

}

func (cache *Verifications) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Verifications map[string]*structs.PackageVerification `json:"verifications"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, verification := range tmp.Verifications {
		cache.Add(*verification)
	}

	return nil

}
