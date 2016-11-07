#!/usr/bin/env bash

set -eo pipefail

SCRIPT_DIR="$(dirname "$0")"

main() {
    local usage="$0 <bucket-name> <object-name>"
    local bucket=${1:?$usage}
    local object=${2:?$usage}

	local temp_file=

	export BASHMENOT_NO_SELF_UPDATE=1
	export BASHMENOT_AWS_ACCESS_KEY_ID="$AWS_ACCESS_KEY_ID"
    export BASHMENOT_AWS_SECRET_ACCESS_KEY="$AWS_SECRET_ACCESS_KEY"
    export BASHMENOT_S3_ENDPOINT="$AWS_S3_ENDPOINT_URL"

    source "$SCRIPT_DIR/../../../mietek/bashmenot/src.sh"

	temp_file="$(mktemp)"
	echo "Hi there, I was created at $(date -u -R)!" > "$temp_file"

	echo "==> Create temp file and upload it"
	set -x
	s3_upload "$temp_file" "$bucket" "$object" "public-read"
	set +x

	echo "==> Check file exists in S3"
	#set -x
	s3_check "$bucket" "$object"
	#set +x

	echo "==> Get file content"
	#set -x
	s3_download "$bucket" "$object" /dev/stdout
	#set +x
}

main "$@"
