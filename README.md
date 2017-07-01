# DoCSoc Sponsor Portal

## Dependencies
 - Go 1.8+
 - Glide
 - Migrate
 - Docker
 - npm
 - yarn

## Build & Run
 - TBC

---

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

