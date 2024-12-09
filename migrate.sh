#!/bin/bash
TIMESTAMP=$(date +"%Y%m%d%H%M%S")
MS=$1
NAME=$2
DIR="${MS}/internal/db/migrations"
EXT="sql"

if [ -z "$NAME" ]; then
  echo "Usage: $0 <migration_name>"
  exit 1
fi

mkdir -p "$DIR"

UP_FILE="${DIR}/${TIMESTAMP}_${NAME}.up.${EXT}"
DOWN_FILE="${DIR}/${TIMESTAMP}_${NAME}.down.${EXT}"

touch "$UP_FILE" "$DOWN_FILE"

echo "Migration created:"
echo "  $UP_FILE"
echo "  $DOWN_FILE"
