package db

import "github.com/Nordgedanken/Morpheusv2/pkg/matrix"

// DB defines a Interface to allow multiple DB Implementations
type DB interface {
	SaveRoom(Room matrix.Room) error
	GetRooms() (rooms map[string]matrix.Room, err error)
	GetRoom(roomID string) (room matrix.Room, err error)

	SaveUser(user matrix.User) error
	GetUsers() (map[string]matrix.User, error)
	GetCurrentUser() (matrix.User, error)
}
