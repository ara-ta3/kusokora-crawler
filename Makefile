install:
	go get "github.com/ChimeraCoder/anaconda"
	go get "github.com/mitchellh/cli"
	go get "github.com/mattn/go-sqlite3"
	go get "github.com/rubenv/sql-migrate/..."

run:
	go run main.go "aaa" "bbb"

migrate/up:
	sql-migrate up

migrate/status:
	sql-migrate status
