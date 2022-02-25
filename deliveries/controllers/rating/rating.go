package rating

import (
	"group-project2/deliveries/controllers/common"
	"group-project2/deliveries/middlewares"
	_RatingRepo "group-project2/repositories/rating"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RatingController struct {
	repo _RatingRepo.Rating
}

func New(repository _RatingRepo.Rating) *RatingController {
	return &RatingController{
		repo: repository,
	}
}

func (ctl *RatingController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)
		RoomID, _ := strconv.Atoi(c.QueryParam("room_id"))

		NewRating := RequestRating{}
		if err := c.Bind(&NewRating); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari client tidak sesuai"))
		}

		res, err := ctl.repo.Insert(NewRating.ToEntityRating(uint(UserID), uint(RoomID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError("gagal menambahkan rating baru"))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan rating baru", ToResponseCreateRating(res)))
	}
}

func (ctl *RatingController) GetRatingsByRoomID() echo.HandlerFunc {
	return func(c echo.Context) error {
		RoomID, _ := strconv.Atoi(c.QueryParam("room_id"))

		res, err := ctl.repo.GetRatingsByRoomID(uint(RoomID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua rating berdasarkan room_id", ToResponseGetRatingsByRoomID(res)))
	}
}
