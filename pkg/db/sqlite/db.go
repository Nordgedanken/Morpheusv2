package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Go-sqlite3 side effect import needed to use SQLite3 databases
	"github.com/shibukawa/configdir"
	"log"
	"path/filepath"
)

// Init prepares the DB by opening it and creating the required tables if needed
// TODO: Make part of the interface
func (s *SQLite) Init() (err error) {
	log.Println("Start setting up DB")
	var openErr error

	// Open the data.db file. It will be created if it doesn't exist.
	configDirs := configdir.New("Nordgedanken", "Morpheusv2")
	filePath := filepath.ToSlash(configDirs.QueryFolders(configdir.Global)[0].Path)

	log.Println("DBFilePath: ", filePath+"/data.db")
	s.db, openErr = sql.Open("sqlite3", filePath+"/data.db")
	if openErr != nil {
		err = openErr
		return
	}

	log.Println("Creating DB Tables if needed")
	createTables := `CREATE TABLE IF NOT EXISTS users (id integer not null primary key, display_name text, avatar text);
					CREATE TABLE IF NOT EXISTS messages (id integer not null primary key, author text, message text, timestamp text, pure_event text);
					CREATE TABLE IF NOT EXISTS rooms (id integer not null primary key, room_aliases text, room_id text, room_name text, room_avatar text, room_topic text, room_messages text);
					`
	_, execErr := s.db.Exec(createTables)
	if execErr != nil {
		err = fmt.Errorf("DB EXEC ERR: %s", execErr)
		return
	}
	log.Println("Finished setting DB Setup")
	return
}

// Open returns the in Init() created db variable
func (s *SQLite) Open() *sql.DB {
	if s.db != nil {
		return s.db
	}
	return nil
}
