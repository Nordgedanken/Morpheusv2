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

package db

import (
	"database/sql"
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix"
)

// DB defines a Interface to allow multiple DB Implementations
type DB interface {
	Init() error
	Open() *sql.DB

	SaveRoom(Room matrix.Room) error
	GetRooms() (roomsR []matrix.Room, err error)
	GetRoom(roomID string) (roomR matrix.Room, err error)

	SaveUser(user matrix.User) error
	GetUsers() ([]matrix.User, error)
	GetUser(userID string) (userR matrix.User, err error)
	GetCurrentUser() (userR matrix.User, err error)

	SaveMessage(message matrix.Message) error
	GetMessages(eventIDs []string) (messagesR []matrix.Message, err error)
	GetMessage(eventID string) (messageR matrix.Message, err error)
}
