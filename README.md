# System Integrity Check

GTK4 application to validate the integrity of packages, boot setup, and devices.

![Screenshot](./docs/screenshot.png)

## Views

| View     | Requirements                                                     | Description                                          | Ready? |
|:---------|:-----------------------------------------------------------------|:-----------------------------------------------------|--------|
| Antiques | (Packages and Services requirements)                             | Integrity Check for dependencies of running services |        |
| Boot     | `bootctl`, `mkinitcpio`                                          | Integrity Check for boot betup                       | yes    |
| Drives   | `df`                                                             | Integrity Check for mounts and hard drives           |        |
| Packages | Either of `apk`, `apt`, `dnf`, `pacman`, `rpm`, `tdnf`, `zypper` | Integrity Check for packages                         | yes    |
| Programs | `ldd`, `/proc`                                                   | Integrity Check for running programs                 |        |
| Services | `ldd`, `/proc`                                                   | Integrity Check for running services                 |        |
| System   | `coreutils`, `/etc`, `/sys`                                      | Integrity Check for system-wide configurations       |        |
| Users    | `shadow`, `systemd`                                              | Integrity Check for system-wide user accounts        |        |

## Build

```bash
# Requirements: Go 1.26+, GTK 4.22+ development headers
make build;
```

## Install

```bash
sudo make install;
```

## License

[AGPL-3.0](./LICENSE)

