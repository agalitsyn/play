#!/usr/bin/env bash

set -xe

USAGE="Usage: $0 <api http://localhost:8182/v2> <entry http://localhost:8181>"
API=${1:?$USAGE}
ENTRY=${2:?$USAGE}

SERVERS_GROUP_ONE="http://localhost:5000 http://localhost:5001"
SERVERS_GROUP_TW0="http://localhost:5002 http://localhost:5003"


function do_curl() {
	curl --output /dev/null --silent --fail "$@"
}

function api_set() {
	do_curl -X POST -H "Content-Type: application/json" "$@"
}

function start_servers() {
	python -mSimpleHTTPServer 5000 &
	FIRST_SERVER_PID=$!
	python -mSimpleHTTPServer 5001 &
	SECOND_SERVER_PID=$!
	python -mSimpleHTTPServer 5002 &
	THIRD_SERVER_PID=$!
	python -mSimpleHTTPServer 5003 &
	FOURTH_SERVER_PID=$!
}

function upsert_servers() {
	local usage="usage: upsert_servers <backend> <servers>"
	local backend=${1:?$usage}
	local servers=(${2:?$usage})

	for server in "${servers[@]}"; do
		api_set $API/backends/$backend/servers -d \
		"{
			\"Server\": {
				\"Id\": \"$(LC_ALL=C tr -dc 'a-zA-Z' < /dev/urandom | head -c 12)\",
				\"URL\": \"$server\"
			}
		}"
	done
}

function configure_first_group() {
	api_set $API/backends -d@- <<-EOF
	{
		"Backend": {
			"Id": "b1",
			"Type": "http",
			"Settings": {
				"Timeouts": {
					"Read": "1.5s",
					"Dial": "50ms"
				},
				"KeepAlive": {
					"MaxIdleConnsPerHost": 100,
					"Period": "65s"
				}
			}
		}
	}
	EOF

	api_set $API/backends -d@- <<-EOF
	{
		"Backend": {
			"Id": "b1",
			"Type": "http",
			"Settings": {
				"Timeouts": {
					"Read": "1.5s",
					"Dial": "50ms"
				},
				"KeepAlive": {
					"MaxIdleConnsPerHost": 100,
					"Period": "65s"
				}
			}
		}
	}
	EOF
	upsert_servers "b1" "$SERVERS_GROUP_ONE"

	api_set $API/frontends -d@- <<-EOF
	{
		"Frontend": {
			"Id": "f1",
			"Type": "http",
			"BackendId": "b1",
			"Route": "Path(\"/v1\")",
			"Settings": {
				"FailoverPredicate": "(IsNetworkError() || ResponseCode() == 500 || ResponseCode() == 502 || ResponseCode() == 503 || ResponseCode() == 504) && Attempts() <= 5"
			}
		}
	}
	EOF

	api_set $API/frontends/f1/middlewares -d@- <<-EOF
	{
		"Middleware": {
			"Id": "r1",
			"Priority": 1,
			"Type": "rewrite",
			"Middleware": {
				"Regexp": "/v1(.*)",
				"Replacement": "$1",
				"RewriteBody": false,
				"Redirect": false
			}
		}
	}
	EOF
}

function configure_second_group() {
	api_set $API/backends -d@- <<-EOF
	{
		"Backend": {
			"Id": "b2",
			"Type": "http",
			"Settings": {
				"Timeouts": {
					"Read": "1.5s",
					"Dial": "150ms"
				},
				"KeepAlive": {
					"MaxIdleConnsPerHost": 100,
					"Period": "65s"
				}
			}
		}
	}
	EOF
	upsert_servers "b2" "$SERVERS_GROUP_TW0"

	api_set $API/frontends -d@- <<-EOF
	{
		"Frontend": {
			"Id": "f2",
			"Type": "http",
			"BackendId": "b2",
			"Route": "Path(\"/v2\")",
			"Settings": {
				"FailoverPredicate": "(IsNetworkError() || ResponseCode() == 500 || ResponseCode() == 502 || ResponseCode() == 503 || ResponseCode() == 504) && Attempts() <= 5"
			}
		}
	}
	EOF

	api_set $API/frontends/f2/middlewares -d@- <<-EOF
	{
		"Middleware": {
			"Id": "r2",
			"Priority": 1,
			"Type": "rewrite",
			"Middleware": {
				"Regexp": "/v2(.*)",
				"Replacement": "$1",
				"RewriteBody": false,
				"Redirect": false
			}
		}
	}
	EOF
}

function test_settings() {
	echo 'Send requests to first group'
	for i in {1..5}; do
		echo "=> $i"
		do_curl $ENTRY/v1
	done

	echo 'Send requests to second group'
	for i in {1..5}; do
		echo "=> $i"
		do_curl $ENTRY/v2
	done
}

function main() {
	start_servers
	configure_first_group
	configure_second_group
	test_settings
	kill -SIGTERM $FIRST_SERVER_PID $SECOND_SERVER_PID $THIRD_SERVER_PID $FOURTH_SERVER_PID
}


main