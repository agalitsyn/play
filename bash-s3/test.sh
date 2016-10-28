#!/usr/bin/env bash

set -eo pipefail; [[ $TRACE ]] && set -x

SCRIPT_DIR="$(dirname "$0")"

main() {
    local usage="$0 <bucket-name> <object-name>"
    local bucket=${1:?$usage}
    local object=${2:?$usage}

    source "$SCRIPT_DIR/../../../mietek/bashmenot/src.sh"

    export BASHMENOT_AWS_ACCESS_KEY_ID="$AWS_ACCESS_KEY_ID"
    export BASHMENOT_AWS_SECRET_ACCESS_KEY="$AWS_SECRET_ACCESS_KEY"
    export BASHMENOT_S3_ENDPOINT="$AWS_S3_ENDPOINT_URL"

    s3_check "$bucket" "$object"
}

main "$@"
