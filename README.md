# go-clean-rest-api
Sample CRUD API using Golang and SQLite3. This project follows SOLID principles which contains of this structure
```
    /
    |- controllers
    |- db
    |- infrastructures
    |- interfaces
    |- models
    |- repositories
    |- services
    container.go
    main.go
    router.go
```

Dependencies:
- [Chi Router](github.com/go-chi/chi)
- [Goose (Database Migration)](bitbucket.org/liamstask/goose)
- [Ozzo (Attribute Validation)](github.com/go-ozzo/ozzo-validation/v4)
- [JWT (Security)](github.com/dgrijalva/jwt-go)
- [Crypto (Data Encryption)](golang.org/x/crypto)

## Installing

Clone source
```
    git clone https://github.com/irahardianto/service-pattern-go
```

Set dependencies
```
    go get github.com/go-chi/chi
    go get github.com/mattn/go-sqlite3
    go get golang.org/x/crypto
    go get github.com/dgrijalva/jwt-go
    go get github.com/go-ozzo/ozzo-validation/v
    go get bitbucket.org/liamstask/goose
```

Copy Env File
```
    cp env.sample .env
```

Run the test to make sure all run well (TODO)
```
    go test ./... -v
```

Build and run the server
```
    go build && ./go-clean-rest-api
```

## Endpoints
```
POST /users
POST /users/login
GET /users/:id
GET /users
PATCH /users/:id
DELETE /users/:id
```
