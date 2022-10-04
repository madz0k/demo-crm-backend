# demo-crm-backend

Apply migrations:
```shell
export POSTGRESQL_URL="postgres://crm:crm@localhost:5432/crm?sslmode=disable"
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```