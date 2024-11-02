#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Использование: $0 coverage.out exclude.txt"
    exit 1
fi

COVERAGE_FILE=$1
EXCLUDE_FILE=$2

grep -v -F -f "$EXCLUDE_FILE" "$COVERAGE_FILE" > "${COVERAGE_FILE}.tmp" && mv "${COVERAGE_FILE}.tmp" "$COVERAGE_FILE"

echo "Файлы успешно отфильтрованы и $COVERAGE_FILE"
