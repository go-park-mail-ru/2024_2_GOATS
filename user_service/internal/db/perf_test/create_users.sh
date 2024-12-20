#!/bin/bash

CSRF_URL="https://cassette-world.ru/api/csrf-token"
SIGNUP_URL="https://cassette-world.ru/api/auth/signup"

CSRF_RESPONSE=$(curl -i \
    -H "Content-Type: application/json" \
    "$CSRF_URL")

echo "$CSRF_RESPONSE"
CSRF_SESSION=$(echo "$CSRF_RESPONSE" | grep -i "set-cookie" | sed -n 's/^set-cookie: \([^;]*\).*/\1/p' | tr -d '\r\n')
CSRF_TOKEN=$(echo "$CSRF_RESPONSE" | grep -i "x-csrf-token" | sed -n 's/^x-csrf-token: \([^;]*\).*/\1/p' | tr -d '\r\n')

echo "$CSRF_SESSION"
echo "$CSRF_TOKEN"

WRK_SCRIPT="create_users.lua"
WRK_URL="https://cassette-world.ru/api/auth/signup"

wrk -t100 -c150 -d1500s -s "$WRK_SCRIPT" --latency "$WRK_URL" -H "Cookie: $CSRF_SESSION" -H "X-Csrf-Token: $CSRF_TOKEN"
