include .env

export PGPASSWORD
DB_CONN := postgres://${DB_USER}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}

.PHONY: start seed server client

install:
	glide install

setupdb:
	migrate -database ${DB_CONN} -path migrations up

SEED_CMD := psql -U ${DB_USER} -d ${DB_NAME} -a
USER_CMD := "${SEED_CMD} -c \"COPY users (name, email) FROM STDIN WITH CSV\""
SPONSOR_CMD := ${SEED_CMD} -f /seed/sponsors.sql
seed:
	docker exec -i $$(docker-compose ps -q db) sh -c ${USER_CMD} < seed/users.csv
	docker exec -i $$(docker-compose ps -q db) sh -c "${SPONSOR_CMD}"

setup: install setupdb seed

build:
	cd client; yarn build

build-dev:
	cd client; yarn build:dev

client:
	cd client; yarn && yarn start

server:
	go run web/*.go
