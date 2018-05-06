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

package matrix

import (
	"github.com/matrix-org/gomatrix"
	"time"
)

// User defines a Interface to allow multiple User type Implementations
type User interface {
	SetCli(cli *gomatrix.Client)
	SetMXID(id string)
	SetDisplayName(roomID string, name string)
	SetAvatar(roomID string, avatar []byte)

	GetMXID() string
	GetDisplayName(roomID string) (string, error)
	GetAvatar(roomID string) ([]byte, error)
	GetCli() (cli *gomatrix.Client)

	GetAccessToken() string

	//Implement  MarshalJSON() (b []byte, e error) as in http://gregtrowbridge.com/golang-json-serialization-with-interfaces/ to support json
}

// Room defines a Interface to allow multiple Room type Implementations
type Room interface {
	// Handled using global "own User"
	//SetCli(cli *gomatrix.Client)
	SetRoomID(id string)
	SetRoomAliases(aliases []string)
	SetName(name string)
	SetAvatar(avatar []byte)
	SetTopic(topic string)
	SetMessages(messages []Message)

	GetRoomID() string
	GetRoomAliases() []string
	GetName() (string, error)
	GetAvatar() ([]byte, error)
	GetTopic() (string, error)
	GetMessages() []Message

	//Implement  MarshalJSON() (b []byte, e error) as in http://gregtrowbridge.com/golang-json-serialization-with-interfaces/ to support json
}

// Message defines a Interface to allow multiple Message type Implementations
type Message interface {
	// Handled using global "own User"
	//SetCli(cli *gomatrix.Client)
	SetEventID(id string)
	SetEvent(event *gomatrix.Event)
	SetAuthorMXID(mxid string)
	SetMessage(message string)
	SetTimestamp(ts *time.Time)

	GetEventID() (id string)
	GetEvent() (event *gomatrix.Event)
	GetAuthorMXID() (mxid string)
	GetMessage() (message string)
	GetTimestamp() (ts *time.Time)

	Show() error

	//Implement  MarshalJSON() (b []byte, e error) as in http://gregtrowbridge.com/golang-json-serialization-with-interfaces/ to support json
}
