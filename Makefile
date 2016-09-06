picUrl=https://pbs.twimg.com/media/CjeQ9FwUUAAeJrS.jpg
db=./kusokora.db

install:
	go get "github.com/ChimeraCoder/anaconda"
	go get "github.com/mitchellh/cli"
	go get "github.com/mattn/go-sqlite3"
	go get "github.com/rubenv/sql-migrate/..."

run-add:
	go run main.go a $(picUrl)

run-crawl:
	go run main.go c

migrate/up:
	sql-migrate up

migrate/down:
	sql-migrate down

migrate/status:
	sql-migrate status

show:
	sqlite3 $(db) ".tables"
	sqlite3 $(db) "SELECT * FROM kusokoras"

clean:
	rm -f $(db)
