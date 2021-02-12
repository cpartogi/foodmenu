PROJECT_NAME=foodmenu

install:
	cd .. && go get -u github.com/swaggo/swag/cmd/swag && go get -u github.com/cosmtrek/air && cd ${PROJECT_NAME} && swag init

local:
	air -c config/.air.toml

test:
	go test -v -cover ./...

compose-up:
	docker-compose up -d --build

compose-down:
	docker-compose down