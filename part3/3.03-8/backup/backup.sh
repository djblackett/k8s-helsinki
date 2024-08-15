#!/usr/bin/env sh

set -e

# Construct PostgreSQL connection string
POSTGRES_URL="postgres://${USER}:${PASSWORD}@${HOST}:${POSTGRES_PORT}/${DB_NAME}"

# Print the connection string for verification (optional)
echo "PostgreSQL connection string: $POSTGRES_URL"

if [ "$POSTGRES_URL" ]; then
    pg_dump -v "$POSTGRES_URL" >/usr/src/app/backup.sql

    echo "Attempting to send backup to google storage"
    gcloud storage cp /usr/src/app/backup.sql gs://db_backup1029
fi

echo "I'm running. Yay!"
