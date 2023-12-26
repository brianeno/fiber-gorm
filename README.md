# project-go-fiber

## Get t Started -

- Create a database projects in your Postgres local instance
- Rename `.env.exmaple` to `.env`.
- Add in the database user and password for your postgres instance containing the database projects.
- In the root folder run `go run main.go`.
- Get the API docs at http://localhost:3000/swagger/index.html


## Make

go install github.com/cespare/reflex@latest
make watch

## Installs

go install github.com/swaggo/swag/cmd/swag@latest

go get -u github.com/gofiber/fiber/v2
go get github.com/google/uuid
go get github.com/joho/godotenv
go get gorm.io/gorm
go get gorm.io/driver/postgres

## Install Swagger docs
swag init --parseDependency --parseInternal


