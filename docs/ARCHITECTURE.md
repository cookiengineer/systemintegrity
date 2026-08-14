# Architecture

This document is a precise and concise onboarding guide for the
**System Integrity Check** codebase. It explains the project layout, the
dependency direction, and every type, struct, enum, matcher, cache, adapter
and action so that new developers can navigate the code quickly.

Module path: `github.com/cookiengineer/systemintegrity` (Go 1.26+).

## 1. Goal

The application collects a local Linux system's state and compares it against
its distribution's package metadata. A user triggers a collection run and gets
a report of what is known, what is unexpected, and — for packages — exactly
which files have been changed compared to the package manager's records.

The first concrete deliverable is the **Packages** view: installed packages
(the **Package index**) and per-package integrity results (the
**PackageVerification index**).

## 2. Project layout

```
cmds/systemintegrity/main.go   # thin entrypoint: boots GTK, runs the app
gui/                           # all GTK4 GUI code
  Window.go                    # top-level window (sidebar + stack + refresh)
  views/                       # widget construction, no logic
  controllers/                 # data ownership + collection orchestration
bindings/gtk/                  # cgo bindings against GTK4
actions/                       # accumulates datasets into structs.System
adapters/<program>/            # program-specific collectors (pacman, df, ...)
structs/                       # unified cross-system data model
matchers/                      # query/match types used by caches and structs
types/                         # primitives (enums + value types)
caches/                        # in-memory indexes used by insights/*
insights/                      # embedded static datasets (devices, distros, countries)
utils/                         # small stdlib helpers (fmt, strings, ...)
docs/                          # MASTERPLAN, ARCHITECTURE, PACKAGES_SUPPORT, TODO
```

### Dependency direction (no cycles)

```
gui -> gui/views -> gui/controllers -> actions -> adapters -> structs
gui/views -> bindings/gtk
gui/controllers -> bindings/gtk
actions -> adapters -> structs
structs -> matchers -> types
caches  -> structs -> matchers -> types
insights -> caches -> structs
```

Rules:

- `structs` depends on `matchers` and `types`; `matchers` depends on `types`.
- Nothing outside `gui` and `cmds` may import GTK bindings.
- External dependencies are minimal (stdlib first); `CGo` is only required by
  `bindings/gtk`. `golang.org/x/term` and `golang.org/x/sys` are used for
  terminal detection.

## 3. Layer overview

| Layer      | Responsibility                                                              |
|------------|-----------------------------------------------------------------------------|
| `types/`   | Pure primitives: enums (`Manager`, `Architecture`, `Protocol`, `PackageVerificationIssue`) and value types (`Version`, `Datetime`, `Time`, IP/network types, `User`, `Group`, `Maintainer`). No logic beyond parsing/validation. |
| `matchers/`| "Query shapes" — lightweight structs with string fields that default to `"any"` and expose `Matches*` methods. Used for dependency/requirement expressions and cache queries. |
| `structs/` | The unified data model — the concrete records that make up a `System` report (`Package`, `Program`, `Device`, `PackageVerification`, ...). |
| `caches/`  | In-memory, mutex-guarded indexes keyed by an identifier string (`Packages`, `Updates`, `Verifications`, ...) with `Add/Get/Query/Remove/JSON`. |
| `adapters/`| Program-specific collectors that shell out to tools (`pacman`, `dpkg`, `rpm`, `df`, `/proc`, ...) and translate the raw output into `structs`. |
| `actions/` | Orchestration: `Init()` seeds a `*structs.System`, `Collect*()` functions fill its slices via the active adapter. |
| `insights/`| Static JSON datasets (devices, distributions, countries) embedded via `//go:embed` and wrapped by `caches`. |
| `gui/`     | GTK4 UI: `views/` build widgets, `controllers/` own the `System`/`Console` and marshal onto the GTK main loop. |

## 4. Enums (`types/`)

Enums are string-based `type X string` constants and follow a strict convention:
`IsX(string) bool`, `ParseX(string) *X`, `ToX(string) X`, `X.String()`,
`(*X).IsValid()`, `X.MarshalJSON`, `(*X).UnmarshalJSON`. `ParseX` returns `nil`
for unknown input; `ToX` returns the zero value.

