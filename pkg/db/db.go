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
	GetRooms() (rooms []matrix.Room, err error)
	GetRoom(roomID string) (room matrix.Room, err error)

	SaveUser(user matrix.User) error
	GetUsers() ([]matrix.User, error)
	GetUser(userID string) (userR matrix.User, err error)
	GetCurrentUser() (userR matrix.User, err error)

	SaveMessage(message matrix.Message) error
	GetMessages(eventIDs []string) (users []matrix.Message, err error)
	GetMessage(eventID string) (matrix.Message, error)
}
