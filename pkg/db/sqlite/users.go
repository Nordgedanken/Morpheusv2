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

func (s *SQLite) GetUsers() ([]matrix.User, error) {
	return nil, nil
}

func (s *SQLite) GetCurrentUser() (matrix.User, error) {
	return nil, nil
}
