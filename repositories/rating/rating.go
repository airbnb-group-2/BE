package rating

import (
	"errors"
	R "group-project2/entities/rating"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type RatingRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *RatingRepository {
	return &RatingRepository{
		db: db,
	}
}

func (repo *RatingRepository) Insert(NewRating R.Ratings) (R.Ratings, error) {
	if err := repo.db.Create(&NewRating).Error; err != nil {
		log.Warn(err)
		return R.Ratings{}, err
	}
	return NewRating, nil
}

func (repo *RatingRepository) GetRatingsByRoomID(RoomID uint) ([]R.Ratings, error) {
	Ratings := []R.Ratings{}
	if RowsAffected := repo.db.Where("room_id = ?", RoomID).Find(&Ratings).RowsAffected; RowsAffected == 0 {
		return nil, errors.New("belum ada penilaian")
	}
	return Ratings, nil
}

// func (repo *RatingRepository) GetImageByID(ImageID uint) (I.Images, error) {
// 	Image := I.Images{}
// 	if err := repo.db.First(&Image, ImageID).Error; err != nil {
// 		log.Warn(err)
// 		return I.Images{}, err
// 	}
// 	return Image, nil
// }

// func (repo *RatingRepository) Update(ImageUpdate I.Images) (I.Images, error) {
// 	if RowsAffected := repo.db.Model(&ImageUpdate).Updates(ImageUpdate).RowsAffected; RowsAffected == 0 {
// 		return I.Images{}, errors.New("tidak ada perubahan pada data gambar")
// 	}
// 	repo.db.First(&ImageUpdate)
// 	return ImageUpdate, nil
// }

// func (repo *RatingRepository) DeleteImageByID(ImageID uint) error {
// 	if RowsAffected := repo.db.Delete(&I.Images{}, ImageID).RowsAffected; RowsAffected == 0 {
// 		return errors.New("tidak ada gambar yang dihapus")
// 	}
// 	return nil
// }

// func (repo *RatingRepository) DeleteImageByRoomID(RoomID uint) error {
// 	if RowsAffected := repo.db.Where("room_id = ?", RoomID).Delete(&I.Images{}).RowsAffected; RowsAffected == 0 {
// 		return errors.New("tidak ada gambar berdasarkan room_id yang dihapus")
// 	}
// 	return nil
// }