### `Manager`

The package manager (or packaging ecosystem). Valid values: `any`, `apk`,
`apt`, `dnf`, `pacman`, `rpm`, `tdnf`, `zypper` (distributions); `pkg`,
`pkgsrc`, `msi` (other OS); plus programming-language managers (`cargo`,
`composer`, `npm`, `pip`, `gem`, ...).

### `Architecture`

CPU architecture, normalized to a small set: `any`, `x86`, `x86_64`, `armv6`,
`armv7`, `armv8`, `riscv32`, `riscv64`, `sparc`, `sparc64`. `ParseArchitecture`
maps common spellings (`amd64`/`x86_64` → `x86_64`, `armhf` → `armv7`, ...).

### `Protocol`

Network/transport protocol: `any`, `dns`, `dns-over-tls`, `http`, `https`,
`icmp`, `ssh`, `socks`, `tcp`, `udp`, `whois`.

### `PackageVerificationIssue`

The package-manager-agnostic enum for everything that can diverge between an
installed file and its package metadata. Full cross-manager mapping lives in
`docs/PACKAGES_SUPPORT.md`; the values are:

| Value                          | Meaning                                                    |
|--------------------------------|------------------------------------------------------------|
| `checksum_mismatch`            | file digest differs from metadata                         |
| `checksum_unavailable`         | digest could not be computed/verified                     |
| `device_mismatch`              | device major/minor differs (rpm only)                     |
| `file_type_mismatch`           | path is no longer a regular file (dpkg `M` heuristic)     |
| `group_mismatch`               | group ownership differs                                   |
| `mode_mismatch`                | permission/mode bits differ                               |
| `modification_time_mismatch`   | mtime differs — **non-issue** (`IsIssue()` returns false) |
| `missing_file`                 | file absent from filesystem                               |
| `permission_denied`            | file could not be read                                    |
| `readlink_mismatch`            | symlink target path differs                               |
| `size_mismatch`                | size differs                                              |
| `user_mismatch`                | user ownership differs                                    |
| `capabilities_mismatch`        | file capabilities differ (rpm only)                       |

Important distinction: dpkg's `M` is *not* the same as pacman's "Permissions
mismatch" or rpm's `M`. dpkg does not track permission bits; its `M` only fires
when a path that should be a regular file is no longer one, so it maps to
`file_type_mismatch`, whereas pacman/rpm permission changes map to
`mode_mismatch`.

Methods on the enum: `IsValid`, `String`, `IsIssue`, `Description`,
`ParsePackageVerificationIssue`, `ToPackageVerificationIssue`, JSON marshaling.

## 5. Value types (`types/`)

### `Version`

Rich semantic version parser. Fields: `Epoche uint`, `Upstream{ Major, Minor,
Patch uint; Release, Hash string }`, `Revision string`. Supports epoch (`1:`),
`major.minor.patch`, prerelease tags, hashes, and revisions. Methods:
`Parse`, `IsValid`, `IsAfter/IsBefore/IsSame`, `Next*`/`Prev*`,
`SemanticString`, `String` (`0:1.2.3alpha+hash~rev`), `ToEarlier/ToEarliest/
ToLater/ToLatest`.

### `Datetime`

Calendar + wall-clock struct (`Year, Month, Day, Hour, Minute, Second uint`).
`Parse` accepts many formats (ISO-8601, RFC-ish, `YYYY-MM-DD`, timezone
offsets); `String()` emits `YYYY-MM-DD HH:MM:SS`. Comparison helpers
(`IsAfter/IsBefore/IsSame`, `ToDatetimeDifference`, `ToTimeDifference`,
`ToWeekday`).

### `Time`

Wall-clock only (`Hour, Minute, Second uint`). Arithmetic (`AddSecond`, ...),
comparison, `Offset` (timezone), `String` (`HH:MM:SS`).

### `Prefix`

