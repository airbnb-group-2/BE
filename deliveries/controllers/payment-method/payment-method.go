package paymentmethod

import (
	"group-project2/deliveries/controllers/common"
	"group-project2/deliveries/middlewares"
	_PMRepo "group-project2/repositories/payment-method"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type PaymentMethodController struct {
	repo _PMRepo.PaymentMethod
}

func New(repository _PMRepo.PaymentMethod) *PaymentMethodController {
	return &PaymentMethodController{
		repo: repository,
	}
}

func (ctl *PaymentMethodController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		IsRenter := middlewares.ExtractTokenIsRenter(c)
		if !IsRenter {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("client tidak terautorisasi, hanya renter yang diizinkan untuk menambahkan payment method"))
		}

		NewPM := RequestCreatePM{}
		if err := c.Bind(&NewPM); err != nil || strings.TrimSpace(NewPM.Name) == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari client tidak sesuai, nama payment method tidak boleh kosong"))
		}

		res, err := ctl.repo.Insert(NewPM.ToEntityPM())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError("gagal menambahkan payment method baru"))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan payment method baru", ToResponseCreatePM(res)))
	}
}

func (ctl *PaymentMethodController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ctl.repo.Get()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua payment method", ToResponseGet(res)))
	}
}

func (ctl *PaymentMethodController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		IsRenter := middlewares.ExtractTokenIsRenter(c)
		if !IsRenter {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("client tidak terautorisasi, hanya renter yang diizinkan untuk menghapus payment method"))
		}

		PMID, _ := strconv.Atoi(c.Param("id"))

		err := ctl.repo.Delete(uint(PMID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus payment method", err))
	}
}
