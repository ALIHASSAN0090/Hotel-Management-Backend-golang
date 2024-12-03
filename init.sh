set -e


until pg_isready -U "$POSTGRES_USER" -d "$POSTGRES_DB"; do
    echo "Waiting for PostgreSQL to be ready..."
    sleep 2
done

echo "Running migrations..."

# Run Go mod tidy
if ! go mod tidy; then
    echo "Error running go mod tidy"
    exit 1
fi


if ! go run db_migrator/main.go db_migrator/migrations.go; then
    echo "Error running migrations"
    exit 1
fi

echo "Migrations completed!"

if [ -z "$ADMIN_DEFAULT_EMAIL" ] || [ -z "$ADMIN_DEFAULT_PASSWORD" ]; then
    echo "Environment variables ADMIN_DEFAULT_EMAIL and ADMIN_DEFAULT_PASSWORD must be set"
    exit 1
fi


ADMIN_PASSWORD_HASH=$(echo -n "$ADMIN_DEFAULT_PASSWORD" | openssl dgst -sha256 | awk '{print $2}')

psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" <<EOF
INSERT INTO users (username, email, password_hash, first_name, last_name, role)
VALUES ('admin', '$ADMIN_DEFAULT_EMAIL', '$ADMIN_PASSWORD_HASH', 'Admin', 'User', 'admin')
ON CONFLICT (email) DO NOTHING;
EOF

echo "Default admin user created!"