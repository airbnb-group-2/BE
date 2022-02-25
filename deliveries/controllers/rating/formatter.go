package rating

import R "group-project2/entities/rating"

type RequestRating struct {
	Star uint `json:"star" form:"star"`
}

func (Req RequestRating) ToEntityRating(UserID, RoomID uint) R.Ratings {
	return R.Ratings{
		Star:   Req.Star,
		UserID: UserID,
		RoomID: RoomID,
	}
}

type ResponseCreateRating struct {
	Star   uint `json:"star"`
	UserID uint `json:"user_id"`
	RoomID uint `json:"room_id"`
}

func ToResponseCreateRating(Rating R.Ratings) ResponseCreateRating {
	return ResponseCreateRating{
		Star:   Rating.Star,
		UserID: Rating.UserID,
		RoomID: Rating.RoomID,
	}
}

type ResponseGetRating struct {
	ID     uint `json:"id"`
	Star   uint `json:"star"`
	UserID uint `json:"user_id"`
	RoomID uint `json:"room_id"`
}

func ToResponseGetRatingsByRoomID(Ratings []R.Ratings) []ResponseGetRating {
	Responses := make([]ResponseGetRating, len(Ratings))

	for i := 0; i < len(Ratings); i++ {
		Responses[i].ID = Ratings[i].ID
		Responses[i].Star = Ratings[i].Star
		Responses[i].UserID = Ratings[i].UserID
		Responses[i].RoomID = Ratings[i].RoomID
	}

	return Responses
}
