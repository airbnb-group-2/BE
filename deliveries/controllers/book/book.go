package book

import (
	"fmt"
	"group-project2/deliveries/controllers/common"
	"group-project2/deliveries/middlewares"
	_B "group-project2/repositories/book"
	midtranspay "group-project2/services/midtrans-pay"
	"net/http"
	"strconv"
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

var (
	midtransConn = midtranspay.InitConnection()
)

func (ctl *BookController) GetStatusID() echo.HandlerFunc {
	return func(c echo.Context) error {
		BookID, _ := strconv.Atoi(c.Param("id"))
		res, err := midtranspay.Notification(midtransConn, uint(BookID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}

		if res == "status pending" {
			return c.JSON(http.StatusOK, common.Success(http.StatusOK, "pembayaran tertunda", nil))
		} else if res == "status settlement" {
			res, err := ctl.repo.SetPaid(uint(BookID))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
			}
			return c.JSON(http.StatusOK, common.Success(http.StatusOK, "pembayaran sukses", ToResponseSetPaid(res)))
		} else if res == "status cancel" {
			res, err := ctl.repo.SetCancel(uint(BookID))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
			}
			return c.JSON(http.StatusOK, common.Success(http.StatusOK, "pembayaran gagal", ToResponseSetCancel(res)))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "status transaksi:", res))
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
		trxResp := midtranspay.CreateTransaction(midtransConn, res.ID)
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan booking baru", ToResponseCreateBook(res, trxResp)))
		// return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan booking baru", res2))
	}
}

func (ctl *BookController) GetAllBooksByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)

		res, err := ctl.repo.GetAllBooksByUserID(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses mendapatkan semua booking", ToResponseGet(res)))
	}
}

func (ctl *BookController) GetBookHistoryByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)

		res, err := ctl.repo.GetBookHistoryByUserID(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses mendapatkan history user", ToResponseGet(res)))
	}
}

func (ctl *BookController) IsAvailable() echo.HandlerFunc {
	return func(c echo.Context) error {
		Checker := RequestIsAvailable{}
		if err := c.Bind(&Checker); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari client tidak sesuai, room_id, check_in_reserved atau check_out_reserved tidak boleh kosong"))
		}

		res := ctl.repo.IsAvailable(uint(Checker.RoomID), Checker.CheckInReserved, Checker.CheckOutReserved)
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengecek ketersediaan room", res))
	}
}

func (ctl *BookController) SetPaid() echo.HandlerFunc {
	return func(c echo.Context) error {
		BookID, _ := strconv.Atoi(c.Param("id"))
		UserID := middlewares.ExtractTokenUserID(c)

		res, err := ctl.repo.SetPaid(uint(BookID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		fmt.Println("user_id:", UserID, "sudah membayar")
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status booking menjadi paid", ToResponseSetPaid(res)))
	}
}

func (ctl *BookController) SetCancel() echo.HandlerFunc {
	return func(c echo.Context) error {
		BookID, _ := strconv.Atoi(c.Param("id"))
		UserID := middlewares.ExtractTokenUserID(c)

		res, err := ctl.repo.SetCancel(uint(BookID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		fmt.Println("user_id:", UserID, "batal booking")
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status booking menjadi cancel", ToResponseSetCancel(res)))
	}
}

func (ctl *BookController) SetCheckInTime() echo.HandlerFunc {
	return func(c echo.Context) error {
		BookID, _ := strconv.Atoi(c.Param("id"))
		UserID := middlewares.ExtractTokenUserID(c)

		CheckInTime := RequestCheckInTime{}
		if err := c.Bind(&CheckInTime); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari client tidak sesuai"))
		}

		res, err := ctl.repo.SetCheckInTime(uint(BookID), CheckInTime.CheckInTime)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		fmt.Println("user_id:", UserID, "melakukan check in")
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menyetel check_in_time", ToResponseCheckInTime(res)))
	}
}

func (ctl *BookController) SetCheckOutTime() echo.HandlerFunc {
	return func(c echo.Context) error {
		BookID, _ := strconv.Atoi(c.Param("id"))
		UserID := middlewares.ExtractTokenUserID(c)

		CheckOutTime := RequestCheckOutTime{}
		if err := c.Bind(&CheckOutTime); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari client tidak sesuai"))
		}

		res, err := ctl.repo.SetCheckOutTime(uint(BookID), CheckOutTime.CheckOutTime)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		fmt.Println("user_id:", UserID, "melakukan check out")
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menyetel check_out_time", ToResponseCheckOutTime(res)))
	}
}
