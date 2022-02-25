package rating

import R "group-project2/entities/rating"

type Rating interface {
	Insert(NewRating R.Ratings) (R.Ratings, error)
	GetRatingsByRoomID(RoomID uint) ([]R.Ratings, error)
}