CIDR prefix length (`uint8`). `String()` renders `/24`. `ParsePrefix` handles
`/24`, `[::]/64`, `::/64`, `1.2.3.0/24`, bare numbers, etc.

### Network types

- `IPv4 [4]byte`, `IPv6 [16]byte` — with `Is*`, `Parse*`, `Parse*AndPort`.
- `Domain string` — hostname validation/parsing.
- `ASN int` — autonomous system number.
- `Socket`, `Server`, `Connection`, `Geolocation` — structured network/geo
  primitives used by `matchers` and the connection cache.

### Identity types

- `User { ID, Name, Password, Folder, Groups []Group, Shell, Type }`.
- `Group { ID, Name, Password, Type }`.
- `Maintainer { Name, Email }` — parses `Name <email>`, `<email>`, `email`,
  or bare `Name`; `String()` re-renders `Name <email>`.

## 6. Matchers (`matchers/`)

Matchers are query shapes: structs whose fields are strings defaulting to
`"any"` and which expose `Matches*(value) bool`. `"any"` matches everything.
They are used for dependency/requirement expressions inside `structs.Package`
and for cache queries.

Key matchers:

- `Package { Name, Version, Architecture, Manager, Vendor }` — the most used
  one. `Parse` splits a requirement string (via `parseVersionCondition`) into
  name + version condition + architecture. `MatchesVersion` understands
  `>=`, `>`, `<=`, `<`, `=` operators. `Hash()` computes a CRC32 id.
- `Unresolved { Candidates []Package }` — represents `a | b` or `a || b`
  alternative dependencies before they are resolved against installed packages.
- `Manager`, `User`, `Connection`, `Country`, `Device`, `Distribution`,
  `Drive`, `Network`, `Program`, `Subnet`, `Timezone`, `Datetime`, `Time`,
  `Timeslot`, `Update`, `Antique`, `System` — matcher counterparts of the
  concrete structs, plus security-domain matchers (`Vulnerability`,
  `Weakness`, `Incident`, `Mitigation`, `Response`, `Credential`).

## 7. Structs (`structs/`)

The unified data model. Every struct follows a convention: a `New*`
constructor initializing slices, `Set*`/`Add*/Remove*` mutators that validate
and deduplicate, and `IsValid()` (plus `IsIdentical` where needed).

### `System` (the root report)

```go
type System struct {
    Name, Hostname        string
    Datetime              types.Datetime
    Distribution          Distribution
    Fingerprint           struct{ Country, Locale, Timezone, Token string }
    BIOS, Board           Device
    Boot                  Boot
    Devices               []Device
    Drives                []Drive
    Networks              []Network
    Packages              []Package            // Package index
    Programs              []Program
    Services              []Program
    Antiques              []Antique
    Updates               []Update
    Users                 []types.User
    Verifications         []PackageVerification // PackageVerification index
}
```

`ToJSON()` renders the whole report as indented JSON.

### `Distribution`

`Name, Version, Kernel, KernelArchitecture, KernelModules []string,
KernelVersion, Manager, Vendor, Keywords *map[string]string`. Kernel strings
are normalized (`linux`, `freebsd`, ...).

### `Device` / `DeviceSystem`

`DeviceSystem { Name, Vendor, Device }` (a vendor:device tuple). `Device
{ Name, Bus, System *DeviceSystem, Subsystem *DeviceSystem }` where `Bus` is
one of `pci`, `usb`, `hid`, `i2c`, `scsi`, `other`.

### `Drive`

`Name, Mountpoint, Type, Size, Free` (plus `ToJSON`, `IsValid`). `Size`/`Free`
are human-formatted via `utils/fmt.FormatBytes`.

### `Network`

A network interface (name, IP addresses, subnet, protocol, ...).

### `Package`

The heart of the Package index:

```go
type Package struct {
    Name, Version, Architecture, Manager, Vendor, URL, Datetime
    Maintainers  []types.Maintainer
    Filesystem   []string            // absolute paths owned by the package
    Conflicts    []matchers.Package
    Dependencies []matchers.Package
    Provides     []matchers.Package
    Replaces     []matchers.Package
    Unresolved   []matchers.Unresolved
}
```

