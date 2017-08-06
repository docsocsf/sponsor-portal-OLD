include .env

export PGPASSWORD
DB_CONN := postgres://${DB_USER}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}

.PHONY: start seed server client

install:
	glide install

setupdb:
	migrate -database ${DB_CONN} -path migrations up

SEED_CMD := "psql -U ${DB_USER} -d ${DB_NAME} -c \"COPY users (name, email) FROM STDIN WITH CSV\""
seed:
	docker exec -i $$(docker-compose ps -q db) sh -c ${SEED_CMD} < seed/users.csv

setup: install setupdb seed

build:
	cd client; yarn build

build-dev:
	cd client; yarn build:dev

client:
	cd client; yarn && yarn start

server:
	go run web/*.go
