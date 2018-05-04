package sqlite

import "github.com/Nordgedanken/Morpheusv2/pkg/matrix"

// SaveMessage saves message Events to the DB
func (s *SQLite) SaveMessage(message matrix.Message) error {
	if s.db == nil {
		s.db = s.Open()
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO messages (id, author_id, message, timestamp, pure_event) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	id := message.GetEventID()
	authorID := message.GetAuthorMXID()
	messageS := message.GetMessage()
	timestampR := message.GetTimestamp()
	timestamp := timestampR.Format("2006-01-02 15:04:05")
	pureEvent := message.GetEvent()
	_, err = stmt.Exec(id, authorID, messageS, timestamp, pureEvent)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *SQLite) GetMessages(eventIDs []string) (users []matrix.Message, err error) {
	return nil, nil
}

func (s *SQLite) GetMessage(eventID string) (matrix.Message, error) {
	return nil, nil
}
