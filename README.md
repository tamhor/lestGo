# lestGo

CRUD Rest API using Golang (Gorm & Fiber).

## Development
This project was developed to speed up me creating Rest API using golang.

## Package Using

### Gorm
- <a href="https://gorm.io/docs/">Gorm</a> is ORM Library for Golang.

### Fiber v2
- <a href="https://docs.gofiber.io/">Fiber</a> is designed to ease things up for fast development with zero memory allocation and performance in mind.

# Requirement

Download and install Golang <a href="https://go.dev/doc/install">here</a>.

# Project Setup

- Fork or clone this repo \
`git clone https://github.com/tamhor/lestGo.git`
- Go to directory \
`cd lestGo`
- Rename environtment to .env \
`mv environtment .env`
- Setup Database on .env \
-- Sqlite: just using DB_CONNECTION
`DB_CONNECTION="sqlite"` \
-- Postgresql \
`DB_CONNECTION="postgres"` \
`DB_HOST="DB_HOST"` \
`DB_PORT="DB_PORT"` \
`DB_NAME="DB_NAME"` \
`DB_USER="DB_USER"` \
`DB_PASSWORD="DB_PASSWORD"` 
- Get all requirement package \
`go get`
- Migrate Database \
`go run main.go -migrate`
- (Optional) If u using Air for hot reload, copy air to /usr/sbin/ path \
`sudo cp bin/air /usr/sbin/` \
--then run \
`air`
- If u dont using air, just run \
`go run main.go`
- All setup is complete, lestGo!