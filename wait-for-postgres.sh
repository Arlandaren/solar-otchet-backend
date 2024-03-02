1#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
port="$2"
shift 2
cmd="$@"

until PGPASSWORD=$DB_PASSWORD psql -h "$host" -p "$port" -U "postgres" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd
