package matrix

import (
	"github.com/matrix-org/gomatrix"
	"time"
)

// User defines a Interface to allow multiple User type Implementations
type User interface {
	SetCli(cli *gomatrix.Client)
	SetMXID(id string)
	GetDisplayName(roomID string) (string, error)
	GetAvatar(roomID string) (string, error)

	//Implement  MarshalJSON() (b []byte, e error) as in http://gregtrowbridge.com/golang-json-serialization-with-interfaces/ to support json
}

// Room defines a Interface to allow multiple Room type Implementations
type Room interface {
	// Handled using global "own User"
	//SetCli(cli *gomatrix.Client)
	SetRoomID(id string)
	SetRoomAliases([]string)
	SetName(string)
	SetAvatar(string)
	SetTopic(string)
	SetMessages([]Message)

	GetRoomID() string
	GetRoomAliases() []string
	GetName() (string, error)
	GetAvatar() (string, error)
	GetTopic() (string, error)
	GetMessages() ([]Message, error)

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
	Show() error

	//Implement  MarshalJSON() (b []byte, e error) as in http://gregtrowbridge.com/golang-json-serialization-with-interfaces/ to support json
}
