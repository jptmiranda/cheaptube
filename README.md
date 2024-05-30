# CheapTube

## Database

### Create migrations

```bash
podman run --rm --net=host \
  -v $(pwd)/db/migrations:/migrations \
  -v $(pwd)/sqlc:/sqlc \
  docker.io/arigaio/atlas migrate diff ${MIGRATION_NAME} \
  --to "file:///sqlc/schema.sql" \
  --dev-url "postgresql://admin:admin@:5432/cheaptube?sslmode=disable&search_path=public"
```

### Run migrations

```bash
podman run --rm --net=host \
  -v $(pwd)/db/migrations:/migrations \
  docker.io/arigaio/atlas migrate apply \
  --url "postgresql://admin:admin@:5432/cheaptube?sslmode=disable"
```