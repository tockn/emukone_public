.PHONY: server front test docker_up deps mysql migrate_up test/migrate_up

mysql:
	mysql -u root --protocol=tcp -D emukone -p

run: server front

server:
	which gin || go get github.com/codegangsta/gin
	gin -a 8081 -p 8080

front:
	npm run dev --prefix frontend

test:
	go test -v ./...

docker_up:
	docker-compose up -d

deps:
	which sql-migrate || go get github.com/rubenv/sql-migrate/...

migrate_up: deps
	sql-migrate up

test/migrate_up:
	sql-migrate up -env=test

swagger:
	which swag || go get -u github.com/swaggo/swag/cmd/swag
	swag init

lint:
	npm --prefix frontend run lintfix
