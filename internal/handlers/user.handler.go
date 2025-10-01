package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m16yusuf/communicare/internal/models"
	"github.com/m16yusuf/communicare/internal/repositories"
	"github.com/m16yusuf/communicare/internal/utils"
)

type UserHandler struct {
	userRep *repositories.UserRepository
}

func NewUserHandler(userRep *repositories.UserRepository) *UserHandler {
	return &UserHandler{userRep: userRep}
}

// update profile
// @Tags User
// @Router /users [post]
// @Summary Perbarui Profil Pengguna
// @Description Memperbarui detail profil pengguna termasuk nama lengkap, bio, dan avatar. Membutuhkan otentikasi.
// @Accept multipart/form-data
// @Produce json
// @Security JWTtoken
// @Param fullname formData string false "Nama lengkap pengguna"
// @Param bio formData string false "Bio/deskripsi singkat pengguna"
// @Param profile_picture formData file false "File gambar untuk Avatar (Max 2MB)"
// @Success 200 {object} models.ResponseData{data=models.ProfileResponse} "Profil berhasil diperbarui"
// @Failure 400 {object} models.BadRequestResponse "Input tidak valid atau kegagalan upload file"
// @Failure 401 {object} models.UnauthorizedResponse "Token tidak valid atau tidak ada"
// @Failure 500 {object} models.InternalErrorResponse "Kesalahan server internal"
func (uh *UserHandler) UpdateProfile(ctx *gin.Context) {
	// get user it from token
	userID, err := utils.GetUserFromCtx(ctx)
	if err != nil {
		log.Println("error cause: ", err.Error())
		ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      http.StatusUnauthorized,
			},
			Err: err.Error(),
		})
		return
	}

	var body models.ProfileRequest
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println("error cause: ", err.Error())
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      http.StatusBadRequest,
			},
			Err: err.Error(),
		})
		return
	}

	// process the image
	file := body.Avatar
	filename, err := utils.FileUpload(ctx, file, "avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      http.StatusBadRequest,
			},
			Err: err.Error(),
		})
		return
	}

	profile := models.Profile{
		Id:       userID,
		Fullname: body.Fullname,
		Avatar:   &filename,
		Bio:      body.Bio,
	}
	// query updated user to database
	updatedUser, err := uh.userRep.UpdateProfileUser(ctx.Request.Context(), profile)
	if err != nil {
		log.Println("error cause: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      http.StatusInternalServerError,
			},
			Err: "internal server error",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, models.ResponseData{
			Response: models.Response{
				IsSuccess: true,
				Code:      http.StatusOK,
			},
			Data: updatedUser,
		})
	}
}
