#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")/.."

command_console() {
  echo '```console'
  echo "$ $*"
  "$@"
  echo '```'
}

commands() {
  for cmd in init run; do
    echo "
## pinact $cmd

$(command_console pinact help $cmd)"
  done
}

echo "# Usage

<!-- This is generated by scripts/generate-usage.sh. Don't edit this file directly. -->

$(command_console pinact help)
$(commands)" > USAGE.md
