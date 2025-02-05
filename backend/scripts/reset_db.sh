#!/bin/bash

# Exit on error
set -e

echo "🗑️  Dropping existing database..."
dropdb --if-exists squawker_db

echo "🛠️  Creating fresh database..."
createdb squawker_db

echo "📊 Applying schema..."
psql squawker_db < "$(dirname "$0")/../internal/database/schema.sql"

echo "✅ Database reset complete!"
