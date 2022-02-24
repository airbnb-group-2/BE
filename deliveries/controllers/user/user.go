package user

import (
	"group-project2/deliveries/controllers/common"
	"group-project2/deliveries/middlewares"
	_UserRepo "group-project2/repositories/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo _UserRepo.User
}

func New(repository _UserRepo.User) *UserController {
	return &UserController{
		repo: repository,
	}
}

func (ctl *UserController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		NewUser := RequestCreateUser{}

		if err := c.Bind(&NewUser); err != nil || NewUser.Name == "" || NewUser.Email == "" || NewUser.Password == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai, nama, email atau password tidak boleh kosong"))
		}

		res, err := ctl.repo.Insert(NewUser.ToEntityUser())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan user baru", ToResponseCreateUser(res)))
	}
}

func (ctl *UserController) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)

		res, err := ctl.repo.GetUserByID(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan user berdasarkan ID", ToResponseGetByID(res)))
	}
}

func (ctl *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)
		var UpdatedUser = RequestUpdateUser{}

		if err := c.Bind(&UpdatedUser); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("terdapat kesalahan input dari client"))
		}

		res, err := ctl.repo.Update(UpdatedUser.ToEntityUser(uint(UserID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses update user", ToResponseUpdate(res)))
	}
}

func (ctl *UserController) SetRenter() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)

		res, err := ctl.repo.SetRenter((uint(UserID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menjadikan renter", ToResponseUpdate(res)))
	}
}

func (ctl *UserController) DeleteByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)

		err := ctl.repo.DeleteByID(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus user", err))
	}
}
