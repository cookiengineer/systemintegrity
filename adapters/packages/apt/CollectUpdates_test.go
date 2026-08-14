//go:build debian || linuxmint || trisquel || ubuntu

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
