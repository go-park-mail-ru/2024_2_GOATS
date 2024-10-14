#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Использование: $0 coverage.out exclude.txt"
    exit 1
fi

COVERAGE_FILE=$1
EXCLUDE_FILE=$2


EXCLUDED_FILES=()
while IFS= read -r line; do
    EXCLUDED_FILES+=("$line")
done < "$EXCLUDE_FILE"

TEMP_FILE=$(mktemp)
while IFS= read -r line; do
    EXCLUDE=false
    for file in "${EXCLUDED_FILES[@]}"; do
        if [[ "$line" == *"$file"* ]]; then
            EXCLUDE=true
            break
        fi
    done
    if [ "$EXCLUDE" = false ]; then
        echo "$line" >> "$TEMP_FILE"
    fi
done < "$COVERAGE_FILE"

mv "$TEMP_FILE" "$COVERAGE_FILE"

echo "Файлы успешно отфильтрованы из $COVERAGE_FILE"
