package sqlite

import (
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix"
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

// GetUsers returns all Users of the Database
func (s *SQLite) GetUsers() (users []matrix.User, err error) {
	if s.db == nil {
		s.db = s.Open()
	}

	rows, err := s.db.Query("SELECT id, display_name, avatar FROM users")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var mxid string
		var displayName string
		var avatar string
		err = rows.Scan(&mxid, &displayName, &avatar)
		if err != nil {
			return
		}

		// TODO replace with implementation
		userI := matrix.User{}
		userI.SetMXID(mxid)
		userI.SetDisplayName("", displayName)
		userI.SetAvatar("", avatar)

		users = append(users, userI)
	}

	// get any error encountered during iteration
	err = rows.Err()
	return
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

	// TODO replace with implementation
	userI := matrix.User{}
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

	// TODO replace with implementation
	userI := matrix.User{}
	userI.SetMXID(userID)
	userI.SetDisplayName("", displayName)
	userI.SetAvatar("", avatar)

	userR = userI
	return
}
