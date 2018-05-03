package sqlite

import (
	"bytes"
	"encoding/binary"
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
	i := 0
	var aliasesBuffer bytes.Buffer
	for _, v := range aliases {
		if i == 0 {
			aliasesBuffer.WriteString(v)
			i++
		} else {
			aliasesBuffer.WriteString("," + v)
		}
	}
	aliasesS := aliasesBuffer.String()
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

	// TODO This looks very fragile. REDO
	messages, err := Room.GetMessages()
	if err != nil {
		return err
	}
	var messagesBuffer bytes.Buffer
	for _, v := range messages {
		if i == 0 {
			binary.Write(&messagesBuffer, binary.BigEndian, v)
			i++
		} else {
			messagesBuffer.WriteString(",")
			binary.Write(&messagesBuffer, binary.BigEndian, v)
		}
	}
	messagesS := messagesBuffer.String()

	_, err = stmt.Exec(aliasesS, roomID, name, avatar, topic, messagesS)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (s *SQLite) GetRooms() (rooms map[string]matrix.Room, err error) {
	return nil, nil
}

func (s *SQLite) GetRoom(roomID string) (room matrix.Room, err error) {
	return nil, nil
}
