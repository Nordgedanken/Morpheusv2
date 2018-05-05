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
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix"
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix/users"
	"github.com/matrix-org/gomatrix"
	"strings"
)

// SaveUser saves a matrix User to the User table
func (s *SQLite) SaveUser(user matrix.User) error {
	if s.db == nil {
		s.db = s.Open()
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO users (id, display_name, avatar) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	displayName, err := user.GetDisplayName("")
	if err != nil {
		return err
	}
	avatar, err := user.GetAvatar("")
	if err != nil {
		return err
	}
	mxid := user.GetMXID()
	_, err = stmt.Exec(mxid, displayName, avatar)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// GetCurrentUser returns the current user including a ready gomatrix client
func (s *SQLite) GetCurrentUser() (userR matrix.User, err error) {
	if s.db == nil {
		s.db = s.Open()
	}

	// TODO Get from a file which user is current and accessToken
	var mxid string
	var accessToken string

	stmt, err := s.db.Prepare(`SELECT display_name, avatar FROM users WHERE id=$1`)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(mxid)

	var displayName string
	var avatar string
	err = row.Scan(&displayName, &avatar)
	if err != nil {
		return
	}

	userI := &users.User{}
	userI.SetMXID(mxid)
	userI.SetDisplayName("", displayName)
	userI.SetAvatar("", avatar)

	splitUser := strings.Split(mxid, ":")
	domain := strings.TrimPrefix(mxid, splitUser[0]+":")
	client, err := gomatrix.NewClient(domain, mxid, accessToken)
	if err != nil {
		return
	}
	userI.SetCli(client)

	userR = userI
	return
}

// GetUser returns a user matching the mxid of the argument
func (s *SQLite) GetUser(userID string) (userR matrix.User, err error) {
	if s.db == nil {
		s.db = s.Open()
	}

	stmt, err := s.db.Prepare(`SELECT display_name, avatar FROM users WHERE id=$1`)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(userID)

	var displayName string
	var avatar string
	err = row.Scan(&displayName, &avatar)
	if err != nil {
		return
	}

	userI := &users.User{}
	userI.SetMXID(userID)
	userI.SetDisplayName("", displayName)
	userI.SetAvatar("", avatar)

	userR = userI
	return
}
