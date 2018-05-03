package sqlite

import (
	"encoding/json"
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix"
)

// SaveRoom saves a Room into the sqlite DB
func (s *SQLite) SaveRoom(Room matrix.Room) error {
	if s.db == nil {
		s.db = s.Open()
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO rooms (room_aliases, id, room_name, room_avatar, room_topic, room_messages) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	aliases := Room.GetRoomAliases()
	aliasesBytes, err := json.Marshal(aliases)
	if err != nil {
		return err
	}
	aliasesS := string(aliasesBytes)
	roomID := Room.GetRoomID()
	name, err := Room.GetName()
	if err != nil {
		return err
	}
	avatar, err := Room.GetAvatar()
	if err != nil {
		return err
	}
	topic, err := Room.GetTopic()
	if err != nil {
		return err
	}

	messages, err := Room.GetMessages()
	if err != nil {
		return err
	}
	messagesBytes, err := json.Marshal(messages)
	if err != nil {
		return err
	}
	messagesS := string(messagesBytes)

	_, err = stmt.Exec(aliasesS, roomID, name, avatar, topic, messagesS)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// GetRooms returns all Rooms from the Database
func (s *SQLite) GetRooms() (rooms []matrix.Room, err error) {
	if s.db == nil {
		s.db = s.Open()
	}

	rows, err := s.db.Query("SELECT id, room_aliases, room_name, room_avatar, room_topic, room_messages FROM rooms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var roomID string
		var roomAliases string
		var roomName string
		var roomAvatar string
		var roomTopic string
		var roomMessages string
		err = rows.Scan(&roomID, &roomAliases, &roomName, &roomAvatar, &roomTopic, &roomMessages)
		if err != nil {
			return nil, err
		}

		// TODO replace with implementation
		roomI := matrix.Room{}
		roomI.SetRoomID(roomID)
		var aliases []string
		err := json.Unmarshal([]byte(roomAliases), &aliases)
		if err != nil {
			return nil, err
		}
		roomI.SetRoomAliases(aliases)
		roomI.SetName(roomName)
		roomI.SetAvatar(roomAvatar)
		roomI.SetTopic(roomTopic)
		var messages []matrix.Message
		err = json.Unmarshal([]byte(roomMessages), &messages)
		if err != nil {
			return nil, err
		}
		roomI.SetMessages(messages)

		rooms = append(rooms, roomI)
	}

	// get any error encountered during iteration
	return rooms, rows.Err()
}

func (s *SQLite) GetRoom(roomID string) (room matrix.Room, err error) {
	return nil, nil
}
