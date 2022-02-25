package room

import (
	"group-project2/deliveries/controllers/common"
	"group-project2/deliveries/middlewares"
	_RoomRepo "group-project2/repositories/room"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type RoomController struct {
	repo _RoomRepo.Room
}

func New(repository _RoomRepo.Room) *RoomController {
	return &RoomController{
		repo: repository,
	}
}

func (ctl *RoomController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)
		IsRenter := middlewares.ExtractTokenIsRenter(c)
		if !IsRenter {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("client tidak terautorisasi, hanya renter yang diizinkan untuk menambahkan room"))
		}

		NewRoom := RequestCreateRoom{}
		if err := c.Bind(&NewRoom); err != nil || strings.TrimSpace(NewRoom.Name) == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari client tidak sesuai, nama tidak boleh kosong"))
		}

		res, err := ctl.repo.Insert(NewRoom.ToEntityRoom(uint(UserID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError("gagal membuat room baru"))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan room baru", ToResponseCreateRoom(res)))
	}
}

func (ctl *RoomController) GetAllRooms() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ctl.repo.GetAllRooms()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua room", ToResponseGetAllRooms(res)))
	}
}

func (ctl *RoomController) GetRoomByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		RoomID, _ := strconv.Atoi(c.Param("id"))

		res, err := ctl.repo.GetRoomByID(uint(RoomID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError("room tidak ditemukan"))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan room", ToResponseGetRoomByID(res)))
	}
}

func (ctl *RoomController) GetRoomsByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID, _ := strconv.Atoi(c.QueryParam("user_id"))

		res, err := ctl.repo.GetRoomsByUserID(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua room berdasarkan user", ToGetRoomsByUserID(res)))
	}
}

func (ctl *RoomController) GetRoomsByCity() echo.HandlerFunc {
	return func(c echo.Context) error {
		City := c.QueryParam("city")

		res, err := ctl.repo.GetRoomsByCity(City)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua room berdasarkan kota", ToResponseGetRoomsByCity(res)))
	}
}

func (ctl *RoomController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		RoomID, _ := strconv.Atoi(c.Param("id"))
		IsRenter := middlewares.ExtractTokenIsRenter(c)
		RoomUpdate := RequestUpdate{}

		if !IsRenter {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("client tidak terautorisasi, hanya renter yang diizinkan untuk mengupdate room"))
		}

		if err := c.Bind(&RoomUpdate); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari client tidak sesuai"))
		}

		res, err := ctl.repo.Update(RoomUpdate.ToEntityRoom(uint(RoomID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengupdate room", ToResponseUpdate(res)))
	}
}

func (ctl *RoomController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)
		IsRenter := middlewares.ExtractTokenIsRenter(c)
		if !IsRenter {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("client tidak terautorisasi, hanya renter yang diizinkan untuk menghapus room"))
		}

		RoomID, _ := strconv.Atoi(c.Param("id"))

		err := ctl.repo.Delete(uint(UserID), uint(RoomID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus room", err))
	}
}
