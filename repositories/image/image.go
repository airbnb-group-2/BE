package image

import (
	"errors"
	I "group-project2/entities/image"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ImageRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ImageRepository {
	return &ImageRepository{
		db: db,
	}
}

func (repo *ImageRepository) Insert(NewImage I.Images) (I.Images, error) {
	if err := repo.db.Create(&NewImage).Error; err != nil {
		log.Warn(err)
		return I.Images{}, err
	}
	return NewImage, nil
}

func (repo *ImageRepository) GetImagesByRoomID(RoomID uint) ([]I.Images, error) {
	Images := []I.Images{}
	if RowsAffected := repo.db.Where("room_id = ?", RoomID).Find(&Images).RowsAffected; RowsAffected == 0 {
		return nil, errors.New("belum ada gambar yang tersimpan")
	}
	return Images, nil
}

func (repo *ImageRepository) GetImageByID(ImageID uint) (I.Images, error) {
	Image := I.Images{}
	if err := repo.db.First(&Image, ImageID).Error; err != nil {
		log.Warn(err)
		return I.Images{}, err
	}
	return Image, nil
}

func (repo *ImageRepository) Update(ImageUpdate I.Images) (I.Images, error) {
	if RowsAffected := repo.db.Model(&ImageUpdate).Updates(ImageUpdate).RowsAffected; RowsAffected == 0 {
		return I.Images{}, errors.New("tidak ada perubahan pada data gambar")
	}
	repo.db.First(&ImageUpdate)
	return ImageUpdate, nil
}

func (repo *ImageRepository) DeleteImageByID(ImageID uint) error {
	if RowsAffected := repo.db.Delete(&I.Images{}, ImageID).RowsAffected; RowsAffected == 0 {
		return errors.New("tidak ada gambar yang dihapus")
	}
	return nil
}

func (repo *ImageRepository) DeleteImageByRoomID(RoomID uint) error {
	if RowsAffected := repo.db.Where("room_id = ?", RoomID).Delete(&I.Images{}).RowsAffected; RowsAffected == 0 {
		return errors.New("tidak ada gambar berdasarkan room_id yang dihapus")
	}
	return nil
}
