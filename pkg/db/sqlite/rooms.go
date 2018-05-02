package sqlite

import "github.com/Nordgedanken/Morpheusv2/pkg/matrix"

func (s *SQLite) SaveRoom(Room matrix.Room) error {
	return nil
}

func (s *SQLite) GetRooms() (rooms map[string]matrix.Room, err error) {
	return nil, nil
}

func (s *SQLite) GetRoom(roomID string) (room matrix.Room, err error) {
	return nil, nil
}
