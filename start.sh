#!/bin/sh

set -e
echo "run db migration"
/app/migrate -path db/migration/ -database "$DATABASESOURCE" up
exec "$@"

