package user

import (
	"errors"
	U "group-project2/entities/user"
	"group-project2/repositories/hash"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Insert(NewUser U.Users) (U.Users, error) {
	NewUser.Password, _ = hash.HashPassword(NewUser.Password)
	if err := repo.db.Create(&NewUser).Error; err != nil {
		log.Warn(err)
		return U.Users{}, err
	}
	return NewUser, nil
}

func (repo *UserRepository) GetUserByID(UserID uint) (U.Users, error) {
	User := U.Users{}
	if err := repo.db.First(&User, UserID).Error; err != nil {
		log.Warn(err)
		return User, err
	}
	return User, nil
}

func (repo *UserRepository) Update(UpdatedUser U.Users) (U.Users, error) {
	if UpdatedUser.Password != "" {
		UpdatedUser.Password, _ = hash.HashPassword(UpdatedUser.Password)
	}

	res := repo.db.Model(&UpdatedUser).Updates(UpdatedUser)
	if res.RowsAffected == 0 {
		return U.Users{}, errors.New("tidak ada perubahan pada data user")
	}
	repo.db.First(&UpdatedUser)
	return UpdatedUser, nil
}

func (repo *UserRepository) SetRenter(UserID uint) (U.Users, error) {
	User := U.Users{}
	res := repo.db.Model(&User).Where("id = ?", UserID).Update("is_renter", true)
	if res.RowsAffected == 0 {
		return U.Users{}, errors.New("gagal menjadikan renter")
	}
	repo.db.First(&User, UserID)
	return User, nil
}

func (repo *UserRepository) DeleteByID(UserID uint) error {
	if RowsAffected := repo.db.Delete(&U.Users{}, UserID).RowsAffected; RowsAffected == 0 {
		return errors.New("tidak ada user yang dihapus")
	}
	return nil
}
