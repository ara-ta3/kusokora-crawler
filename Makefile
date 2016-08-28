picUrl=https://pbs.twimg.com/media/CjeQ9FwUUAAeJrS.jpg
twitterUrl=https://twitter.com/gomaaburamax/status/736216444275294208
db=./kusokora.db

install:
	go get "github.com/ChimeraCoder/anaconda"
	go get "github.com/mitchellh/cli"
	go get "github.com/mattn/go-sqlite3"
	go get "github.com/rubenv/sql-migrate/..."

run:
	go run main.go $(picUrl) $(twitterUrl)

migrate/up:
	sql-migrate up

migrate/down:
	sql-migrate down

migrate/status:
	sql-migrate status

show:
	sqlite3 $(db) "SELECT * FROM kusokoras"

clean:
	rm -f $(db)
