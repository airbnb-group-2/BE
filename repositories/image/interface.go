package image

import I "group-project2/entities/image"

type Image interface {
	Insert(NewImage I.Images) (I.Images, error)
	GetImagesByRoomID(RoomID uint) ([]I.Images, error)
	GetImageByID(ImageID uint) (I.Images, error)
	Update(ImageUpdate I.Images) (I.Images, error)
	DeleteImageByID(ImageID uint) error
	DeleteImageByRoomID(RoomID uint) error
}
