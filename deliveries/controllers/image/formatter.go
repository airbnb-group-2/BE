package image

import (
	I "group-project2/entities/image"

	"gorm.io/gorm"
)

type RequestImage struct {
	Link   string
	RoomID uint `json:"room_id" form:"room_id"`
}

func (Req RequestImage) ToEntityImage() I.Images {
	return I.Images{
		Link:   Req.Link,
		RoomID: Req.RoomID,
	}
}

type ResponseCreateImage struct {
	Link   string `json:"link"`
	RoomID uint   `json:"room_id"`
}

func ToResponseCreateImage(Image I.Images) ResponseCreateImage {
	return ResponseCreateImage{
		Link:   Image.Link,
		RoomID: Image.RoomID,
	}
}

type ResponseGetImage struct {
	ID     uint   `json:"id"`
	Link   string `json:"link"`
	RoomID uint   `json:"room_id"`
}

func ToResponseGetImagesByRoomID(Images []I.Images) []ResponseGetImage {
	Responses := make([]ResponseGetImage, len(Images))

	for i := 0; i < len(Images); i++ {
		Responses[i].ID = Images[i].ID
		Responses[i].Link = Images[i].Link
		Responses[i].RoomID = Images[i].RoomID
	}

	return Responses
}

func ToResponseGetImageByID(Image I.Images) ResponseGetImage {
	return ResponseGetImage{
		ID:     Image.ID,
		Link:   Image.Link,
		RoomID: Image.RoomID,
	}
}

type RequestImageUpdate struct {
	Link string `json:"link" form:"link"`
}

func (Req RequestImageUpdate) ToEntityImage(ImageID uint) I.Images {
	return I.Images{
		Model: gorm.Model{ID: ImageID},
		Link:  Req.Link,
	}
}

type ResponseImageUpdate struct {
	ID     uint   `json:"id"`
	Link   string `json:"link"`
	RoomID uint   `json:"room_id"`
}

func ToResponseImageUpdate(Image I.Images) ResponseImageUpdate {
	return ResponseImageUpdate{
		ID:     Image.ID,
		Link:   Image.Link,
		RoomID: Image.RoomID,
	}
}
