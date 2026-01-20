# Hashwood

Hashwood is a utility to calculate and store MD5 hashes and log any changes it detects. 
It can be used to launch a local web server which will display the history and changes 
to the hash of specified files.

## Why
When working with updates for various application it's often times difficult to find what files have changed.
I made Hashwood to help keep track of changes and highlight when something has been modified to quickly identify
potential faults or malicious changes. 

## Requirements for setup :
- Go
- Postgres
- Goose
  - Used to "goose up" and build the tables within the db and perform edits. 

--------------
## .env Example
Currently the .env must share a directory with the main hashwood file.

 ```
 DB_URL="postgres://example_user:@localhost:5432/hashwood?sslmode=disable"
 ```

--------------
## Commands
### Scan
- scan file ./example/file.png - Scans a single file and adding it into the database if it does not exist
- scan directory ./example/    - Scans selected directory showing the hashes and placing them into the db
  - Note: Sub directories of the target directory will be skipped.

### Autoscan
- autoscan start 60 ./example/    - Scans the target directory "./example/" every 60 seconds
- autoscan stop ./example/        - If an ./example/ scan is running this will stop it.

### Help
- help     - Prints a list of commands and their functionality.

### History
- history - List all current history entries
- history ./example/file.txt - returns the history for the specified file.

### Ignore
- ignore add ./example/file.txt      - Adds the file to the ignore list
- ignore remove ./example/file.txt   - Removes item from ignore list
- ignore list                        - List all items on the ignore list.

### Reset
- reset           - Resets the entire DB

### Server
- server start         - Launch the reporting webserver to http://localhost:8080/report
- server stop          - Stops the active server

### Quit
- quit or exit - exits the program


--------------
## Setup Process
- With postgres set up a new database
  - ``` 
    psql postgres
    CREATE DATABASE hashwood;
    ```
- Setup a .env in the main folder (see above for .env example)
- goose up from the schema directory
  - goose postgres "postgres://username:password@host:5432/hashwood" up
- return to main directory
- Launch using Go
  - ```
    go run .
    ```
Further build/standalone instructions will be available at a later time.
