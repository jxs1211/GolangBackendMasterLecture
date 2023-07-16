#!/bin/sh

set -e
echo "run db migration"
source /app/app.env
/app/migrate -path db/migration/ -database "$DATABASESOURCE" up
exec "$@"

