package room

import R "group-project2/entities/room"

type Room interface {
	Insert(NewRoom R.Rooms) (R.Rooms, error)
	Get() ([]R.Rooms, error)
	GetRoomByID(RoomID uint) (R.Rooms, error)
	GetRoomsByCity(City string) ([]R.Rooms, error)
	Update(RoomUpdate R.Rooms) (R.Rooms, error)
	Delete(RoomID uint) error
}
