# System Integrity Check for Linux — Master Plan

This document describes the architecture, conventions, and the implementation
plan for the **System Integrity Check** GTK4 GUI and its integrity checkers.

## 1. Goal

Show the local system's integrity compared to its distribution's package
metadata. A user opens the GUI, triggers a collection run, and gets a report of
what is known, what is unexpected, and — for packages — exactly which files
have been changed compared to the package manager's records.

The first concrete deliverable is a **Package Integrity** view that collects
installed packages and compares their files against the package manager's
metadata (e.g. `pacman -Qkk`), producing a list of affected packages with an
expandable per-package view of the changed files.

## 2. Project layout & dependency direction

```
cmds/systemintegrity/main.go   # thin entrypoint: boots GTK, runs the app
gui/                           # all GTK4 GUI code
  views/                       # widget construction, no logic
  controllers/                 # data ownership + collection orchestration
bindings/gtk/                  # cgo bindings against GTK4 (GTK4-compatible)
actions/                       # accumulates datasets into structs.System
adapters/<program>/            # program-specific collectors (pacman, df, ...)
structs/                       # unified cross-system data model
matchers/                      # query/match types used by caches and structs
types/                         # primitives (Version, Architecture, Manager, ...)
caches/                        # in-memory indexes used by insights/*
insights/                      # embedded static datasets (devices, distributions, countries)
utils/                         # small stdlib helpers (strings, ...)
```

**Dependency direction (no cycles):**

```
gui/controllers -> gui/views -> bindings/gtk
gui/controllers -> actions -> adapters -> structs
actions -> adapters -> structs
structs -> matchers -> types
caches  -> structs -> matchers -> types
insights -> caches -> structs
```

- `structs` depends on `matchers` and `types`.
- `matchers` depends on `types`.
- Nothing outside `gui` and `cmds` may import GTK bindings.
- External dependencies are kept to a minimum (stdlib first); `CGo` is only
  required for the GTK4 bindings.

## 3. Report dataset

`structs/System.go` (`structs.System`) is the single report dataset used by both
the GUI and the collectors. `actions.Init()` creates and seeds it, then
`actions.Collect()` fills the following slices:

| Field         | Type            | Source                                  |
|---------------|-----------------|-----------------------------------------|
| Distribution  | `Distribution`  | `adapters/system/etc` + `coreutils`     |
| BIOS / Board  | `Device`        | `adapters/system/sys`                   |
| Devices       | `[]Device`      | `adapters/devices/sys`                  |
| Drives        | `[]Drive`       | `adapters/drives/df`                    |
| Networks      | `[]Network`     | `actions` (stdlib `net`)                |
| Packages      | `[]Package`     | `adapters/packages/<manager>`           |
| Programs      | `[]Program`     | `adapters/programs/proc` + `ldd`        |
| Services      | `[]Program`     | `adapters/programs/proc`                |
| Updates       | `[]Update`      | `adapters/packages/<manager>`           |
| Users         | `[]types.User`  | `adapters/users/shadow` + `systemd`     |
| Antiques      | `[]Antique`     | derived in `actions/CollectAntiques`    |
| Verifications | `[]PackageVerification` | `adapters/packages/pacman/CollectVerification` |
| Boot          | `Boot`          | `adapters/boot/*` (skeleton)            |

## 4. GUI architecture

### 4.1 Sidebar

A `GtkStack` + `GtkStackSidebar` (native GTK4 sidebar) with four pages:

1. **Welcome** — landing page with a "Collect" button and a live progress/log.
2. **Package Integrity** — the primary integrity report.
3. **Boot Integrity** — skeleton (bootctl / mkinitcpio planned).
4. **Hardware Integrity** — read-only hardware inventory.

### 4.2 Controllers vs views

- **views/** build widgets and expose small update methods; they contain no
  collection logic and never block the main loop.
- **controllers/** own the `*structs.System` and `*structs.Console`, spawn
  background goroutines for collection, and marshal UI updates onto the GTK
  main thread via `gtk.RunOnMain`.

### 4.3 Threading & progress

Collection shells out to external programs and reads `/proc`/`/sys`, so it must
never run on the GTK main loop:

```
goroutine: actions.Init + actions.Collect  (console messages accumulate)
ticker:    console.Snapshot() -> gtk.RunOnMain -> update Welcome view
```

Progress is derived from the grouped `console.Messages` slice: the current step
is the deepest open `Group`, completed steps are closed top-level groups.
A `ProgressBar` pulses while the scan is running.

## 5. Package Integrity (primary feature)

1. `actions.CollectPackages` (already implemented) collects every package via
   the active package-manager adapter (pacman on Arch Linux).
2. `adapters/packages/pacman/CollectVerification.go` runs `pacman -Qkk
   --noconfirm` and parses `pkgname: /path (reason)` lines on stderr into
   `[]structs.PackageVerification`.
3. `actions/CollectVerifications.go` aggregates across managers and stores the
   result in `system.Verifications`.
4. The **Package Integrity** view renders a summary header and a
   `GtkScrolledWindow` of `GtkExpander` rows — one per affected package —
   whose child lists the changed files and the detected reason.

Because `-Qkk` is a full-filesystem checksum scan, verification is **not** part
of the general `actions.Collect`; the view triggers it on demand via a
"Verify" button.

## 6. Boot Integrity (skeleton + plan)

- `structs/Boot.go` — unified boot dataset (`Bootloader`, `Mode`, `Kernel`,
  `KernelVersion`, `Initramfs`, `SecureBoot`, `ESP`).
- `adapters/boot/bootctl/CollectBoot.go` — parse `bootctl status` / `bootctl list`.
- `adapters/boot/mkinitcpio/CollectInitramfs.go` — parse `/etc/mkinitcpio.conf`,
  `/etc/mkinitcpio.d/*.preset`, `/usr/lib/initcpio/`.
- `actions/CollectBoot.go` — aggregate into `system.Boot`.

Implemented as skeletons in the first pass; wired into the Boot view later.

## 7. Hardware Integrity

Read-only inventory from already-collected data: `system.BIOS`, `system.Board`,
`system.Devices` (name/bus/vendor:device), and `system.Drives`
(name/mountpoint/type/size/free). A baseline-comparison mode is planned for a
later milestone.

## 8. GTK bindings to add

- `Stack.go` — `GtkStack`: `NewStack`, `AddTitled`, `SetVisibleChildName`.
- `StackSidebar.go` — `GtkStackSidebar`: `NewStackSidebar`, `SetStack`.
- `Expander.go` — `GtkExpander`: `NewExpander`, `SetChild`, `SetExpanded`, `SetLabel`.
- `ProgressBar.go` — `GtkProgressBar`: `NewProgressBar`, `SetFraction`, `Pulse`, `SetText`.

## 9. Definition of done

- `go build ./...` and `go vet ./...` clean.
- `go test -tags guard_archlinux ./adapters/packages/pacman/...` passes.
- App launches, Welcome view collects with live progress, Package Integrity
  lists affected packages with expandable changed-file details, Boot and
  Hardware views render (skeleton/inventory).

## 10. Roadmap (future integrity checks)

- Orphaned packages and downgraded/version anomalies.
- Config drift vs `.pacnew`/`.rpmsave`.
- SUID/setuid and world-writable binaries.
- Running-process ↔ package ownership mismatch.
- Open listening ports and unexpected connections.
- Kernel module + initramfs hash verification.
- Bootloader / EFI / Secure Boot / TPM measurements.
- Immutable (`chattr +i`) and capability (`setcap`) audit.
- Persisted baseline for "changed since last run".
