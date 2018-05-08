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

package sync

import (
	"database/sql"
	"encoding/json"
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix/messages"
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix/rooms"
	"github.com/Nordgedanken/Morpheusv2/pkg/util"
	"github.com/matrix-org/gomatrix"
	"log"
	"time"
)

func parseEventTimestamp(unixTime int64) *time.Time {
	time := time.Unix(0, unixTime*int64(time.Millisecond))
	return &time
}

func NewSync() error {
	syncer := util.User.GetCli().Syncer.(*gomatrix.DefaultSyncer)
	filter := json.RawMessage(`{"room":{"state":{"types":["m.room.*"]},"timeline":{"limit":50,"types":["m.room.message"]}}}`)
	resp, err := util.User.GetCli().CreateFilter(filter)
	if err != nil {
		return err
	}
	filterID := resp.FilterID
	util.User.GetCli().Store.SaveFilterID(util.User.GetCli().UserID, filterID)

	syncer.OnEventType("m.room.message", func(ev *gomatrix.Event) {
		log.Println("New Message")
		msg := &messages.Message{}
		msg.SetEvent(ev)
		msg.SetTimestamp(parseEventTimestamp(ev.Timestamp))
		msg.SetAuthorMXID(ev.Sender)
		msg.SetEventID(ev.ID)
		room, err := util.DB.GetRoom(ev.RoomID)
		if err == sql.ErrNoRows {
			room = &rooms.Room{}
			room.SetRoomID(ev.RoomID)
		} else if err != nil && err != sql.ErrNoRows {
			log.Panicln(err)
		}
		messages := room.GetMessages()
		messages = append(messages, msg)
		room.SetMessages(messages)

		go func() {
			err := util.DB.UpdateRoom(room)
			if err != nil {
				log.Panicln(err)
			}
		}()
		go func() {
			err := util.DB.SaveMessage(msg)
			if err != nil {
				log.Panicln(err)
			}
		}()
	})

	go func() {
		log.Println("Start Sync...")
		for {
			if err := util.User.GetCli().Sync(); err != nil {
				log.Panicln("Sync err:", err)
			}
		}
	}()
	return nil
}

func Stop() {
	log.Println("Stop Sync...")
	util.User.GetCli().StopSync()
}
