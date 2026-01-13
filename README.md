# Hashwood

WIP

## Requirements :
- Go
- Postgres
- Goose
  - Used to "goose up" and build the tables within the db and perform edits. 
- Go's UUID system
  - More info to come. It may be dropped.
- SQLC for quick Go SQL functions.
WIP

--------------
## How to launch
WIP

--------------
## .env Example
Currently the .env must share a directory with the main hashwood file.

 ```
 DB_URL="postgres://example_user:@localhost:5432/hashwood?sslmode=disable"
 ```

--------------
## Commands
- scan file - Not Implemented
- scan directory ./example/    - Scans selected directory showing the hashes and placing them into the db

--------------
## Setup Process
- goose up from the schema directory
  - goose postgres "postgres://username:password@host:5432/hashwood" up
