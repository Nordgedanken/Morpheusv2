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
	"encoding/json"
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix"
	"github.com/matrix-org/gomatrix"
	"time"
)

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

// GetMessages returns all Messages from the Database
func (s *SQLite) GetMessages(eventIDs []string) (messages []matrix.Message, err error) {
	if s.db == nil {
		s.db = s.Open()
	}

	rows, err := s.db.Query("SELECT id, author_id, message, timestamp, pure_event FROM messages")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var eventID string
		var authorID string
		var messageS string
		var timestamp time.Time
		var pureEvent string
		err = rows.Scan(&eventID, &authorID, &messageS, &timestamp, &pureEvent)
		if err != nil {
			return
		}

		// TODO replace with implementation
		messageI := matrix.Message{}
		messageI.SetEventID(eventID)
		messageI.SetAuthorMXID(authorID)
		messageI.SetMessage(messageS)
		messageI.SetTimestamp(&timestamp)
		var gomatrixEvent gomatrix.Event
		err = json.Unmarshal([]byte(pureEvent), &gomatrixEvent)
		if err != nil {
			return
		}
		messageI.SetEvent(&gomatrixEvent)

		messages = append(messages, messageI)
	}

	// get any error encountered during iteration
	err = rows.Err()
	return
}

func (s *SQLite) GetMessage(eventID string) (matrix.Message, error) {
	return nil, nil
}
