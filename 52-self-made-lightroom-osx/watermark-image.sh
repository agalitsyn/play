#!/usr/bin/env bash

# bash watermark-image.sh ~/Pictures/catalog/winter/resampled/19 jpg 50
USAGE="$0 <logo-image-path> <watermark-image-path> <target> <file-extension> [scale % <none>] [copy-with-postfix <none>]"
LOGO_IMAGE_FILE_PATH=${1:?"$USAGE"}
WATERMARK_IMAGE_FILE_PATH=${2:?"$USAGE"}
TARGET=${3:?"$USAGE"}
EXTENSION=${4:?"$USAGE"}
SCALE=${5}
POSTFIX=${6}

WATERMARK_IMAGE_SETTINGS=${WATERMARK_IMAGE_SETTINGS:-"-resample 400 +level 70%"}
LOGO_IMAGE_SETTINGS=${LOGO_IMAGE_SETTINGS:-"-resample 200"}

PATTERN="*.$EXTENSION"

if [[ -n "$SCALE" ]]; then SCALE="-resize $SCALE%"; fi
if [[ -n "$POSTFIX" ]]; then POSTFIX="-$POSTFIX"; fi
time find "$TARGET" -type f -iname "$PATTERN" -print0 \
    | xargs -I{} -0 -P3 -L1 sh -c \
        "echo \"==> \$0\"; convert \"\$0\" \\
            \( $LOGO_IMAGE_FILE_PATH $LOGO_IMAGE_SETTINGS \) -gravity southeast -composite \\
            \( $WATERMARK_IMAGE_FILE_PATH $WATERMARK_IMAGE_SETTINGS \) -gravity center -compose screen -composite \\
            $SCALE \\
            \"\${0%.*}$POSTFIX.$EXTENSION\""
