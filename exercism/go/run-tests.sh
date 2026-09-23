#!/usr/bin/env bash
# Usage:
#   ./run-tests.sh <exercise-dir> [TestName] [-v]
#
# Examples:
#   ./run-tests.sh logs-logs-logs                   # run all tests, quiet
#   ./run-tests.sh logs-logs-logs -v                # run all tests, verbose
#   ./run-tests.sh logs-logs-logs TestReplace        # run one test, quiet
#   ./run-tests.sh logs-logs-logs TestReplace -v     # run one test, verbose

set -euo pipefail

VERBOSE=""
ARGS=()

for arg in "$@"; do
  if [ "$arg" = "-v" ]; then
    VERBOSE="-v"
  else
    ARGS+=("$arg")
  fi
done

if [ "${#ARGS[@]}" -lt 1 ]; then
  echo "Usage: $0 <exercise-dir> [TestName] [-v]"
  exit 1
fi

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
EXERCISE="${ARGS[0]}"
TEST_NAME="${ARGS[1]:-}"

if [ -n "$TEST_NAME" ]; then
  go test -C "$SCRIPT_DIR/$EXERCISE" -run "$TEST_NAME" $VERBOSE ./...
else
  go test -C "$SCRIPT_DIR/$EXERCISE" $VERBOSE ./...
fi
