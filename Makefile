.PHONY: postgresinit migrate-up migrate-down

postgresinit:
	docker run --name chart-postgres -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DATABASE=chartDB -d postgres

createdb:
	docker exec -it chart-postgres createdb --username=admin --owner=admin chartDB

migrate-up:
	migrate -path storage/migration -database "postgresql://admin:password@localhost:5432/chartDB?sslmode=disable" -verbose up

migrate-down:
	migrate -path storage/migration -database "postgresql://admin:password@localhost:5432/chartDB?sslmode=disable" -verbose down
