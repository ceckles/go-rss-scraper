# Go RSS Scraper

# Installation

```bash
git clone https://github.com/ceckles/go-rss-scraper.git
```
#ENV
create a .env file within the root dir with the following content:
```bash
PORT=3000
DB_URL=<your db url>
```
# Usage
Nav to cloned directory and run the following commands:
```bash
#install
> go install
> go mode vendor # if you want to use vendor local packages
> go mod tidy  # clean up the go.mod file

#run
> go run main.go

#build and run
> go build && ./go-rss-scraper
```

# Goose
```bash
> cd sql/schema
> goose postgres 'DB_URL' down
> goose postgres 'DB_URL' up
```

# SQLC
```bash
> cd to project root
> sqlc generate
```