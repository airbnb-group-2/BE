package room

import R "group-project2/entities/room"

type Room interface {
	Insert(NewRoom R.Rooms) (R.Rooms, error)
	GetAllRooms() ([]R.Rooms, error)
	GetRoomByID(RoomID uint) (R.Rooms, error)
	GetRoomsByUserID(UserID uint) ([]R.Rooms, error)
	GetRoomsByCity(City string) ([]R.Rooms, error)
	Update(RoomUpdate R.Rooms) (R.Rooms, error)
	Delete(UserID, RoomID uint) error
}
