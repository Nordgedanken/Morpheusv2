package matrix

import (
	"github.com/matrix-org/gomatrix"
	"time"
)

type User interface {
	SetCli(cli *gomatrix.Client)
	SetMXID(id string)
	GetDisplayName(roomID string) (string, error)
	GetAvatar(roomID string) (string, error)
}

type Room interface {
	// Handled using global "own User"
	//SetCli(cli *gomatrix.Client)
	SetRoomID(id string)
	SetRoomAliases(aliases map[int64]string)
	GetName() (string, error)
	GetAvatar() (string, error)
	GetTopic() (string, error)
	GetMessages() (map[string]Message, error)
}

type Message interface {
	// Handled using global "own User"
	//SetCli(cli *gomatrix.Client)
	SetEventID(id string)
	SetEvent(event *gomatrix.Event)
	SetAuthorMXID(mxid string)
	SetMessage(message string)
	SetTimestamp(ts *time.Time)
	Show() error
}
