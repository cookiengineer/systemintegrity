package controllers

import "github.com/cookiengineer/systemintegrity/bindings/gtk"
import "github.com/cookiengineer/systemintegrity/actions"
import "github.com/cookiengineer/systemintegrity/structs"
import "sync"
import "time"

type App struct {
	console    *structs.Console
	system     *structs.System
	mutex      sync.RWMutex
	collecting bool
	verifying  bool
}

func NewApp() *App {

	var app App

	system := structs.NewSystem()

	app.console = structs.NewConsole(nil, nil, 0)
	app.system = &system
	app.collecting = false
	app.verifying = false

	return &app

}

func (app *App) Console() *structs.Console {

	app.mutex.RLock()
	console := app.console
	app.mutex.RUnlock()

	return console

}

func (app *App) System() *structs.System {

	app.mutex.RLock()
	system := app.system
	app.mutex.RUnlock()

	return system

}

func (app *App) IsCollecting() bool {

	app.mutex.RLock()
	collecting := app.collecting
	app.mutex.RUnlock()

	return collecting

}

func (app *App) IsVerifying() bool {

	app.mutex.RLock()
	verifying := app.verifying
	app.mutex.RUnlock()

	return verifying

}

func (app *App) StartCollect(onProgress func(Progress), onDone func()) {

	app.mutex.Lock()

	if app.collecting == true {
		app.mutex.Unlock()
		return
	}

	app.collecting = true

	console := structs.NewConsole(nil, nil, 0)
	app.console = console

	app.mutex.Unlock()

	done := make(chan struct{})

	go func() {

		system := actions.Init(console)
		actions.Collect(console, system)

		app.mutex.Lock()
		app.system = system
		app.collecting = false
		app.mutex.Unlock()

		close(done)

		gtk.RunOnMain(onDone)

	}()

	go func() {

		ticker := time.NewTicker(150 * time.Millisecond)
		defer ticker.Stop()

		for {

			select {

			case <-done:
				return

			case <-ticker.C:

				progress := ProgressOf(console.Snapshot())

				gtk.RunOnMain(func() {
					onProgress(progress)
				})

			}

		}

	}()

}

func (app *App) StartVerification(onProgress func(string), onDone func()) {

	app.mutex.Lock()

	if app.verifying == true {
		app.mutex.Unlock()
		return
	}

	app.verifying = true

	console := structs.NewConsole(nil, nil, 0)

	app.mutex.Unlock()

	done := make(chan struct{})

	go func() {

		system := app.System()

		actions.CollectVerifications(console, system)

		app.mutex.Lock()
		app.verifying = false
		app.mutex.Unlock()

		close(done)

		gtk.RunOnMain(onDone)

	}()

	go func() {

		ticker := time.NewTicker(250 * time.Millisecond)
		defer ticker.Stop()

		for {

			select {

			case <-done:
				return

			case <-ticker.C:

				step := CurrentStep(console.Snapshot())

				gtk.RunOnMain(func() {
					onProgress(step)
				})

			}

		}

	}()

}
