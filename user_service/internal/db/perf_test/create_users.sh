#!/bin/bash

CSRF_URL="https://cassette-world.ru/api/csrf-token"
SIGNUP_URL="https://cassette-world.ru/api/auth/signup"

CSRF_RESPONSE=$(curl -i \
    -H "Content-Type: application/json" \
    "$CSRF_URL")

CSRF_SESSION=$(echo "$CSRF_RESPONSE" | grep -i "Set-Cookie" | sed -n 's/^Set-Cookie: \([^;]*\).*/\1/p' | tr -d '\r\n')
CSRF_TOKEN=$(echo "$CSRF_RESPONSE" | grep -i "X-Csrf-Token" | sed -n 's/^X-Csrf-Token: \([^;]*\).*/\1/p' | tr -d '\r\n')

WRK_SCRIPT="create_users.lua"
WRK_URL="https://cassette-world.ru/api/auth/signup"

wrk -t100 -c200  -d1200s --timeout 10s -s "$WRK_SCRIPT" --latency "$WRK_URL" -H "Cookie: $CSRF_SESSION" -H "X-Csrf-Token: $CSRF_TOKEN"
