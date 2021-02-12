PROJECT_NAME=foodmenu

# DATABASE
DB_USER=root
DB_PASSWORD=root
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=foodmenu

install:
	cd .. && go get -u github.com/swaggo/swag/cmd/swag && go get -u github.com/cosmtrek/air && cd ${PROJECT_NAME} && swag init

local:
	air -c config/.air.toml

test:
	go test -v -cover ./...

migrate-up:
	migrate -source file:./scripts/migrations/ -database mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME} up

migrate-down:
	migrate -source file:./scripts/migrations/ -database mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME} down


compose-up:
	docker-compose up -d --build

compose-down:
	docker-compose down