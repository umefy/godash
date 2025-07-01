#! /bin/bash

BASE_PATH="$(cd "$(dirname "$0")/.." && pwd)"

lefthook install

cd "$BASE_PATH/scripts/commitlint" && pnpm install && cd "$BASE_PATH" || return
