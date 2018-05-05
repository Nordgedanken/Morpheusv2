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
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix/rooms"
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

	messages := Room.GetMessages()
	if err != nil {
		return err
	}
	var messageIDs []string
	for _, v := range messages {
		messageIDs = append(messageIDs, v.GetEventID())
	}

	_, err = stmt.Exec(aliasesS, roomID, name, string(avatar), topic, messageIDs)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// GetRooms returns all Rooms from the Database
func (s *SQLite) GetRooms() (roomsR []matrix.Room, err error) {
	if s.db == nil {
		s.db = s.Open()
	}

	rows, err := s.db.Query("SELECT id, room_aliases, room_name, room_avatar, room_topic, room_messages FROM rooms")
	if err != nil {
		return
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
			return
		}

		roomI := &rooms.Room{}
		roomI.SetRoomID(roomID)
		var aliases []string
		err = json.Unmarshal([]byte(roomAliases), &aliases)
		if err != nil {
			return
		}
		roomI.SetRoomAliases(aliases)
		roomI.SetName(roomName)
		roomI.SetAvatar([]byte(roomAvatar))
		roomI.SetTopic(roomTopic)
		var messageIDs []string
		err = json.Unmarshal([]byte(roomMessages), &messageIDs)
		if err != nil {
			return
		}
		// TODO Convert IDs to messages slice using another call or from the beginning on using a JOIN
		//roomI.SetMessages(messages)

		roomsR = append(roomsR, roomI)
	}

	// get any error encountered during iteration
	err = rows.Err()
	return
}

// GetRoom returns the Room where the id matches the roomID
func (s *SQLite) GetRoom(roomID string) (roomR matrix.Room, err error) {
	if s.db == nil {
		s.db = s.Open()
	}

	stmt, err := s.db.Prepare(`SELECT room_aliases, room_name, room_avatar, room_topic, room_messages FROM rooms WHERE id=$1`)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(roomID)

	var roomAliases string
	var roomName string
	var roomAvatar string
	var roomTopic string
	var roomMessages string
	err = row.Scan(&roomAliases, &roomName, &roomAvatar, &roomTopic, &roomMessages)
	if err != nil {
		return
	}

	roomI := &rooms.Room{}
	roomI.SetRoomID(roomID)
	var aliases []string
	err = json.Unmarshal([]byte(roomAliases), &aliases)
	if err != nil {
		return
	}
	roomI.SetRoomAliases(aliases)
	roomI.SetName(roomName)
	roomI.SetAvatar([]byte(roomAvatar))
	roomI.SetTopic(roomTopic)
	var messageIDs []string
	err = json.Unmarshal([]byte(roomMessages), &messageIDs)
	if err != nil {
		return
	}
	// TODO Convert IDs to messages slice using another call or from the beginning on using a JOIN
	//roomI.SetMessages(messages)

	roomR = roomI
	return
}
