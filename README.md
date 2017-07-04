# DoCSoc Sponsor Portal

## Dependencies
 - Go 1.8+
 - Glide
 - Migrate
 - Docker
 - npm
 - yarn

## Build & Run
 - `make install` to install npm and go packages

 - `make client` to build the front-end assets for development and watch for changes (recommended)
 - `make build:dev` to build the front-end assets for development
 - `make build` to build the front-end assets for production

 - `make server` to start the go server

---

_I'll turn these into issues when they are more concrete_

### TODOs (packages):
 - auth/jwt: middleware for JWT tokens
 - cvs
 - profile
 - job board
 - logging, GitHub education has some free options for logging services

### TODOs (client side)
 - React router
 - decide on styles
 - bundle 3 javacsript files, index, student and sponsor

### TODOs (features):
 - user roles/validation
 - CV upload
 - CV viewing
 - CV commenting
 - Job board

