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

// get followed post
// @Tags Social
// @Router /social/post [get]
// @Summary Ambil Postingan dari Pengguna yang Diikuti (Feed)
// @Description Mengambil daftar postingan terbaru dari pengguna yang sudah di-follow oleh pengguna yang sedang login.
// @Produce json
// @Security JWTtoken
// @Success 200 {object} models.ResponseData{data=[]models.PostDetail} "Daftar postingan berhasil dimuat"
// @Failure 401 {object} models.UnauthorizedResponse "Token tidak valid atau tidak ada"
// @Failure 500 {object} models.InternalErrorResponse "Kesalahan server internal"
func (s *SocialHandler) GetFollowedPost(ctx *gin.Context) {
	// get user id from token
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

	// query to database
	posts, err := s.socRep.GetPost(ctx.Request.Context(), userID)
	if err != nil {
		log.Println("error cause \n", err)
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
			Data: posts,
		})
	}
}

// @Tags Social
// @Router /social/post [post]
// @Summary Like atau Komentar Postingan
// @Description Memberi like pada postingan, menghapus like, atau menambahkan komentar pada postingan.
// @Accept json
// @Produce json
// @Security JWTtoken
// @Param request body models.InteracTionPost true "Detail interaksi (Post ID, status like, dan/atau komentar)"
// @Success 200 {object} models.Response "Interaksi (like/komentar) berhasil diproses"
// @Failure 400 {object} models.BadRequestResponse "Input tidak valid (misalnya Post ID kosong)"
// @Failure 401 {object} models.UnauthorizedResponse "Token tidak valid atau tidak ada"
// @Failure 500 {object} models.InternalErrorResponse "Kesalahan server internal"
func (s *SocialHandler) LikeAndOrCommentPost(ctx *gin.Context) {
	// get user id from token
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

	// binding data
	var body models.InteracTionPost
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println("failed binding data \ncause, ", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      400,
			},
			Err: "Failed binding data...",
		})
		return
	}

	// query to database
	if err := s.socRep.LikeAndOrCommentPost(ctx.Request.Context(), userID, body); err != nil {
		log.Println("failed query data to database \ncause, ", err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      http.StatusInternalServerError,
			},
			Err: "Internal server error",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, models.Response{
			IsSuccess: true,
			Code:      http.StatusOK,
			Msg:       "like or comment successfully",
		})
	}
}
