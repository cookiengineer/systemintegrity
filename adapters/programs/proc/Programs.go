package proc

import "github.com/cookiengineer/systemintegrity/caches"

var Programs *caches.Programs

func init() {
	Programs = caches.NewPrograms()
}
