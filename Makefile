postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root socials

dropdb:
	docker exec -it postgres12 dropdb socials

migrateup:
	migrate -path cmd/migrate/migrations -database postgresql://root:postgres@localhost:5432/socials?sslmode=disable up

migratedown:
	migrate -path cmd/migrate/migrations -database postgresql://root:postgres@localhost:5432/socials?sslmode=disable down

test:
	go test -v -cover ./ ...


server: 
	go run cmd/api/**.go

.PHONY: postgres createdb dropdb migrateup migratedown test server