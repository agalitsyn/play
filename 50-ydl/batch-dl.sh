#!/usr/bin/env bash

while IFS='' read -r line || [[ -n "$line" ]]; do
    echo "$line"
    youtube-dl "$line" &
done < "$1"

