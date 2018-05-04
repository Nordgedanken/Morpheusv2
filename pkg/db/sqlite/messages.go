package sqlite

import "github.com/Nordgedanken/Morpheusv2/pkg/matrix"

func (s *SQLite) SaveMessage(message matrix.Message) error {
	return nil
}

func (s *SQLite) GetMessages(eventIDs []string) (users []matrix.Message, err error) {
	return nil, nil
}

func (s *SQLite) GetMessage(eventID string) (matrix.Message, error) {
	return nil, nil
}
