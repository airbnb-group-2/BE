package book

import (
	"group-project2/deliveries/controllers/common"
	"group-project2/deliveries/middlewares"
	_B "group-project2/repositories/book"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	repo _B.Book
}

func New(repository _B.Book) *BookController {
	return &BookController{
		repo: repository,
	}
}

func (ctl *BookController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)

		NewBook := RequestCreateBook{}
		if err := c.Bind(&NewBook); err != nil || strings.TrimSpace(NewBook.Phone) == "" || NewBook.PaymentMethodID == 0 || NewBook.RoomID == 0 {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari client tidak sesuai, phone, payment_method_id atau room_id tidak boleh kosong"))
		}

		res, err := ctl.repo.Insert(NewBook.ToEntityBook(uint(UserID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError("gagal membuat booking baru"))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan booking baru", ToResponseCreateBook(res)))
	}
}

func (ctl *BookController) GetAllBooksByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)

		res, err := ctl.repo.GetAllBooksByUserID(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses mendapatkan semua booking", res))
	}
}

func (ctl *BookController) GetBookHistoryByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)

		res, err := ctl.repo.GetBookHistoryByUserID(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses mendapatkan history user", res))
	}
}
