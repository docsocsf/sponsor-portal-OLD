# DoCSoc Sponsor Portal

## Dependencies
 - Go 1.8+
 - [Glide](https://github.com/Masterminds/glide#install)
 - [Migrate](https://github.com/mattes/migrate#cli-usage)
 - Docker
 - npm
 - [yarn](https://yarnpkg.com/en/docs/install)

## Build & Run
 - `make install` to install npm and go packages

 - `make client` to build the front-end assets for development and watch for changes (recommended)
 - `make build:dev` to build the front-end assets for development
 - `make build` to build the front-end assets for production

 - `make server` to start the go server

---

_I'll turn these into issues when they are more concrete_

### TODOs (packages):
 - web: main entrypoint for the server, sets up auth, routes and index page
 - auth: OAuth login/logout package
   - auth/jwt: middleware for JWT tokens
 - model: all the main models needed for json and db
 - persistence: package to keep track of storing data to s3 and db
 - cvs
 - profile
 - job board
 - middleware? need a way to serve the student app or the sponsor app
 - logging, GitHub education has some free options for logging services

### TODOs (client side)
 - React router
 - basic index page
 - decide on styles
 - bundle 3 javacsript files, index, student and sponsor

### TODOs (features):
 - oAuth/signup
 - user roles/validation
 - CV upload
 - CV viewing
 - CV commenting
 - Job board

