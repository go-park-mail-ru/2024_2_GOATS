#!/bin/bash

CSRF_URL="https://cassette-world.ru/api/csrf-token"
LOGIN_URL="https://cassette-world.ru/api/auth/login"
SIGNUP_URL="https://cassette-world.ru/api/auth/signup"

CSRF_RESPONSE=$(curl -i \
    -H "Content-Type: application/json" \
    "$CSRF_URL")

CSRF_SESSION=$(echo "$CSRF_RESPONSE" | grep -i "Set-Cookie" | sed -n 's/^Set-Cookie: \([^;]*\).*/\1/p' | tr -d '\r\n')
CSRF_TOKEN=$(echo "$CSRF_RESPONSE" | grep -i "X-Csrf-Token" | sed -n 's/^X-Csrf-Token: \([^;]*\).*/\1/p' | tr -d '\r\n')

EMAIL="test_$(date +%s)@example.com"
PASSWORD="password$(date +%s)"
USERNAME="username$(date +%s)"

SIGNUP_RESPONSE=$(curl -i -X POST \
    -H "Content-Type: application/json" \
    -H "X-Csrf-Token: $CSRF_TOKEN" \
    -H "Cookie: $CSRF_SESSION" \
    -d '{"email": "'$EMAIL'", "password": "'$PASSWORD'", "passwordConfirmation": "'$PASSWORD'", "username": "'$USERNAME'"}' \
    "$SIGNUP_URL")

LOGIN_RESPONSE=$(curl -i -X POST \
    -H "Content-Type: application/json" \
    -H "X-Csrf-Token: $CSRF_TOKEN" \
    -H "Cookie: $CSRF_SESSION" \
    -d '{"email": "'$EMAIL'", "password": "'$PASSWORD'"}' \
    "$LOGIN_URL")

COOKIE=$(echo "$LOGIN_RESPONSE" | grep -i "Set-Cookie" | sed -n 's/^Set-Cookie: \([^;]*\).*/\1/p')

if [ -z "$COOKIE" ]; then
    echo "Не удалось получить куку"
    exit 1
fi

echo "Кука успешно получена: $COOKIE"

WRK_SCRIPT="session_data.lua"
WRK_URL="https://cassette-world.ru/api/auth/session"

wrk -t4 -c100 -d30s -s "$WRK_SCRIPT" --latency "$WRK_URL" -H "Cookie: $COOKIE"
