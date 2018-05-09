// Copyright © 2018 MTRNord <info@nordgedanken.de>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	createTables := `CREATE TABLE IF NOT EXISTS users (id varchar not null primary key, display_name text, avatar text, access_token text, own integer);
					CREATE TABLE IF NOT EXISTS messages (id varchar not null primary key, author_id varchar, message text, timestamp datetime, pure_event text);
					CREATE TABLE IF NOT EXISTS rooms (id varchar not null primary key, room_aliases text, room_name text, room_avatar text, room_topic text, room_messages text);
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

func (s *SQLite) RemoveAll() error {
	if s.db == nil {
		s.db = s.Open()
	}

	_, err := s.db.Exec("DELETE FROM users;")
	if err != nil {
		return err
	}

	_, err = s.db.Exec("DELETE FROM messages;")
	if err != nil {
		return err
	}

	_, err = s.db.Exec("DELETE FROM rooms;")
	if err != nil {
		return err
	}
	return nil
}
