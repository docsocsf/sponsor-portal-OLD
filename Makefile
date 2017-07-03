export HOST = :8080
export STATIC_FILES = client/dist

.PHONY: start

install:
	cd client; yarn install
	glide install

build:
	cd client; yarn build

build-dev:
	cd client; yarn build:dev

client: install
	cd client; yarn start

server: install
	go run web/main.go
