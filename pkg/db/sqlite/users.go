package sqlite

import "github.com/Nordgedanken/Morpheusv2/pkg/matrix"

func (s *SQLite) SaveUser(user *matrix.User) error {
	return nil
}

func (s *SQLite) GetUsers() (map[string]*matrix.User, error) {
	return nil, nil
}

func (s *SQLite) GetCurrentUser() (*matrix.User, error) {
	return nil, nil
}
