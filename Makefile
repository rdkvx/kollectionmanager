MIGRATIONS_PATH=deployment/migrations

migrate-creation:
	migrate create -ext sql -dir ${MIGRATIONS_PATH} -seq 02_all_tables

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path ${MIGRATIONS_PATH} up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path ${MIGRATIONS_PATH} down