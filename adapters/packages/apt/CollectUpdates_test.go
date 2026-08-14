//go:build guard_debian || guard_linuxmint || guard_trisquel || guard_ubuntu || intel_debian || intel_linuxmint || intel_trisquel || intel_ubuntu

package apt

import "github.com/cookiengineer/systemintegrity/structs"
import "slices"
import "testing"

func TestCollectUpdates(t *testing.T) {

	t.Run("CollectUpdates()", func(t *testing.T) {

		updates := CollectUpdates()

		if slices.ContainsFunc[[]structs.Update, structs.Update](updates, func(update structs.Update) bool {
			return update.Name != "" && update.Version.IsValid()
		}) == false {
			t.Errorf("Expected %v to contain %s", updates, "Name=any")
		}

	})

}
