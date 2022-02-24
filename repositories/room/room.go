package room

import (
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"

	R "group-project2/entities/room"
)

type RoomRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

func (repo *RoomRepository) Insert(NewRoom R.Rooms) (R.Rooms, error) {
	if err := repo.db.Create(&NewRoom).Error; err != nil {
		log.Warn(err)
		return R.Rooms{}, err
	}
	repo.db.Table("rooms").Where("user_id = ? AND name = ?", NewRoom.UserID, NewRoom.Name).First(&NewRoom)
	return NewRoom, nil
}

func (repo *RoomRepository) GetAllRooms() ([]R.Rooms, error) {
	Rooms := []R.Rooms{}
	if RowsAffected := repo.db.Find(&Rooms).RowsAffected; RowsAffected == 0 {
		return []R.Rooms{}, errors.New("belum ada room yang terdaftar")
	}
	return Rooms, nil
}

func (repo *RoomRepository) GetRoomByID(RoomID uint) (R.Rooms, error) {
	Room := R.Rooms{}
	if err := repo.db.First(&Room, RoomID).Error; err != nil {
		log.Warn(err)
		return R.Rooms{}, err
	}
	return Room, nil
}

func (repo *RoomRepository) GetRoomsByUserID(UserID uint) ([]R.Rooms, error) {
	Rooms := []R.Rooms{}
	if RowsAffected := repo.db.Table("rooms").Where("user_id = ?", UserID).Find(&Rooms).RowsAffected; RowsAffected == 0 {
		return nil, errors.New("tidak terdapat room dari user tersebut")
	}
	return Rooms, nil
}

func (repo *RoomRepository) GetRoomsByCity(City string) ([]R.Rooms, error) {
	Rooms := []R.Rooms{}
	if RowsAffected := repo.db.Table("rooms").Where("city = ?", City).Find(&Rooms).RowsAffected; RowsAffected == 0 {
		return nil, errors.New("tidak terdapat room di kota tersebut")
	}
	return Rooms, nil
}

func (repo *RoomRepository) Update(RoomUpdate R.Rooms) (R.Rooms, error) {
	if RowsAffected := repo.db.Model(&RoomUpdate).Updates(RoomUpdate).RowsAffected; RowsAffected == 0 {
		return R.Rooms{}, errors.New("tidak ada perubahan pada data room")
	}
	repo.db.First(&RoomUpdate)
	return RoomUpdate, nil
}

func (repo *RoomRepository) Delete(UserID, RoomID uint) error {
	if RowsAffected := repo.db.Table("room").Where("user_id = ? AND room_id = ?", UserID, RoomID).Delete(&R.Rooms{}, RoomID).RowsAffected; RowsAffected == 0 {
		return errors.New("tidak ada room yang dihapus")
	}
	return nil
}
