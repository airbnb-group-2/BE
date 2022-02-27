package image

import (
	"errors"
	"fmt"
	"group-project2/configs"
	"group-project2/deliveries/controllers/common"
	"group-project2/deliveries/middlewares"
	_ImageRepo "group-project2/repositories/image"
	awss3 "group-project2/services/aws-s3"
	"net/http"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ImageController struct {
	repo    _ImageRepo.Image
	awsConf *configs.AppConfig
	awsSess *session.Session
}

func New(repository _ImageRepo.Image, config *configs.AppConfig, aws *session.Session) *ImageController {
	return &ImageController{
		repo:    repository,
		awsConf: config,
		awsSess: aws,
	}
}

func Upload(ctl *ImageController, c echo.Context) (string, error) {
	imageUpload := RequestImage{}
	if err := c.Bind(&imageUpload); err != nil || imageUpload.RoomID == 0 {
		log.Info(errors.New("input dari client tidak sesuai, room_id tidak boleh kosong"))
		return "", err
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Info(err)
		return "", err
	}

	file.Filename = fmt.Sprintf("room%s-%s", strconv.Itoa(int(imageUpload.RoomID)), file.Filename)

	link, err := awss3.DoUpload(ctl.awsSess, *file, ctl.awsConf.S3_REGION)
	if err != nil {
		log.Info(err)
		return "", err
	}

	return link, nil
}

// func Delete(ctl *ImageController, c echo.Context) (string, error) {
// 	imageUpload := RequestImage{}
// 	if err := c.Bind(&imageUpload); err != nil || imageUpload.RoomID == 0 {
// 		log.Info(errors.New("input dari client tidak sesuai, room_id tidak boleh kosong"))
// 		return "", err
// 	}

// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		log.Info(err)
// 		return "", err
// 	}

// 	file.Filename = fmt.Sprintf("room%s-%s", strconv.Itoa(int(imageUpload.RoomID)), file.Filename)

// 	link, err := awss3.DoUpload(ctl.awsSess, *file, ctl.awsConf.S3_REGION)
// 	if err != nil {
// 		log.Info(err)
// 		return "", err
// 	}

// 	return link, nil
// }

func (ctl *ImageController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		IsRenter := middlewares.ExtractTokenIsRenter(c)
		if !IsRenter {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("client tidak terautorisasi, hanya renter yang diizinkan untuk menambahkan gambar"))
		}

		NewImage := RequestImage{}
		if err := c.Bind(&NewImage); err != nil || NewImage.RoomID == 0 {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari client tidak sesuai, room_id tidak boleh kosong"))
		}

		link, err := Upload(ctl, c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("gagal mengunggah gambar"))
		}
		NewImage.Link = link

		res, err := ctl.repo.Insert(NewImage.ToEntityImage())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError("gagal menambahkan gambar baru"))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan gambar baru", ToResponseCreateImage(res)))
	}
}

func (ctl *ImageController) GetImagesByRoomID() echo.HandlerFunc {
	return func(c echo.Context) error {
		RoomID, _ := strconv.Atoi(c.QueryParam("room_id"))

		res, err := ctl.repo.GetImagesByRoomID(uint(RoomID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua gambar berdasarkan room_id", ToResponseGetImagesByRoomID(res)))
	}
}

func (ctl *ImageController) GetImageByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ImageID, _ := strconv.Atoi(c.Param("id"))

		res, err := ctl.repo.GetImageByID(uint(ImageID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError("gagal mendapatkan gambar berdasarkan id"))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan gambar berdasarkan id", ToResponseGetImageByID(res)))
	}
}

func (ctl *ImageController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ImageID, _ := strconv.Atoi(c.Param("id"))
		IsRenter := middlewares.ExtractTokenIsRenter(c)
		if !IsRenter {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("client tidak terautorisasi, hanya renter yang diizinkan untuk menambahkan gambar"))
		}

		ImageUpdate := RequestImageUpdate{}
		if err := c.Bind(&ImageUpdate); err != nil || strings.TrimSpace(ImageUpdate.Link) == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari client tidak sesuai, link tidak boleh kosong"))
		}

		res, err := ctl.repo.Update(ImageUpdate.ToEntityImage(uint(ImageID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengupdate gambar", ToResponseImageUpdate(res)))
	}
}

func (ctl *ImageController) DeleteImageByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ImageID, _ := strconv.Atoi(c.Param("id"))
		IsRenter := middlewares.ExtractTokenIsRenter(c)
		if !IsRenter {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("client tidak terautorisasi, hanya renter yang diizinkan untuk menambahkan gambar"))
		}

		// res, _ := ctl.repo.GetImageByID(uint(ImageID))

		// link := ToResponseGetImageByID(res).Link
		// pattern, _ := regexp.Compile(`(?:[^/][\d\w\.]+)$(?<=\.\w{3,4})`)
		// fileName := pattern.FindString(link)

		err := ctl.repo.DeleteImageByID(uint(ImageID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus gambar", err))
	}
}

func (ctl *ImageController) DeleteImageByRoomID() echo.HandlerFunc {
	return func(c echo.Context) error {
		RoomID, _ := strconv.Atoi(c.QueryParam("room_id"))
		IsRenter := middlewares.ExtractTokenIsRenter(c)
		if !IsRenter {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("client tidak terautorisasi, hanya renter yang diizinkan untuk menambahkan gambar"))
		}

		err := ctl.repo.DeleteImageByRoomID(uint(RoomID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus gambar berdasarkan room_id", err))
	}
}