`ResolveDependencies(packages []Package)` resolves `Unresolved` alternatives
against the collected set. `HasFilesystem(path)` is used to attribute running
processes to owning packages.

### `PackageVerification` / `FileVerification`

The PackageVerification index:

```go
type FileVerification struct {
    Path         string
    Issues       []types.PackageVerificationIssue
    Remediations []types.Remediation
}

type PackageVerification struct {
    Name    string
    Version types.Version
    Manager types.Manager
    Files   []FileVerification
}
```

One file can carry multiple issues (e.g. rpm `S.5....T.` → size + digest +
mtime). `types.Remediation { Manager, Issue, Command, Description }` stores the
advisory fix command per issue.

### `Program`

A running process (`PID, Name, Command, Arguments, Folder, Environment, User,
Manager, Filesystem, Connections, Dependencies, Packages`). `IsProgram()` /
`IsService()` distinguish clients from listeners based on connection type.
`ResolveDependencies` maps filesystem paths back to owning packages.

### `Antique`

A service dependency ("what keeps an old service running"): `Name, Version,
Architecture, Manager, Vendor, Service, URL`.

### `Boot`

`Bootloader, Mode, Kernel, KernelArchitecture, KernelVersion, Initramfs,
SecureBoot, ESP` — populated by `adapters/boot/*` (skeleton).

### `Update`

`Name, Version, Architecture, Manager, Vendor, URL` — an available upgrade.

### `Console` / `ConsoleMessage`

A thread-safe, color-aware log with `Group/GroupEnd`, `Log/Info/Warn/Error`,
`Progress`, `Inspect`. `Snapshot()` returns a copy of `Messages`; the GUI
derives progress from the grouped messages. `ConsoleMessage { Method, Value }`.

## 8. Caches (`caches/`)

Thread-safe in-memory indexes (mutex-guarded maps) keyed by a deterministic
identifier string. All implement `New*`, `Add`, `Get`, `Length`, `Query`,
`Remove`, and custom `MarshalJSON`/`UnmarshalJSON`.

- `Packages` — keyed `manager:vendor:name:version:architecture`, with a
  dependency reverse-index for `QueryByDependency`.
- `Updates` — keyed `manager:vendor:name:version:architecture`.
- `Verifications` — keyed `manager:name:version` (the PackageVerification index).
- `Programs`, `Antiques`, `Networks`, `Devices`, `Countries`, `Distributions`,
  `Timezones` — additional indexes; the last three wrap `insights/*` JSON.

## 9. Adapters (`adapters/`)

Adapters shell out to external programs and translate output into `structs`.
Each adapter package has a `<name>.go` file with `init()` that sets `SUPPORTED`
(and sometimes `OPTIMIZED`) by checking for the tool's binary.

Package manager adapters (`adapters/packages/<manager>`):

| Manager | Files                                     | Notes |
|---------|-------------------------------------------|-------|
| `pacman`| `CollectPackages`, `CollectPackage`, `CollectUpdates`, `CollectUpdate`, `CollectVerification`, `ParsePackage`, `ParseUpdate` | `pacman -Qi/-Ql/-Qkk`. |
| `apt`   | `CollectPackages`, `CollectPackage`, `CollectUpdates`, `CollectUpdate`, `CollectVerification`, `ParsePackage`, `ParseUpdate`, `ParseVerificationLine` | `apt list`, `apt-cache show`, `dpkg --verify`. |
| `apk`   | `CollectPackages`, `CollectPackage`, `CollectUpdates`, `CollectUpdate`, `parsePackageIndex` | Alpine `apk`; parses `/lib/apk/db/installed`-style index. |
| `rpm`   | `CollectPackages`, `ParsePackage`, `ParseConflict/Provide/Require/SharedLibrary`, `CollectVerification`, `ParseVerificationLine` | `rpm -qa/-V`; shared parsing reused by dnf/zypper/tdnf. |
| `dnf`   | `CollectPackages`, `CollectPackage`, `CollectUpdates`, `CollectUpdate`, `CollectVerification` | delegates dependency parsing to `rpm`. |
| `zypper`| `CollectPackages`, `CollectPackage`, `CollectUpdates`, `CollectUpdate`, `CollectVerification` | delegates dependency parsing to `rpm`. |
| `tdnf`  | stub                                        | TODO. |

