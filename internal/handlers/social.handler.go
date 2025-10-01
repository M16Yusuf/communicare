package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m16yusuf/communicare/internal/models"
	"github.com/m16yusuf/communicare/internal/repositories"
	"github.com/m16yusuf/communicare/internal/utils"
)

type SocialHandler struct {
	socRep *repositories.SocialRepository
}

func NewSocialHandler(sr *repositories.SocialRepository) *SocialHandler {
	return &SocialHandler{socRep: sr}
}

// follow
// @Tags Social
// @Router /social/follow/{user_id} [post]
// @Summary Mengikuti Pengguna
// @Description Memulai mengikuti pengguna lain berdasarkan ID pengguna. Membutuhkan otentikasi.
// @Produce json
// @Security JWTtoken
// @Param user_id path string true "ID pengguna yang ingin diikuti"
// @Success 200 {object} models.Response "Berhasil mengikuti pengguna"
// @Failure 401 {object} models.UnauthorizedResponse "Token tidak valid atau tidak ada"
// @Failure 500 {object} models.InternalErrorResponse "Kesalahan server internal"
func (s *SocialHandler) FollowAUser(ctx *gin.Context) {
	// get user id from token
	follower, err := utils.GetUserFromCtx(ctx)
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

	following := ctx.Param("user_id")
	// add database follow a user
	if err := s.socRep.FollowUser(ctx.Request.Context(), following, follower); err != nil {
		log.Println("error cause \n", err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      http.StatusInternalServerError,
			},
			Err: "internal server error",
		})
	} else {
		ctx.JSON(http.StatusOK, models.Response{
			IsSuccess: true,
			Code:      http.StatusOK,
			Msg:       "success follow a user",
		})
	}
}
