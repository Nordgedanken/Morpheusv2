package sqlite

import (
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix"
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

func (s *SQLite) GetCurrentUser() (matrix.User, error) {
	return nil, nil
}