**Verification adapters** all expose `CollectVerification() []structs.PackageVerification`
and map native output to `types.PackageVerificationIssue`:

- pacman: `pacman -Qkk --noconfirm` → `pkg: /path (reason)` lines.
- apt: `dpkg --verify <pkg>` per package → `??5?????? [c] path` / `missing path`.
- rpm: `rpm -V <pkg>` per package → flag run (`S.5....T.`) + path; `missing`.

Other adapters: `boot/{bootctl,mkinitcpio}`, `devices/sys`, `drives/df`,
`programs/{ldd,proc}`, `system/{coreutils,etc,proc,sys}`, `users/{shadow,systemd}`.

## 10. Actions (`actions/`)

`Init(console) *structs.System` seeds the system (machine-id, hostname,
kernel, BIOS/board). `Collect(console, system)` runs the collection steps:

```
CollectBoot, CollectDrives, CollectDevices, CollectBoot, CollectNetworks,
CollectPrograms, CollectServices, CollectPackages, CollectVerifications,
CollectUpdates, CollectUsers, CollectAntiques
```

Each `Collect*` function picks the first `SUPPORTED` adapter (mirroring the
`pacman → apt → apk → rpm → dnf → tdnf → zypper` ordering), fills the relevant
`system` slice via `Set*`, and logs progress through the `Console`. The GUI
progress bar advances one of 12 steps per top-level group
(`gui/controllers/progress.go`, `CollectSteps = 12`).

## 11. GUI (`gui/` + `bindings/gtk/`)

- `bindings/gtk` — cgo bindings over GTK4.
- `gui/Window.go` — assembles `GtkStackSidebar` + `GtkStack`; disables the
  sidebar until the first collection completes; `Refresh()` repopulates views.
- `gui/views/*` — pure widget construction exposing `Refresh(*structs.System)`;
  never block the main loop.
- `gui/controllers/App.go` — owns `*structs.System` and `*structs.Console`,
  runs `actions.Init`+`actions.Collect` (or re-verification) in goroutines, and
  marshals UI updates via `gtk.RunOnMain`.

## 12. Testing

- **Host unit tests** (`make test`): `types`, `structs`, `caches` — no
  package manager or container required. Includes full coverage of every
  `types.PackageVerificationIssue` enum value and its remediation.
- **Adapter parser tests**: `adapters/packages/*/CollectVerification_test.go`
  are guarded by build tags (`antergos || archlinux || manjaro`, `debian ||
  linuxmint || trisquel || ubuntu`, `redhat || ... || fedora || amazonlinux`,
  `opensuse || suse_desktop || suse_server`) and test the parsers against
  representative native output strings.
- **Container integration tests** (`make test-integration` →
  `./test-integration.sh`): builds tagged test binaries on the host and runs
  them inside podman containers (archlinux, ubuntu, fedora, opensuse). Each
  `actions/Dockerfile.<platform>` installs `openssh` and tampers with its files
  so the tagged `actions/CollectVerifications_<platform>_test.go` asserts the
  expected enum per tampered file. See `docs/PACKAGES_SUPPORT.md`.

## 13. Conventions

- Package-manager specificity is isolated in `adapters/`; `structs`, `matchers`,
  and `types` never import adapters.
- All structs/types serialize cleanly to JSON (explicit `MarshalJSON` /
  `UnmarshalJSON` where custom rendering is needed).
- Parsing helpers are total (never panic) and return `nil`/zero on failure.
- Slices are always initialized (non-nil) in constructors.
- "any" is the universal wildcard across enums and matchers.

## 14. Further reading

- `docs/MASTERPLAN.md` — original architecture and roadmap.
- `docs/PACKAGES_SUPPORT.md` — Packages/PackageVerification plan, enum mapping,
  remediation matrix, and container test details.
- `docs/TODO.md` — task tracking.
