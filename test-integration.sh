#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT"

mkdir -p "${ROOT}/build";

# test_platform <platform>
#
# Builds the tagged test binary for ./actions on the host, ships it into the
# platform's container image, and runs it inside a podman container. The build
# tag matches the platform name.
test_platform() {

	local platform="$1"
	local binary="build/actions_${platform}.test";
	local image="systemintegrity-test-${platform}";
	local dockerfile="actions/Dockerfile.${platform}";

	echo "==> [${platform}] building tagged test binary (-tags ${platform})";
	CGO_ENABLED=0 go test -c -tags "${platform}" -o "${binary}" ./actions;

	echo "==> [${platform}] building container image";
	podman build -f "${dockerfile}" -t "${image}" .;

	echo "==> [${platform}] running container tests";
	podman run --rm "${image}";

}

test_platform archlinux;
test_platform ubuntu;
test_platform fedora;
test_platform opensuse;
test_platform alpinelinux;

echo "All container tests passed."

