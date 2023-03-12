makesql:
	migrate create -ext sql -dir db/migrations -seq init_tables

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root movies

dropdb:
	docker exec -it postgres12 dropdb --username=root movies

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/movies?sslmode=disable" -verbose up

migrateup_ver:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/movies?sslmode=disable" -verbose up $(ver)

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/movies?sslmode=disable" -verbose down

migratedown_ver:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/movies?sslmode=disable" -verbose down $(ver)

.PHONY: makesql
