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

### Goals

To provide a place for DoCSoc sponsors and DoC students to connect

Sponsors can:
* Search through student profiles
* View CVs
* Reach out to students (if the student has given permission)
* Manage their profile page, i.e. logo, description, contact details, application details, what they are looking for, roles they are advertising

Students can:
* Insert all details which companies require so that they don’t have to fill any other forms
* Exactly which fields is up for discussion
* Companies usually hire for multiple positions/countries. How will we handle that?
* Select which companies they wish to intern/placement/grad at
* Give permission to share personal information (ie. CV, email, mobile, LinkedIn) with sponsors (very important)
* Give permission to allow other students to see their profile
* Filter sponsors based on what the sponsors are looking for (eg a fresher should see that Company X doesn’t take freshers)
* Track their applications
  * My idea for this is some kind of funky Trello-esque view which gives you a quick glance summary of how things are looking
* Select where they have chosen as a destination
* Share reviews of the application process or questions they have been asked (careful about NDA)  - all internal (sponsors should not be able to see)


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

