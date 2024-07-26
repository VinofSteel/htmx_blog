#!/bin/sh

# Wait for the database to be ready
echo "Waiting for database to be ready..."
MAX_TRIES=30
TRIES=0

until pg_isready -h $PGHOST -p $PGPORT -U $PGUSER || [ $TRIES -eq $MAX_TRIES ]; do
  echo "Waiting for database container to be up... $TRIES/$MAX_TRIES"
  TRIES=$((TRIES+1))
  sleep 2
done

if [ $TRIES -eq $MAX_TRIES ]; then
  echo "Database container is down after $MAX_TRIES attempts, exiting."
  exit 1
fi

echo "Database container is up."

# Check if the database exists
DB_EXISTS=$(psql -h $PGHOST -p $PGPORT -U $PGUSER -tAc "SELECT 1 FROM pg_database WHERE datname='$PGDATABASE'")

if [ "$DB_EXISTS" != "1" ]; then
  echo "Database $PGDATABASE does not exist. Creating it..."
  createdb -h $PGHOST -p $PGPORT -U $PGUSER $PGDATABASE
else
  echo "Database $PGDATABASE already exists. Skipping creation."
fi

# Run migrations
echo "Running migrations to $PGDATABASE db..."
goose -dir sql/schema postgres "$PG_CONN_STRING" up

# Start the application
if [ "$ENV" = "production" ]; then
  make build && ./htmx_blog
else
  make run
fi
